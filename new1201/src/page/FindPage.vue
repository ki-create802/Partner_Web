<template>
    <body>
      <GuideBar />
      <div class="searchContainer">
        <SearchBox class="searchBox" :initialQuery="searchQuery" @search="handleSearch"/>
      </div>
      <div class="basicSelect">
        <button>拼车</button>
        <button>运动</button>
        <button>娱乐</button>
        <button>拼单</button>
        <button>约拍</button>
        <button>其他</button>
      </div>
      <div v-if="searchResults.length > 0" class="searchResults">
        <FindListItem v-for="(result, index) in searchResults" :key="index" :item="result" />
      </div>
      <div v-else class="noResults">
        <p>没有找到相关结果</p>
      </div>
    </body>
  </template>
  
  <script>
  import GuideBar from '../components/GuideBar.vue';
  import SearchBox from '../components/SearchBox.vue';
  import FindListItem from '../components/FindListItem.vue';
  import axios from 'axios';
  export default {
  name: 'FindPage',
  components: {
    GuideBar,
    SearchBox,
    FindListItem,
  },
  data() {
    return {
        searchQuery: this.$route.query.q || '',
        searchResults: []
    };
  },
  created() {
    // 从 query 参数中获取搜索结果
    if (this.$route.query.results) {
        this.searchResults = JSON.parse(this.$route.query.results);
    }
    // 如果没有路由传参，执行 handleSearch
    if (!this.searchQuery && this.searchResults.length === 0) {
      this.handleSearch('');
    }
  },
  methods: {
    async handleSearch(query) {
      try {
        // 向后端发送搜索请求
        const response = await axios.get(`http://localhost:3000/api/search?q=${query}`);
        this.searchResults = response.data;
        this.searchQuery = query;
      } catch (error) {
        console.error('搜索请求失败:', error);
      }
    }
  },
}  </script>
    


<style scoped>
.searchContainer{
    margin-top: 20px;
    display: flex;
    justify-content: center;
}
.basicSelect {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  gap:20px;
}

.basicSelect button {
  background-color: #cceaff;
  border: none;
  border-radius: 4px;
  height: 35px;
  width: 60px;
  font-size: 17px;
  color: #657c8c;
}
.basicSelect button:hover{
    background-color: #657c8c;
    color: #ffffff;
}


</style>