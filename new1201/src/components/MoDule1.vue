<template>
    <!-- 模块 -->
    <div class="module">
        <div class="card-group">
            <div class="card">
                <img src="@/assets/car.jpg" alt="Image" class="card-img">
                <div class="card-body">
                    <h2 class="card-title">拼车</h2>
                    <button class="bbutton" @click="handleSearch({ query: '', scope: '拼车' })">Go</button>
                </div>
            </div>

            <div class="card">
                <img src="@/assets/sports.jpg" alt="Image" class="card-img">
                <div class="card-body">
                    <h2 class="card-title">运动</h2>
                    <button class="bbutton" @click="handleSearch({ query: '', scope: '运动' })">Go</button>
                </div>
            </div>

            <div class="card">
                <img src="@/assets/play.jpg" alt="Image" class="card-img">
                <div class="card-body">
                    <h2 class="card-title">娱乐</h2>
                    <button class="bbutton" @click="handleSearch({ query: '', scope: '娱乐' })">Go</button>
                </div>
            </div>
        </div>
        <div class="card-group mt-3">
            <div class="card">
                <img src="@/assets/shopping.jpg" alt="Image" class="card-img">
                <div class="card-body">
                    <h2 class="card-title">拼单</h2>
                    <button class="bbutton" @click="handleSearch({ query: '', scope: '拼单' })">Go</button>
                </div>
            </div>

            <div class="card">
                <img src="@/assets/photo.jpg" alt="Image" class="card-img">
                <div class="card-body">
                    <h2 class="card-title">约拍</h2>
                    <button class="bbutton" @click="handleSearch({ query: '', scope: '约拍' })">Go</button>
                </div>
            </div>

            <div class="card">
                <img src="@/assets/others.jpg" alt="Image" class="card-img">
                <div class="card-body">
                    <h2 class="card-title">其他</h2>
                    <button class="bbutton" @click="handleSearch({ query: '', scope: '其他' })">Go</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { search } from '@/api.js';

export default {
    name: 'MoDule',
    methods: {
        async handleSearch(searchParams) {     
            try {
                const { query, scope } = searchParams;
                // 向后端发送搜索请求
                const searchResults = await search(query, scope);
                // 导航到 FindPage 并传递搜索内容和结果
                this.$router.push({
                    name: 'FindPage',
                    query: {
                        q: query,
                        scope: scope,
                        results: JSON.stringify(searchResults) // 将搜索结果编码为 JSON 字符串
                }
                });
            } catch (error) {
                alert("搜索信息失败");
                console.error('搜索请求失败:', error);
            }
        }
    }
}
</script>

<style>
.module {
    margin-top: 20px;
}

.card-group {
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;
}

.card {
    max-width: 20rem;
    border: 1px solid #ddd;
    border-radius: 8px;
    overflow: hidden;
    margin: 0 10px; /* 添加左右间隔 */
}

.card:first-child {
    margin-left: 0; /* 第一个卡片左边无间隔 */
}

.card:last-child {
    margin-right: 0; /* 最后一个卡片右边无间隔 */
}

.card-img {
    width: 100%;
    height: 200px;
    object-fit: cover;
}

.card-body {
    padding: 15px;
    text-align: center;
}

.card-title {
    font-size: 1.25rem;
    margin-bottom: 10px;
}

.bbutton {
    background-color: rgb(234, 156, 187) !important;
    border: none;
    padding: 10px 20px;
    border-radius: 5px;
    cursor: pointer;
    color: white;
}

.mt-3 {
    margin-top: 1rem;
}
</style>