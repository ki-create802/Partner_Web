<template>
  <div class="chat-list">
    <div class="searchContainer">
      <input type="text" v-model="searchWord" @keyup.enter="searchChats" placeholder="查找聊天室..." />
      <button @click="searchChats">查找</button>
    </div>
    <div class="mylist">
      <ul>
        <li v-for="chat in chats" :key="chat.id" @click="selectChat(chat)">
          <img class="roomOwnerImg" :src="'http://localhost:3000/'+chat.ownerImage"/>
          {{ chat.name }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>

export default {
  name: 'ChatList',
  props: {
    chats:{
      type:Array,
      Required:true
    }
  },
  data() {
    return {
      searchWord:'',
    };
  },
  methods: {
    async searchChats(){
      if(this.searchWord.trim()=="")return;
      else {
        this.$emit('search-chats', this.searchWord);
      }
    },
    async selectChat(chat) {
      //向前端发送点击房间的房间信息
      this.$emit('chat-selected', chat);
    },
  },
};
</script>

<style scoped>
.searchContainer{
  height: 40px;
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
}

.searchContainer input{
  height: 100%;
  width: 80%;
  border: none;
  font-size: 16px;
}
.searchContainer button{
  background-color: #b8f1ff;
  border: none;
  height: 100%;
  width: 20%;
  color:rgb(93, 145, 194);
}
.searchContainer button:hover{
  background-color: #a7cbef;;
}
.chat-list {
  background-color: #e8f6fd;
  width: 300px;
  border-right: 1px solid #ccc;
  padding: 10px;
}
.mylist{
  height: 95%;
  overflow-y: auto;
}
/* 自定义滚动条样式 */
.mylist::-webkit-scrollbar {
  width: 2px; /* 滚动条宽度 */
}

.mylist::-webkit-scrollbar-track {
  background: #f1f1f1; /* 滚动条轨道背景色 */
}

.mylist::-webkit-scrollbar-thumb {
  background: #9bcbe4; /* 滚动条滑块颜色 */
  border-radius: 4px; /* 滚动条滑块圆角 */
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  cursor: pointer;
  padding: 5px;
  border-bottom: 1px solid #9d9d9d82;
  display: flex; /* 使用 Flexbox 布局 */
  align-items: center; /* 垂直居中对齐 */
}

li:hover {
  background-color: #b1b1b129;
}

.roomOwnerImg {
  width: 40px; /* 设置图像的宽度 */
  height: 40px; /* 设置图像的高度 */
  border-radius: 50%; /* 使图像显示为圆形 */
  vertical-align: middle; /* 使图像垂直居中 */
  margin-right: 10px; /* 给图像和文本之间添加一些间距 */
}
</style>