<template>
  <div class="chat-area">
    <div class="part1">
      <h2>{{ selectedChat.name }}</h2>
      <p>{{ selectedChat.remark }}</p>
      <div class="roomInfo">
        <div class="left">
        </div>
        <div class="right">
          <button v-if="isRoomOwner" class="addMemberButton" @click="showPopup('addMember')"></button>
          <button class="setButton" @click="showPopup('settings')"></button>
          <button class="memberButton" @click="showPopup('members')"></button>
        </div>
      </div>
    </div>
    <div ref="messagesContainer" class="messages">
      <AMessage v-for="message in messages" :key="message.id" :message="message" :is-current-user="message.UID === uid"/>
    </div>
    <div class="input-area">
      <input type="text" v-model="newMessage" @keyup.enter="sendMessage" placeholder="输入消息..." />
      <button @click="sendMessage">发送</button>
    </div>

    <!-- 弹出窗口 -->
    <PopupWindow v-if="isPopupVisible" @close="closePopup" class="popup-window">
      <div v-if="popupType === 'addMember'">
        <p class="title">添加成员</p>
        <div class="mylist">
          <ul>
            <li v-for="chatmember in chatMembers" :key="chatmember.id">
              <div class="addMemberList">
                <div class="addmemberInfo">
                  <img class="avatar" :src="`http://localhost:8082/avatars/${chatmember.userID}.jpg`">
                  成员名{{ chatmember.userName }}
                </div>
                <button class="list-button" @click="addMember(chatmember.userID)">添加</button>
              </div>
            </li>
          </ul>
        </div>
      </div>
      <div v-else-if="popupType === 'settings'">
        <p class="title">设置</p>
        <div class="setting">
          <button v-if="!isRoomOwner" @click="getOutRoom_">退出房间</button>
          <button v-if="isRoomOwner" @click="dissolveRoom">解散房间</button>
        </div>
      </div>
      <div v-else-if="popupType === 'members'">
        <p class="title">成员列表</p>
        <ul>
          <li v-for="roomMember in roomMembers" :key="roomMember.id">
            <div class="roomMemberList">
              <div class="roomMemberInfo">
                <img class="avatar" :src="`http://localhost:8082/avatars/${roomMember.userID}.jpg`">
                {{ roomMember.userName }}
              </div>
              <button class="list-button" v-if="isRoomOwner" @click="removeMember(roomMember.userID)">删除</button>
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
    chatMembers:{
      type: Array,
    },
    roomMembers:{
      type: Array,
    },
    isRoomOwner:{
      type:Boolean,
      required:true,
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
  mounted() {
    this.scrollToBottom();
  },
  updated() {
    this.scrollToBottom();
  },
  watch: {
    messages() {
      this.$nextTick(() => {
        this.scrollToBottom();
      });
    },
  },
  methods: {
    scrollToBottom() {
      const container = this.$refs.messagesContainer;
      if (container) {
        container.scrollTop = container.scrollHeight;
      }
    },
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
    range(n) {
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
    getOutRoom_(){
      this.$emit('getout-room');
    }
  },
  created() {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      const user = JSON.parse(storedUser);
      this.uid = user.UID;
    }
  },
};
</script>

<style scoped>
/* 整体布局 */
.chat-area {
  flex: 1;
  padding: 20px;
  background-color: #f9f9f97e;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.part1 {
  height: 15%;
  margin-bottom: 20px;
}

.roomInfo {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.right {
  display: flex;
  gap: 10px;
}

.right button {
  width: 40px;
  height: 40px;
  border: none;
  border-radius: 20%;
  background-color: #ffffff00;
  color: white;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.right button:hover {
  background-color: hsla(0, 0%, 100%, 0.754);
}
.right .setButton {
  background-image: url('@/assets/齿轮.png');
  background-size: cover;
}
.right .addMemberButton {
  background-image: url('@/assets/加号.png');
  background-size: cover;
}
.right .memberButton {
  background-image: url('@/assets/成员.png');
  background-size: cover;
}

/* 消息区域 */
.messages {
  height: 72%;
  overflow-y: auto;
  background-color: rgba(255, 255, 255, 0.575);
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

/* 自定义滚动条样式 */
.messages::-webkit-scrollbar {
  width: 6px;
}

.messages::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.messages::-webkit-scrollbar-thumb {
  background: #ff9191;
  border-radius: 4px;
}

.messages::-webkit-scrollbar-thumb:hover {
  background: #ff6b6b;
}

/* 输入区域 */
.input-area {
  background-color: #ffffff00;
  display: flex;
  gap: 10px;
}

input {
  flex: 1;
  padding: 10px;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  font-size: 16px;
  background-color: rgba(255, 255, 255, 0.526);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

button {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  background-color: #ff9191;
  color: white;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #ff6b6b;
}

/* 弹出窗口 */
.popup-window {
  background-color: rgba(255, 255, 255, 0.412);
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.title {
  font-size: 18px;
  font-weight: bold;
  color: #ff6b6b;
  margin-bottom: 20px;
}

.mylist ul, .members ul {
  list-style: none;
  padding: 0;
}

.addMemberList, .roomMemberList {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #e0e0e0;
}

.addmemberInfo, .roomMemberInfo {
  display: flex;
  align-items: center;
  gap: 10px;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.list-button {
  padding: 6px 12px;
  border: none;
  border-radius: 6px;
  background-color: #ff9191;
  color: white;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.list-button:hover {
  background-color: #ff6b6b;
}

/* 设置区域 */
.setting {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.setting button {
  width: 100%;
  padding: 10px;
  border: none;
  border-radius: 8px;
  background-color: #ff9191;
  color: white;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.setting button:hover {
  background-color: #ff6b6b;
}

.mylist, ul {
  list-style: none; /* 移除列表项前的标记 */
  padding-left: 0; /* 可选：移除默认的左内边距 */
}

</style>