<template>
  <body>
    <TopBar />
    <div class="searchContainer">
      <SearchBox class="searchBox" :initialQuery="searchQuery" :initialScope="searchScope" @search="handleSearch" />
    </div>
    <div class="basicSelect">
      <button @click="handleUserSearch">用户</button>
      <button @click="handleSearch({ query: searchQuery, scope: '拼车' })" :class="{ active: selectedScope === '拼车' }">拼车</button>
      <button @click="handleSearch({ query: searchQuery, scope: '运动' })" :class="{ active: selectedScope === '运动' }">运动</button>
      <button @click="handleSearch({ query: searchQuery, scope: '娱乐' })" :class="{ active: selectedScope === '娱乐' }">娱乐</button>
      <button @click="handleSearch({ query: searchQuery, scope: '拼单' })" :class="{ active: selectedScope === '拼单' }">拼单</button>
      <button @click="handleSearch({ query: searchQuery, scope: '约拍' })" :class="{ active: selectedScope === '约拍' }">约拍</button>
      <button @click="handleSearch({ query: searchQuery, scope: '其他' })" :class="{ active: selectedScope === '其他' }">其他</button>
    </div>
    <div v-if="searchUser">
      <div v-if="userResults.length > 0" class="userResults">
        <UserListItem v-for="(user, index) in userResults" :key="index" :user="user" />
      </div>
      <div v-else class="noResults">
        <p>没有找到相关用户</p>
      </div>
    </div>
    <div v-else>
      <div v-if="searchResults.length > 0" class="searchResults">
        <FindListItem v-for="(result, index) in searchResults" :key="index" :item="result" />
      </div>
      <div v-else class="noResults">
        <p>没有找到相关结果</p>
      </div>
    </div>
  </body>
</template>

<script>
import TopBar from '../components/TopBar.vue';
import SearchBox from '../components/SearchBox.vue';
import FindListItem from '../components/FindListItem.vue';
import UserListItem from '../components/UserListItem.vue'; // 引入 UserListItem 组件
import axios from 'axios';

export default {
  name: 'FindPage',
  components: {
    TopBar,
    SearchBox,
    FindListItem,
    UserListItem, // 注册 UserListItem 组件
  },
  data() {
    return {
      searchQuery: this.$route.query.q || '',
      searchScope: this.$route.query.scope || '全部',
      searchResults: [],
      userResults: [], // 存储用户搜索结果
      selectedScope: '',
      searchUser: false,
    };
  },
  created() {
    // 从 query 参数中获取搜索结果
    if (this.$route.query.results) {
      this.searchResults = JSON.parse(this.$route.query.results);
    }
    // 如果没有路由传参，执行 handleSearch
    if (!this.searchQuery && this.searchResults.length === 0) {
      this.handleSearch({ query: '', scope: '全部' });
    }
  },
  methods: {
    async handleSearch(searchParams) {
      try {
        const { query, scope } = searchParams;
        // 向后端发送搜索请求
        const response = await axios.get(`http://localhost:3000/api/search?Searchword=${query}&Range=${scope}`);
        this.searchResults = response.data;
        this.searchQuery = query;
        this.searchScope = scope;
        this.selectedScope = scope; // 更新选中的范围
        this.searchUser = false; // 重置用户搜索状态
      } catch (error) {
        console.error('搜索请求失败:', error);
      }
    },
    async handleUserSearch() {
      try {
        // 向后端发送用户搜索请求
        const response = await axios.get(`http://localhost:3000/api/users`); // 假设后端有 /api/users 接口
        this.userResults = response.data;
        this.searchUser = true; // 设置用户搜索状态
        console.log(this.userResults); // 使用 console.log 替换 alert
      } catch (error) {
        console.error('用户搜索请求失败:', error);
      }
    }
  },
}
</script>

<style scoped>
.searchContainer {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}

.basicSelect {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  gap: 20px;
}

.basicSelect button {
  background-color: #cceaff;
  border: none;
  border-radius: 4px;
  height: 35px;
  width: 60px;
  font-size: 17px;
  color: #657c8c;
  cursor: pointer;
}

.basicSelect button:hover {
  background-color: #bccfff;
  color: #ffffff;
}

.basicSelect button.active {
  background-color: #bccfff;
  color: #ffffff;
  cursor: not-allowed;
}
</style>