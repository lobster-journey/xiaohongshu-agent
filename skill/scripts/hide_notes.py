#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
小红书笔记隐藏工具
使用Playwright自动化隐藏笔记
Created by 🦞 Lobster Journey Studio
"""

import asyncio
import json
import time
from pathlib import Path
from playwright.async_api import async_playwright


class NoteHider:
    """笔记隐藏器"""

    def __init__(self, cookies_file: str = None):
        self.cookies_file = cookies_file or "/home/gem/.openclaw/mcp/cookies.json"
        self.browser = None
        self.context = None
        self.page = None

    async def init_browser(self):
        """初始化浏览器"""
        print("🚀 启动浏览器...")

        playwright = await async_playwright().start()

        # 使用持久化上下文，保存登录状态
        user_data_dir = "/home/gem/.openclaw/workspace/xiaohongshu-browser-data"

        self.browser = await playwright.chromium.launch_persistent_context(
            user_data_dir,
            headless=False,  # 显示浏览器，方便观察
            viewport={"width": 1280, "height": 800},
            locale="zh-CN",
            timezone_id="Asia/Shanghai"
        )

        self.page = await self.browser.new_page()
        print("✅ 浏览器启动成功")

    async def load_cookies(self):
        """加载已保存的cookies"""
        print("🍪 加载登录状态...")

        cookies_path = Path(self.cookies_file)
        if not cookies_path.exists():
            print("⚠️  未找到cookies文件，需要手动登录")
            return False

        with open(cookies_path, 'r') as f:
            cookies = json.load(f)

        await self.context.add_cookies(cookies)
        print("✅ 登录状态加载成功")
        return True

    async def visit_my_notes(self):
        """访问我的笔记页面"""
        print("\n📱 访问创作者中心...")

        # 访问创作者中心
        await self.page.goto("https://creator.xiaohongshu.com/")
        await asyncio.sleep(2)

        # 检查是否登录
        current_url = self.page.url
        if "login" in current_url:
            print("⚠️  需要登录，请在浏览器中手动登录")
            print("等待登录完成...")
            await asyncio.sleep(30)  # 等待30秒让用户登录

        print("✅ 已进入创作者中心")

    async def get_note_list(self):
        """获取笔记列表"""
        print("\n📋 获取笔记列表...")

        # 点击"笔记管理"
        try:
            await self.page.click('text=笔记管理')
            await asyncio.sleep(2)
        except:
            pass

        # 获取所有笔记
        notes = []
        note_elements = await self.page.query_selector_all('.note-item')

        print(f"📊 找到 {len(note_elements)} 篇笔记")

        for i, elem in enumerate(note_elements, 1):
            try:
                # 获取笔记标题
                title_elem = await elem.query_selector('.title')
                title = await title_elem.inner_text() if title_elem else f"笔记{i}"

                # 获取笔记ID（从链接中提取）
                link_elem = await elem.query_selector('a')
                href = await link_elem.get_attribute('href') if link_elem else ""
                note_id = href.split('/')[-1] if href else ""

                notes.append({
                    'id': note_id,
                    'title': title,
                    'element': elem
                })

                print(f"  {i}. {title} (ID: {note_id})")

            except Exception as e:
                print(f"  ⚠️  解析笔记{i}失败: {e}")

        return notes

    async def hide_note(self, note_element):
        """隐藏单篇笔记"""
        try:
            # 点击"更多"按钮
            more_btn = await note_element.query_selector('.more-btn, [class*="more"]')
            if not more_btn:
                # 尝试其他选择器
                more_btn = await note_element.query_selector('button:has-text("更多")')

            if more_btn:
                await more_btn.click()
                await asyncio.sleep(1)

            # 点击"设置为仅自己可见"
            hide_option = await self.page.query_selector('text=设置为仅自己可见')
            if not hide_option:
                hide_option = await self.page.query_selector('text=仅自己可见')

            if hide_option:
                await hide_option.click()
                await asyncio.sleep(1)

                # 确认
                confirm_btn = await self.page.query_selector('button:has-text("确定")')
                if confirm_btn:
                    await confirm_btn.click()
                    await asyncio.sleep(1)

                print("  ✅ 隐藏成功")
                return True
            else:
                print("  ⚠️  未找到隐藏选项")
                return False

        except Exception as e:
            print(f"  ❌ 隐藏失败: {e}")
            return False

    async def batch_hide_notes(self, limit=11):
        """批量隐藏笔记"""
        print(f"\n🔒 开始批量隐藏前 {limit} 篇笔记...")

        notes = await self.get_note_list()

        success_count = 0
        for i, note in enumerate(notes[:limit], 1):
            print(f"\n[{i}/{min(len(notes), limit)}] 处理: {note['title']}")
            if await self.hide_note(note['element']):
                success_count += 1

            # 避免操作过快
            if i < len(notes):
                await asyncio.sleep(3)

        print(f"\n✅ 批量隐藏完成: 成功 {success_count} 篇，失败 {limit - success_count} 篇")
        return success_count

    async def close(self):
        """关闭浏览器"""
        if self.browser:
            await self.browser.close()


async def main():
    """主函数"""
    print("=" * 60)
    print("🦞 小红书笔记隐藏工具")
    print("=" * 60)

    hider = NoteHider()

    try:
        # 初始化浏览器
        await hider.init_browser()

        # 访问创作者中心
        await hider.visit_my_notes()

        # 等待页面加载
        await asyncio.sleep(3)

        # 批量隐藏笔记
        await hider.batch_hide_notes(limit=11)

    except Exception as e:
        print(f"\n❌ 发生错误: {e}")
        import traceback
        traceback.print_exc()

    finally:
        print("\n按回车键关闭浏览器...")
        input()
        await hider.close()


if __name__ == "__main__":
    asyncio.run(main())
