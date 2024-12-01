<!-- src/components/SearchBox.vue -->
<template>
    <div class="search-box">
      <input
        type="text"
        v-model="searchQuery"
        @input="handleInput"
        placeholder="请输入搜索内容"
      />
      <button @click="handleSearch">搜索</button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'SearchBox',
    data() {
      return {
        searchQuery: ''
      };
    },
    methods: {
      handleInput() {
        // 当输入框内容变化时触发
        this.$emit('input', this.searchQuery);
      },
      handleSearch() {
        // 当点击搜索按钮时触发
        this.$emit('search', this.searchQuery);
      }
    }
  };
  </script>
  
  <style scoped>
  .search-box {
    width: 80%;
    display: flex;
    align-items: center;
  }
  
  input {
    width: 90%;
    padding: 8px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 4px;
    margin-right: 8px;
  }
  
  button {
    height: 60%;
    padding: 0px 16px;
    font-size: 16px;
    background-color: #9fccf3;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    display: grid;
    place-items: center; /* 水平和垂直居中 */
    height: 36px; /* 设置按钮的高度 */  
    white-space: nowrap;
}
  
  button:hover {
    background-color: #587dc6;
  }
  </style>