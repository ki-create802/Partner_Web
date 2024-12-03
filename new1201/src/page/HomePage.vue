<template>
  <GuideBar />
  <div class="mybody">
    <div class="left">
      <CarouselImage :images="images" />
      <SearchBox class="searchbox" @search="handleSearch" :initialScope="selectedScope" />
      <div class="sections">
        <button @click="handleSearch({ query: '', scope: '拼车' })">拼车</button>
        <button @click="handleSearch({ query: '', scope: '运动' })">运动</button>
        <button @click="handleSearch({ query: '', scope: '娱乐' })">娱乐</button>
        <button @click="handleSearch({ query: '', scope: '拼单' })">拼单</button>
        <button @click="handleSearch({ query: '', scope: '约拍' })">约拍</button>
        <button @click="handleSearch({ query: '', scope: '其他' })">其他</button>
      </div>
      <div class="hot">
        <h2>🔥🔥🔥火热</h2>
        <div class="hotlist">
          <FindListItem v-for="(item, index) in hotlist" :key="index" :item="item" />
        </div>
      </div>
    </div>
    <div class="right">
      <ACalendar class="calendar" />
    </div>
  </div>
</template>

<script>
import GuideBar from '@/components/GuideBar.vue';
import ACalendar from '@/components/ACalendar.vue';
import CarouselImage from '@/components/CarouselImage.vue';
import SearchBox from '@/components/SearchBox.vue';
import FindListItem from '@/components/FindListItem.vue';
import axios from 'axios';

export default {
  name: 'HomePage',
  components: {
    GuideBar,
    ACalendar,
    CarouselImage,
    SearchBox,
    FindListItem
  },
  data() {
    return {
      images: [
        require('@/assets/轮播1.jpg'),
        require('@/assets/轮播2.jpg'),
        require('@/assets/轮播3.jpg'),
        require('@/assets/轮播4.jpg'),
        require('@/assets/轮播5.jpg'),
        require('@/assets/轮播6.jpg')
      ],
      hotlist: [] // 初始化 hotlist 数组
    };
  },
  created() {
    this.fetchHotData(); // 在组件创建时获取热门数据
  },
  methods: {
    async handleSearch(searchParams) {
      try {
        const { query, scope } = searchParams;
        // 向后端发送搜索请求
        const response = await axios.get(`http://localhost:3000/api/search?Searchword=${query}&Range=${scope}`);
        const searchResults = response.data;
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
    },
    async fetchHotData() {
      try {
        // 向后端发送请求获取热门数据
        const response = await axios.get('http://localhost:3000/api/hot');
        this.hotlist = response.data;
      } catch (error) {
        console.error('获取热门数据失败:', error);
      }
    }
  }
}
</script>

<style scoped>
.mybody {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.left {
  flex: 0 0 60%; /* 占据 60% 的宽度 */
  display: flex;
  flex-direction: column;
  align-items: center; /* 水平居中 */
  padding: 10px; /* 添加一些内边距 */
}

.right {
  background-color: #f8feffbd;
  flex: 0 0 30%; /* 占据 40% 的宽度 */
  display: flex;
  flex-direction: column;
  align-items: center; /* 水平居中 */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  border-radius: 30px;
}

.calendar {
  margin-top:0%;
  max-width: 400px;
  width: 90%;
}

.searchbox {
  padding: 10px;
}

.sections {
  display: grid;
  grid-template-columns: repeat(3, 1fr); /* 三列 */
  grid-template-rows: repeat(2, auto); /* 两行 */
  gap: 40px; /* 列和行之间的间距 */
  width: 100%; /* 使网格容器占满父容器的宽度 */
  margin-top: 20px; /* 添加一些上边距 */
}

.sections button {
  width: 100%; /* 使按钮占满网格单元格的宽度 */
  height: 80px;
  padding: 10px;
  border: none;
  background-color: #ffffff76;
  cursor: pointer;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  border-radius: 20px;
}

.sections button:hover {
  background-color: #f4feffaa;
}

.hot {
  display: flex;
  flex-direction: column; /* 纵向排列 */
  align-items: flex-start; /* 靠左对齐 */
  width: 100%; /* 使容器占满父容器的宽度 */
  margin-top: 20px; /* 添加一些上边距 */
}

.hotlist {
  width: 100%;
}
</style>