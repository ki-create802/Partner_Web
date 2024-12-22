<!-- 聊天室界面 组件-->
<template>
  <div class="chat-area">
    <div class="part1">
      <h2>{{ selectedChat.name }}</h2>
      <div class="roomInfo">
        <div class="left">
        </div>
        <div class="right">
          <button class="addMemberButton" @click="showPopup('addMember')" ></button>
          <button class="setButton" @click="showPopup('settings')"></button>
          <button class="memberButton" @click="showPopup('members')"></button>
        </div>
      </div>
    </div>
    <div class="messages">
      <AMessage v-for="message in messages" :key="message.id" :message="message" :is-current-user="message.senderId === uid"/>
    </div>
    <div class="input-area">
      <input type="text" v-model="newMessage" @keyup.enter="sendMessage" placeholder="输入消息..." />
      <button @click="sendMessage">发送</button>
    </div>

    <!-- 弹出窗口 -->
    <PopupWindow v-if="isPopupVisible" @close="closePopup">
      <div v-if="popupType === 'addMember'">
        <h3>添加成员</h3>
        <!-- 添加成员的表单或内容 -->
        <div class="mylist">
        <ul>
          <li v-for="chatmember in chatMembers" :key='chatmember.id' >
            <div class="addMemberList">
              <div class="addmemberInfo">
                <img class="avatar" :src="'http://localhost:3000/' + chatmember.avatar">
                成员名{{ chatmember.name }}
              </div>
              <button @click="addMember(chatmember.id)">添加</button>
            </div>
          </li>
        </ul>
      </div>

      </div>
      <div v-else-if="popupType === 'settings'">
        <h3>设置</h3>
        <!-- 设置的表单或内容 -->
         <div class="setting">
          <button v-if="isSecret" @click="setPublic">将房间设置为公开</button>
          <button v-else @click="setSecret">将房间设置为私密</button>
          <button @click="dissolveRoom">解散房间</button>
         </div>
      </div>


      <div v-else-if="popupType === 'members'">
        <h3>成员列表</h3>
        <!-- 成员列表的表单或内容 -->
        <ul>
          <li v-for="roomMember in roomMembers" :key='roomMember.id' >
            <div class="roomMemberList">
              <div class="roomMemberInfo">
                <img class="avatar" :src="'http://localhost:3000/' + roomMember.avatar">
                成员名{{ roomMember.name }}
              </div>
              <button @click="removeMember(roomMember.id)">删除</button>
            </div>
          </li>
        </ul>
      </div>
    </PopupWindow>
  </div>
</template>

<script>
import AMessage from './AMessage.vue';
import PopupWindow from './PopupWindow.vue';

export default {
  name: 'ChatArea',
  components: {
    AMessage,
    PopupWindow,
  },
  props: {
    selectedChat: {
      type: Object,
      required: true,
    },
    messages: {
      type: Array,
    },
    isSecret:{
      type:Boolean,
      required:true,
    },
    //会话成员列表每一项为{userid,avatar}
    chatMembers:{
      type: Array,
    },
    roomMembers:{
      type: Array,
    }
  },
  data() {
    return {
      newMessage: '',
      uid: null,
      isPopupVisible: false,
      popupType: '',
    };
  },
  methods: {
    removeMember(id){
      this.$emit('remove-member',id)
    },
    getRoomMember(){
      this.$emit('get-room-member');
    },
    addMember(id){
      this.$emit('add-member',id);
    },
    getchatMember(){
      this.$emit('get-chat-member');
    },
    setSecret(){
      this.$emit('set-secrect');
    },
    setPublic(){
      this.$emit('set-public');
    },
    dissolveRoom(){
      this.$emit('set-dissolve');
    },
    range(n) {//一个无用函数，用于调试
      return Array.from({ length: n }, (_, i) => i + 1);
    },
    sendMessage() {
      if (this.newMessage.trim() === '') return;
      this.$emit('send-message', this.newMessage);
      this.newMessage = '';
    },
    showPopup(type) {
      this.popupType = type;
      this.isPopupVisible = true;
      if(type=="addMember") this.getchatMember();
      if(type=="members")this.getRoomMember();
    },
    closePopup() {
      this.isPopupVisible = false;
    },
  },
  created() {
    // 从本地存储中获取用户信息
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      const user = JSON.parse(storedUser);
      this.uid = user.UID;
    }
  },
};
</script>

<style scoped>
.chat-area {
  flex: 1;
  padding: 10px;
}
.part1 {
  height: 15%;
}
.roomInfo {
  width: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
.left {
  align-items: start;
}
.right {
  width: 50%;
  display: flex; /* 确保使用 Flexbox 布局 */
  justify-content: flex-end; /* 将子元素推到容器的末尾 */
  align-items: center; /* 垂直居中对齐 */
  white-space: nowrap; /* 防止内容换行 */
}
.right button {
  width: 40px;
  height: 40px;
  box-shadow: #ccc;
}
.right .setButton {
  background-image: url('@/assets/齿轮.jpg');
  background-size: cover;
}
.right .addMemberButton {
  background-image: url('@/assets/加号.jpg');
  background-size: cover;
}
.right .memberButton {
  background-image: url('@/assets/成员.jpg');
  background-size: cover;
}
.memberAvatar {
  width: 50%;
  display: flex; /* 使用 Flexbox 布局 */
  overflow-x: auto; /* 启用横向滚动条 */
  white-space: nowrap; /* 防止内容换行 */
}
/* 自定义滚动条样式 */
.memberAvatar::-webkit-scrollbar {
  height: 6px; /* 滚动条宽度 */
}

.memberAvatar::-webkit-scrollbar-track {
  background: #f1f1f1; /* 滚动条轨道背景色 */
}

.memberAvatar::-webkit-scrollbar-thumb {
  background: #e3e3e3; /* 滚动条滑块颜色 */
  border-radius: 4px; /* 滚动条滑块圆角 */
}
.avatar {
  width: 40px; /* 设置图像的宽度 */
  min-width: 40px;
  height: 40px; /* 设置图像的高度 */
  border-radius: 50%; /* 使图像显示为圆形 */
  vertical-align: middle; /* 使图像垂直居中 */
}
.messages {
  height: 72%;
  overflow-y: auto;
  border: 1px solid #ccc;
  padding: 10px;
  margin-bottom: 10px;
}

.input-area {
  display: flex;
}

input {
  height: 13%;
  background-color: rgb(224, 244, 255);
  border: none;
  border-radius: 4px;
  height: 40px;
  flex: 1;
  padding: 5px;
  font-size: 18px;
}

button {
  width: 60px;
  background-color: rgb(103, 150, 211);
  border: none;
  border-radius: 4px;
  padding: 5px 10px;
  font-size: 18px;
  color: white;
}

button:hover {
  background-color: rgb(47, 94, 157);
}
.setting{
  display: flex;
  flex-direction: column;
  gap:10px
}
.setting button{
  width: 100%;
  background-color: #b9b9b9;
}
.setting button:hover{
  background-color: #757575;
}
</style>