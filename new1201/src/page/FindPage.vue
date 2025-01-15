<template>
  <body>
    <GuideBar />
    <div class="searchContainer">
      <SearchBox class="searchBox" :initialQuery="searchQuery" :initialScope="searchScope" @search="handleSearch" @input="inputChange" />
    </div>
    <div class="basicSelect">
      <button @click="handleUserSearch" :class="{ active: searchUser }">用户</button>
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
        <div class="hot">
          <div class="hotlist">
            <FindListItem v-for="(result, index) in searchResults" :key="index" :item="result" class="hotitem" />
          </div>
        </div>
      </div>
      <div v-else class="noResults">
        <p>没有找到相关结果</p>
      </div>
    </div>
  </body>
</template>

<script>
import GuideBar from '../components/GuideBar.vue';
import SearchBox from '../components/SearchBox.vue';
import FindListItem from '../components/FindListItem.vue';
import UserListItem from '../components/UserListItem.vue'; // 引入 UserListItem 组件
import {search,searchUsers} from '@/api.js'

export default {
  name: 'FindPage',
  components: {
    GuideBar,
    SearchBox,
    FindListItem,
    UserListItem, // 注册 UserListItem 组件
  },
  data() {
    return {
      searchQuery: this.$route.query.q ,
      searchScope: this.$route.query.scope,
      //searchQuery: this.$route.query.q || '',
      //searchScope: this.$route.query.scope || '全部',
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
    if (!this.searchScope && this.searchResults.length === 0) {
      this.handleSearch({ query: '', scope: '全部' });
    }
    else{
      this.selectedScope=this.searchScope;
    }
  },
  methods: {
    inputChange(input){
      this.searchQuery=input;
    },
    async handleSearch(searchParams) {
      try {
        const data=await search(searchParams.query,searchParams.scope);
        this.searchResults=data;
        this.searchQuery = searchParams.query;
        this.searchScope = searchParams.scope;
        this.selectedScope = searchParams.scope; // 更新选中的范围
        this.searchUser=false;
        this.searchUser = false; // 重置用户搜索状态
      } catch (error) {
        console.error('搜索请求失败:', error);
      }
    },
    async handleUserSearch() {
      this.searchUser = true; // 设置用户搜索状态
      this.selectedScope="";
      try {
        this.userResults = await searchUsers(this.searchQuery);
      } catch (error) {
        alert("请求用户信息失败");
        console.error('用户搜索请求失败:', error);
      }
    }
  },
}
</script>

<style scoped>
.searchContainer {
  margin-top: 130px;
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

/* 新增样式，与 HomePage 的 .hot 和 .hotlist 一致 */
.hot {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin-top: 20px;
  width: 70%; /* 宽度设置为 80% */
  margin: 0 auto; /* 水平居中 */
}

.hotlist {
  width: 100%;
}

.hotitem {
  margin-bottom: 20px;
}

.noResults {
  text-align: center;
  margin-top: 20px;
  font-size: 1.2em;
  color: #666;
}
</style>