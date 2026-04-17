#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
小红书数据统计脚本
Created by 🦞 Lobster Journey Studio
"""

import argparse
import json
import requests
from datetime import datetime, timedelta


class XiaohongshuStats:
    """小红书数据统计器"""

    def __init__(self, base_url="http://localhost:18060"):
        self.base_url = base_url
        self.api_url = f"{base_url}/api/v1"

    def get_account_stats(self):
        """获取账号统计"""
        response = requests.get(f"{self.api_url}/stats/account")
        return response.json()

    def get_post_stats(self, post_id):
        """获取单篇内容统计"""
        response = requests.get(f"{self.api_url}/stats/post/{post_id}")
        return response.json()

    def get_recent_stats(self, days=7):
        """获取近期统计"""
        end_date = datetime.now()
        start_date = end_date - timedelta(days=days)

        params = {
            "start_date": start_date.strftime("%Y-%m-%d"),
            "end_date": end_date.strftime("%Y-%m-%d")
        }

        response = requests.get(
            f"{self.api_url}/stats/recent",
            params=params
        )

        return response.json()

    def get_overview(self):
        """获取概览数据"""
        response = requests.get(f"{self.api_url}/stats/overview")
        return response.json()


def main():
    parser = argparse.ArgumentParser(
        description="🦞 小红书数据统计工具"
    )
    parser.add_argument(
        "--type", "-t",
        choices=["account", "post", "recent", "overview"],
        default="overview",
        help="统计类型"
    )
    parser.add_argument(
        "--post-id", "-p",
        help="内容ID (type=post时必需)"
    )
    parser.add_argument(
        "--days", "-d",
        type=int,
        default=7,
        help="统计天数 (type=recent时有效)"
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

    # 创建统计器
    stats = XiaohongshuStats(base_url=args.url)

    # 执行统计
    if args.type == "account":
        print("📊 账号统计")
        result = stats.get_account_stats()

    elif args.type == "post":
        if not args.post_id:
            print("❌ 请提供 --post-id 参数")
            return
        print(f"📊 内容统计: {args.post_id}")
        result = stats.get_post_stats(args.post_id)

    elif args.type == "recent":
        print(f"📊 近{args.days}天统计")
        result = stats.get_recent_stats(days=args.days)

    else:  # overview
        print("📊 数据概览")
        result = stats.get_overview()

    # 输出结果
    if result.get("success"):
        data = result.get("data", {})

        if args.output == "json":
            print(json.dumps(result, indent=2, ensure_ascii=False))
        else:
            print("\n📈 统计结果:")
            for key, value in data.items():
                print(f"   {key}: {value}")
    else:
        print(f"❌ 统计失败: {result.get('error')}")


if __name__ == "__main__":
    main()
