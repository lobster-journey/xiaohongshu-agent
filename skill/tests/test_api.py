#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
小红书Agent测试套件
Created by 🦞 Lobster Journey Studio
"""

import pytest
import requests
import time


class TestHealthCheck:
    """健康检查测试"""

    def test_health_endpoint(self):
        """测试健康检查接口"""
        response = requests.get("http://localhost:18060/health")
        assert response.status_code == 200

        data = response.json()
        assert data["success"] is True
        assert data["data"]["status"] == "healthy"


class TestPublishAPI:
    """发布API测试"""

    def test_publish_image(self):
        """测试图文发布"""
        payload = {
            "title": "测试标题",
            "content": "测试内容",
            "images": [],
            "tags": ["测试"],
            "draft": True
        }

        response = requests.post(
            "http://localhost:18060/api/v1/publish/image",
            json=payload
        )

        assert response.status_code == 200
        data = response.json()
        assert data["success"] is True


class TestSearchAPI:
    """搜索API测试"""

    def test_search(self):
        """测试搜索功能"""
        params = {
            "keyword": "AI",
            "limit": 10
        }

        response = requests.get(
            "http://localhost:18060/api/v1/search",
            params=params
        )

        assert response.status_code == 200
        data = response.json()
        assert data["success"] is True
        assert len(data["data"]["notes"]) > 0


class TestStatsAPI:
    """统计API测试"""

    def test_account_stats(self):
        """测试账号统计"""
        response = requests.get(
            "http://localhost:18060/api/v1/stats/account"
        )

        assert response.status_code == 200
        data = response.json()
        assert data["success"] is True

    def test_overview(self):
        """测试数据概览"""
        response = requests.get(
            "http://localhost:18060/api/v1/stats/overview"
        )

        assert response.status_code == 200
        data = response.json()
        assert data["success"] is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
