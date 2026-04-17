#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
小红书图文发布脚本
Created by 🦞 Lobster Journey Studio
"""

import argparse
import json
import requests
from pathlib import Path


class XiaohongshuPublisher:
    """小红书图文发布器"""

    def __init__(self, base_url="http://localhost:18060"):
        self.base_url = base_url
        self.api_url = f"{base_url}/api/v1"

    def publish_image(self, title, content, images, tags=None, draft=False):
        """
        发布图文内容

        Args:
            title: 标题
            content: 内容
            images: 图片路径列表
            tags: 标签列表
            draft: 是否草稿

        Returns:
            发布结果
        """
        payload = {
            "title": title,
            "content": content,
            "images": images,
            "tags": tags or [],
            "draft": draft
        }

        response = requests.post(
            f"{self.api_url}/publish/image",
            json=payload
        )

        return response.json()

    def check_status(self):
        """检查服务状态"""
        response = requests.get(f"{self.base_url}/health")
        return response.json()


def main():
    parser = argparse.ArgumentParser(
        description="🦞 小红书图文发布工具"
    )
    parser.add_argument(
        "--title", "-t",
        required=True,
        help="标题"
    )
    parser.add_argument(
        "--content", "-c",
        required=True,
        help="内容"
    )
    parser.add_argument(
        "--images", "-i",
        required=True,
        nargs="+",
        help="图片路径"
    )
    parser.add_argument(
        "--tags", "-g",
        nargs="+",
        default=[],
        help="标签"
    )
    parser.add_argument(
        "--draft", "-d",
        action="store_true",
        help="保存为草稿"
    )
    parser.add_argument(
        "--url", "-u",
        default="http://localhost:18060",
        help="MCP服务地址"
    )

    args = parser.parse_args()

    # 创建发布器
    publisher = XiaohongshuPublisher(base_url=args.url)

    # 检查服务状态
    print("🔍 检查服务状态...")
    status = publisher.check_status()
    if status.get("status") != "healthy":
        print("❌ MCP服务未运行，请先启动服务")
        return

    print(f"✅ 服务正常 (账号: {status.get('data', {}).get('account')})")

    # 发布内容
    print(f"\n📝 发布图文: {args.title}")
    result = publisher.publish_image(
        title=args.title,
        content=args.content,
        images=args.images,
        tags=args.tags,
        draft=args.draft
    )

    if result.get("success"):
        print(f"✅ 发布成功!")
        print(f"   Post ID: {result.get('data', {}).get('post_id')}")
    else:
        print(f"❌ 发布失败: {result.get('error')}")


if __name__ == "__main__":
    main()
