#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
小红书图文发布工具 - MCP API版本
Created by 🦞 Lobster Journey Studio

功能：
- 通过MCP API发布小红书笔记
- 支持多图发布
- 参数验证
- 错误处理
- 发布历史记录
"""

import json
import time
from datetime import datetime
from pathlib import Path
from typing import List, Dict, Optional
import requests


class XiaohongshuPublisher:
    """小红书发布器（MCP API版本）"""

    def __init__(self, base_url: str = "http://localhost:18060"):
        """
        初始化发布器

        Args:
            base_url: MCP服务地址
        """
        self.base_url = base_url
        self.api_url = f"{base_url}/api/v1"
        self.log_file = Path("publish_log.json")

    def check_service(self) -> bool:
        """
        检查MCP服务状态

        Returns:
            bool: 服务是否正常
        """
        try:
            response = requests.get(f"{self.base_url}/health", timeout=5)
            result = response.json()
            return result.get('success', False)
        except Exception as e:
            print(f"❌ MCP服务检查失败: {e}")
            return False

    def publish(
        self,
        title: str,
        content: str,
        images: List[str],
        validate: bool = True
    ) -> Dict:
        """
        发布小红书笔记

        Args:
            title: 标题（最多20字）
            content: 正文（最多1000字）
            images: 图片路径列表（至少1张）
            validate: 是否验证参数

        Returns:
            Dict: 发布结果
        """
        # 参数验证
        if validate:
            if len(title) > 20:
                raise ValueError(f"标题长度超限: {len(title)} > 20")
            if len(content) > 1000:
                raise ValueError(f"正文长度超限: {len(content)} > 1000")
            if not images or len(images) == 0:
                raise ValueError("至少需要1张图片")

            # 检查图片文件是否存在
            for img in images:
                if not Path(img).exists():
                    raise FileNotFoundError(f"图片不存在: {img}")

        # 调用API
        try:
            response = requests.post(
                f"{self.api_url}/publish",
                json={
                    "title": title[:20],
                    "content": content[:1000],
                    "images": images
                },
                timeout=60
            )

            result = response.json()

            # 记录日志
            self._log_publish(result)

            return result

        except requests.Timeout:
            return {
                "success": False,
                "message": "请求超时，请检查网络"
            }
        except requests.ConnectionError:
            return {
                "success": False,
                "message": "连接失败，请检查MCP服务"
            }
        except Exception as e:
            return {
                "success": False,
                "message": f"未知错误: {e}"
            }

    def batch_publish(
        self,
        notes: List[Dict],
        interval: int = 10
    ) -> List[Dict]:
        """
        批量发布多篇笔记

        Args:
            notes: 笔记列表，每个笔记包含 title, content, images
            interval: 发布间隔（秒）

        Returns:
            List[Dict]: 发布结果列表
        """
        results = []

        print(f"\n{'='*60}")
        print(f"🦞 批量发布 {len(notes)} 篇笔记")
        print(f"{'='*60}\n")

        for i, note in enumerate(notes, 1):
            print(f"\n📝 发布第 {i}/{len(notes)} 篇...")
            print(f"   标题: {note['title']}")

            result = self.publish(
                title=note['title'],
                content=note['content'],
                images=note['images'],
                validate=False
            )

            results.append(result)

            if result['success']:
                print(f"   ✅ 发布成功")
            else:
                print(f"   ❌ 发布失败: {result.get('message')}")

            # 避免频率限制
            if i < len(notes):
                print(f"   ⏳ 等待 {interval} 秒...")
                time.sleep(interval)

        print(f"\n{'='*60}")
        print(f"✅ 批量发布完成")
        print(f"   成功: {sum(1 for r in results if r['success'])}")
        print(f"   失败: {sum(1 for r in results if not r['success'])}")
        print(f"{'='*60}\n")

        return results

    def _log_publish(self, result: Dict):
        """记录发布历史"""
        log_entry = {
            "timestamp": datetime.now().isoformat(),
            "title": result.get('data', {}).get('title', ''),
            "images": result.get('data', {}).get('images', 0),
            "status": result.get('data', {}).get('status', ''),
            "success": result.get('success', False),
            "message": result.get('message', '')
        }

        # 追加到日志文件
        with open(self.log_file, 'a', encoding='utf-8') as f:
            f.write(json.dumps(log_entry, ensure_ascii=False) + '\n')

    def get_publish_history(self, limit: int = 10) -> List[Dict]:
        """
        获取发布历史

        Args:
            limit: 返回条数限制

        Returns:
            List[Dict]: 发布历史列表
        """
        if not self.log_file.exists():
            return []

        with open(self.log_file, 'r', encoding='utf-8') as f:
            lines = f.readlines()

        # 返回最近的记录
        history = []
        for line in lines[-limit:]:
            try:
                history.append(json.loads(line))
            except:
                pass

        return history


def main():
    """示例用法"""

    # 创建发布器
    publisher = XiaohongshuPublisher()

    # 检查服务
    print("🔍 检查MCP服务...")
    if not publisher.check_service():
        print("❌ MCP服务未运行，请先启动服务")
        return

    print("✅ MCP服务正常\n")

    # 示例：发布单篇笔记
    print("📝 发布测试笔记...")

    result = publisher.publish(
        title="🦞测试发布",
        content="这是测试内容\n\n多行文本测试",
        images=[
            "/home/gem/.openclaw/workspace/lobster-journey/branding/intro-images/intro-01.jpg"
        ]
    )

    if result['success']:
        print(f"✅ 发布成功: {result['data']['title']}")
        print(f"   图片数量: {result['data']['images']}")
        print(f"   状态: {result['data']['status']}")
    else:
        print(f"❌ 发布失败: {result['message']}")

    # 查看发布历史
    print("\n📊 发布历史（最近5条）:")
    history = publisher.get_publish_history(limit=5)
    for i, entry in enumerate(history, 1):
        status = "✅" if entry['success'] else "❌"
        print(f"   {i}. {status} {entry['title']} ({entry['timestamp']})")


if __name__ == "__main__":
    main()
