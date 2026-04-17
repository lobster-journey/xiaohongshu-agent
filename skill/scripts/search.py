#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
小红书内容搜索脚本
Created by 🦞 Lobster Journey Studio
"""

import argparse
import json
import requests


class XiaohongshuSearcher:
    """小红书内容搜索器"""

    def __init__(self, base_url="http://localhost:18060"):
        self.base_url = base_url
        self.api_url = f"{base_url}/api/v1"

    def search(self, keyword, limit=20, sort="general"):
        """
        搜索内容

        Args:
            keyword: 关键词
            limit: 数量限制
            sort: 排序方式 (general/hot/newest)

        Returns:
            搜索结果
        """
        params = {
            "keyword": keyword,
            "limit": limit,
            "sort": sort
        }

        response = requests.get(
            f"{self.api_url}/search",
            params=params
        )

        return response.json()

    def get_notes(self, keyword, limit=50):
        """获取笔记列表"""
        return self.search(keyword, limit=limit, sort="general")

    def get_hot_notes(self, keyword, limit=20):
        """获取热门笔记"""
        return self.search(keyword, limit=limit, sort="hot")


def main():
    parser = argparse.ArgumentParser(
        description="🦞 小红书内容搜索工具"
    )
    parser.add_argument(
        "keyword",
        help="搜索关键词"
    )
    parser.add_argument(
        "--limit", "-l",
        type=int,
        default=20,
        help="数量限制"
    )
    parser.add_argument(
        "--sort", "-s",
        choices=["general", "hot", "newest"],
        default="general",
        help="排序方式"
    )
    parser.add_argument(
        "--output", "-o",
        choices=["json", "table"],
        default="table",
        help="输出格式"
    )
    parser.add_argument(
        "--url", "-u",
        default="http://localhost:18060",
        help="MCP服务地址"
    )

    args = parser.parse_args()

    # 创建搜索器
    searcher = XiaohongshuSearcher(base_url=args.url)

    # 执行搜索
    print(f"🔍 搜索: {args.keyword}")
    result = searcher.search(
        keyword=args.keyword,
        limit=args.limit,
        sort=args.sort
    )

    if result.get("success"):
        notes = result.get("data", {}).get("notes", [])
        print(f"✅ 找到 {len(notes)} 条结果\n")

        if args.output == "json":
            print(json.dumps(result, indent=2, ensure_ascii=False))
        else:
            for i, note in enumerate(notes, 1):
                print(f"{i}. {note.get('title', 'N/A')}")
                print(f"   点赞: {note.get('likes', 0)} | 评论: {note.get('comments', 0)}")
                print(f"   作者: {note.get('author', {}).get('nickname', 'N/A')}")
                print()
    else:
        print(f"❌ 搜索失败: {result.get('error')}")


if __name__ == "__main__":
    main()
