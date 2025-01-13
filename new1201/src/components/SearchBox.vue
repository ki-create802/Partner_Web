<template>
  <div class="search-container">
    <input
      type="text"
      v-model="searchQuery"
      @input="handleInput"
      class="search-input"
      placeholder="请输入搜索内容"
    />
    <button class="search-button" @click="handleSearch">搜索</button>
  </div>
</template>

<script>
export default {
  name: 'SearchBox',
  props: {
    initialQuery: {
      type: String,
      default: ''
    },
    initialScope: {
      type: String,
      default: '全部'
    }
  },
  data() {
    return {
      searchQuery: this.initialQuery,
      searchScope: this.initialScope
    };
  },
  methods: {
    handleInput() {
      // 当输入框内容变化时触发
      if(this.searchQuery==null)return;
      this.$emit('input', this.searchQuery);
    },
    handleSearch() {
      // 检查 searchQuery 是否为空
      if (!this.searchQuery.trim()) {
        alert('搜索框不能为空！'); // 提示用户输入内容
        return; // 如果为空，不执行搜索操作
      }
      // 当点击搜索按钮时触发
      this.$emit('search', { query: this.searchQuery, scope: this.searchScope });
    }
  }
};
</script>

<style scoped>
/* 搜索框容器 */
.search-container {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  max-width: 600px; /* 最大宽度 */
  margin: 0 auto; /* 居中 */
  border: 2px solid #ccc;
  border-radius: 25px; /* 圆角 */
  overflow: hidden; /* 防止内容溢出 */
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 阴影 */
}

/* 输入框样式 */
.search-input {
  flex: 1; /* 占据剩余空间 */
  padding: 10px 15px;
  font-size: 16px;
  border: none;
  outline: none; /* 去掉输入框的默认边框 */
  background-color: transparent; /* 透明背景 */
}

/* 搜索按钮样式 */
.search-button {
  padding: 10px 20px;
  font-size: 16px;
  color: #fff;
  /*background-color: #4285f4; /* Google 蓝色 */
  background-color: rgb(110, 176, 110); 
  border: none;
  border-radius: 0 25px 25px 0; /* 圆角 */
  cursor: pointer;
  transition: background-color 0.3s ease; /* 过渡效果 */
}

/* 按钮悬停效果 */
.search-button:hover {
  background-color: #357ae8; /* 深一点的蓝色 */
}

/* 输入框聚焦时的效果 */
.search-input:focus {
  border-color: #4285f4; /* 聚焦时边框颜色变化 */
}
</style>