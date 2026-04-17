#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
小红书视频发布脚本
Created by 🦞 Lobster Journey Studio
"""

import argparse
import json
import requests
from pathlib import Path


class XiaohongshuVideoPublisher:
    """小红书视频发布器"""

    def __init__(self, base_url="http://localhost:18060"):
        self.base_url = base_url
        self.api_url = f"{base_url}/api/v1"

    def publish_video(self, title, content, video, cover=None, tags=None, draft=False):
        """
        发布视频内容

        Args:
            title: 标题
            content: 内容
            video: 视频路径
            cover: 封面路径
            tags: 标签列表
            draft: 是否草稿

        Returns:
            发布结果
        """
        payload = {
            "title": title,
            "content": content,
            "video": video,
            "cover": cover,
            "tags": tags or [],
            "draft": draft
        }

        response = requests.post(
            f"{self.api_url}/publish/video",
            json=payload
        )

        return response.json()


def main():
    parser = argparse.ArgumentParser(
        description="🦞 小红书视频发布工具"
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
        "--video", "-v",
        required=True,
        help="视频路径"
    )
    parser.add_argument(
        "--cover", "-C",
        help="封面路径"
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
    publisher = XiaohongshuVideoPublisher(base_url=args.url)

    # 发布视频
    print(f"🎬 发布视频: {args.title}")
    result = publisher.publish_video(
        title=args.title,
        content=args.content,
        video=args.video,
        cover=args.cover,
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
