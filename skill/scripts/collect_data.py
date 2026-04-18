#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
小红书数据采集脚本
支持采集笔记数据、用户数据、互动数据等
Created by 🦞 Lobster Journey Studio
"""

import argparse
import json
import os
import time
from datetime import datetime, timedelta


class XiaohongshuDataCollector:
    """小红书数据采集器"""

    def __init__(self, output_dir=None):
        self.output_dir = output_dir or os.path.expanduser(
            "~/.openclaw/workspace/data/xiaohongshu"
        )
        os.makedirs(self.output_dir, exist_ok=True)

    def collect_note_data(self, note_id):
        """采集单篇笔记数据"""
        print(f"📊 采集笔记数据: {note_id}")

        # TODO: 实际实现需要调用MCP API或浏览器自动化
        # 这里是模拟数据
        data = {
            "note_id": note_id,
            "title": "示例笔记标题",
            "content": "示例内容...",
            "author": {
                "user_id": "user123",
                "nickname": "示例作者",
                "avatar": "https://..."
            },
            "stats": {
                "views": 1000,
                "likes": 100,
                "comments": 20,
                "collects": 50,
                "shares": 10
            },
            "publish_time": "2026-04-19 10:00:00",
            "collect_time": datetime.now().isoformat()
        }

        # 保存数据
        output_file = os.path.join(self.output_dir, f"note_{note_id}.json")
        with open(output_file, "w", encoding="utf-8") as f:
            json.dump(data, f, ensure_ascii=False, indent=2)

        print(f"✅ 数据已保存: {output_file}")
        return data

    def collect_user_data(self, user_id):
        """采集用户数据"""
        print(f"📊 采集用户数据: {user_id}")

        # TODO: 实际实现
        data = {
            "user_id": user_id,
            "nickname": "示例用户",
            "avatar": "https://...",
            "bio": "个人简介",
            "stats": {
                "followers": 1000,
                "following": 100,
                "notes": 50
            },
            "collect_time": datetime.now().isoformat()
        }

        output_file = os.path.join(self.output_dir, f"user_{user_id}.json")
        with open(output_file, "w", encoding="utf-8") as f:
            json.dump(data, f, ensure_ascii=False, indent=2)

        print(f"✅ 数据已保存: {output_file}")
        return data

    def collect_search_results(self, keyword, limit=50):
        """采集搜索结果数据"""
        print(f"📊 采集搜索结果: {keyword} (最多{limit}条)")

        # TODO: 实际实现
        results = []
        for i in range(min(limit, 10)):  # 模拟10条数据
            results.append({
                "note_id": f"note_{i}",
                "title": f"示例笔记 {i+1}",
                "author": f"作者{i+1}",
                "likes": 100 + i * 10,
                "comments": 10 + i
            })

        data = {
            "keyword": keyword,
            "total": len(results),
            "results": results,
            "collect_time": datetime.now().isoformat()
        }

        output_file = os.path.join(
            self.output_dir,
            f"search_{keyword}_{datetime.now().strftime('%Y%m%d_%H%M%S')}.json"
        )
        with open(output_file, "w", encoding="utf-8") as f:
            json.dump(data, f, ensure_ascii=False, indent=2)

        print(f"✅ 数据已保存: {output_file}")
        return data

    def collect_daily_stats(self, note_ids):
        """采集每日统计数据"""
        print(f"📊 采集每日统计数据: {len(note_ids)}篇笔记")

        stats = []
        for i, note_id in enumerate(note_ids, 1):
            print(f"  [{i}/{len(note_ids)}] 采集: {note_id}")

            data = self.collect_note_data(note_id)
            stats.append({
                "note_id": note_id,
                "title": data["title"],
                "views": data["stats"]["views"],
                "likes": data["stats"]["likes"],
                "comments": data["stats"]["comments"],
                "collects": data["stats"]["collects"],
                "shares": data["stats"]["shares"]
            })

            time.sleep(1)  # 避免频繁请求

        # 汇总统计
        summary = {
            "date": datetime.now().strftime("%Y-%m-%d"),
            "total_notes": len(stats),
            "total_views": sum(s["views"] for s in stats),
            "total_likes": sum(s["likes"] for s in stats),
            "total_comments": sum(s["comments"] for s in stats),
            "total_collects": sum(s["collects"] for s in stats),
            "total_shares": sum(s["shares"] for s in stats),
            "details": stats,
            "collect_time": datetime.now().isoformat()
        }

        output_file = os.path.join(
            self.output_dir,
            f"daily_stats_{datetime.now().strftime('%Y%m%d')}.json"
        )
        with open(output_file, "w", encoding="utf-8") as f:
            json.dump(summary, f, ensure_ascii=False, indent=2)

        print(f"\n✅ 统计数据已保存: {output_file}")
        return summary

    def export_to_csv(self, data_type="all"):
        """导出数据为CSV格式"""
        print(f"📊 导出数据: {data_type}")

        import csv

        # 查找所有JSON文件
        json_files = []
        for filename in os.listdir(self.output_dir):
            if filename.endswith(".json"):
                if data_type == "all" or data_type in filename:
                    json_files.append(os.path.join(self.output_dir, filename))

        if not json_files:
            print("❌ 未找到数据文件")
            return

        print(f"   找到 {len(json_files)} 个数据文件")

        # 导出CSV
        csv_file = os.path.join(
            self.output_dir,
            f"export_{datetime.now().strftime('%Y%m%d_%H%M%S')}.csv"
        )

        with open(csv_file, "w", newline="", encoding="utf-8") as f:
            writer = csv.writer(f)

            # 写入表头
            writer.writerow(["文件名", "采集时间"])

            for json_file in json_files:
                with open(json_file, "r", encoding="utf-8") as jf:
                    data = json.load(jf)
                    writer.writerow([
                        os.path.basename(json_file),
                        data.get("collect_time", "N/A")
                    ])

        print(f"✅ CSV已导出: {csv_file}")


def main():
    parser = argparse.ArgumentParser(
        description="🦞 小红书数据采集工具",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
示例:
  # 采集单篇笔记数据
  python collect_data.py --note-id abc123

  # 采集用户数据
  python collect_data.py --user-id user456

  # 采集搜索结果
  python collect_data.py --keyword "AI人工智能" --limit 50

  # 采集每日统计（批量）
  python collect_data.py --batch-file notes.txt --type daily-stats

  # 导出CSV
  python collect_data.py --export-csv
        """
    )

    parser.add_argument(
        "--note-id", "-n",
        help="笔记ID"
    )
    parser.add_argument(
        "--user-id", "-u",
        help="用户ID"
    )
    parser.add_argument(
        "--keyword", "-k",
        help="搜索关键词"
    )
    parser.add_argument(
        "--limit", "-l",
        type=int,
        default=50,
        help="数量限制"
    )
    parser.add_argument(
        "--batch-file", "-b",
        help="批量操作文件（每行一个笔记ID）"
    )
    parser.add_argument(
        "--type", "-t",
        choices=["note", "user", "search", "daily-stats"],
        default="note",
        help="采集类型"
    )
    parser.add_argument(
        "--output-dir", "-o",
        help="输出目录"
    )
    parser.add_argument(
        "--export-csv",
        action="store_true",
        help="导出CSV格式"
    )
    parser.add_argument(
        "--export-type",
        default="all",
        help="导出类型（all/note/user/search）"
    )

    args = parser.parse_args()

    # 创建采集器
    collector = XiaohongshuDataCollector(output_dir=args.output_dir)

    # 执行操作
    if args.export_csv:
        # 导出CSV
        collector.export_to_csv(data_type=args.export_type)

    elif args.batch_file and args.type == "daily-stats":
        # 批量采集每日统计
        with open(args.batch_file, "r") as f:
            note_ids = [line.strip() for line in f if line.strip()]

        result = collector.collect_daily_stats(note_ids)

        print("\n📊 统计汇总:")
        print(f"   总笔记数: {result['total_notes']}")
        print(f"   总阅读量: {result['total_views']}")
        print(f"   总点赞数: {result['total_likes']}")
        print(f"   总评论数: {result['total_comments']}")
        print(f"   总收藏数: {result['total_collects']}")
        print(f"   总转发数: {result['total_shares']}")

    elif args.keyword:
        # 采集搜索结果
        result = collector.collect_search_results(args.keyword, args.limit)
        print(f"\n✅ 采集完成: {result['total']}条结果")

    elif args.note_id:
        # 采集笔记数据
        collector.collect_note_data(args.note_id)

    elif args.user_id:
        # 采集用户数据
        collector.collect_user_data(args.user_id)

    else:
        parser.print_help()


if __name__ == "__main__":
    main()
