<template>
    <!-- 模块 -->
    <div class="module">
        <div>
        <b-card-group deck>
            <b-card
            title="拼车"
            :img-src="require('@/assets/car.jpg')"
            img-alt="Image"
            img-top
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
            >
            <!-- <b-card-text>
                1
            </b-card-text> -->
            <b-button class="bbutton" @click="handleSearch({ query: '', scope: '拼车' })">Go</b-button>
            </b-card>

            <b-card
            title="运动"
            :img-src="require('@/assets/sports.jpg')"
            img-alt="Image"
            img-top
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
            >
            <!-- <b-card-text>
                2
            </b-card-text> -->
            <b-button class="bbutton" @click="handleSearch({ query: '', scope: '运动' })">Go</b-button>
            </b-card>

            <b-card
            title="娱乐"
            :img-src="require('@/assets/play.jpg')"
            img-alt="Image"
            img-top
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
            >
            <!-- <b-card-text>
                3
            </b-card-text> -->
            <b-button class="bbutton" @click="handleSearch({ query: '', scope: '娱乐' })">Go</b-button>
            </b-card>
        </b-card-group>
        </div>
        <div class="mt-3">
        <b-card-group deck>
            <b-card
            title="拼单"
            :img-src="require('@/assets/shopping.jpg')"
            img-alt="Image"
            img-top
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
            >
            <!-- <b-card-text>
                4
            </b-card-text> -->
            <b-button class="bbutton" @click="handleSearch({ query: '', scope: '拼单' })">Go</b-button>
            </b-card>

            <b-card
            title="约拍"
            :img-src="require('@/assets/photo.jpg')"
            img-alt="Image"
            img-top
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
            >
            <!-- <b-card-text>
                5
            </b-card-text> -->
            <b-button class="bbutton" @click="handleSearch({ query: '', scope: '约拍' })">Go</b-button>
            </b-card>

            <b-card
            title="其他"
            :img-src="require('@/assets/others.jpg')"
            img-alt="Image"
            img-top
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
            >
            <!-- <b-card-text>
                4
            </b-card-text> -->
            <b-button class="bbutton" @click="handleSearch({ query: '', scope: '其他' })">Go</b-button>
            </b-card>
        </b-card-group>
        </div>
    </div>

</template>

<script>
//import axios from 'axios';
import { search } from '@/api.js';

export default{
    name: 'MoDule',
    methods: {
        async handleSearch(searchParams) {
            const response = await search(searchParams);
            alert(JSON.stringify(response));
            // try {
            //     const { query, scope } = searchParams;
            //     // 向后端发送搜索请求
            //     const response = await axios.get(`http://localhost:3000/api/search`);///
            //     const searchResults = response.data;
            //     // 导航到 FindPage 并传递搜索内容和结果
            //     this.$router.push({
            //         name: 'FindPage',
            //         query: {
            //             q: query,
            //             scope: scope,
            //             results: JSON.stringify(searchResults) // 将搜索结果编码为 JSON 字符串
            //     }
            //     });
            // } catch (error) {
            //     alert("搜索信息失败");
            //     console.error('搜索请求失败:', error);
            // }

        }
    }
}

</script>

<style>
.module{
  margin-top: 20px;
}

.bbutton{
    background-color: rgb(234, 156, 187) !important; /* 背景色 */
}

.module img {
  height: 200px; /* 固定高度 */
  object-fit: cover; /* 保持图片比例，裁剪多余部分 */
}

</style>


