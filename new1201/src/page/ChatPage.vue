<template>
  <GuideBar />
  <div class="chat-container">
    <ChatList 
      @chat-selected="onChatSelected" 
      @search-chats="getChatsList_"
      :chats="chats"
    />
    <ChatArea v-if="selectedChat"
      @set-dissolve="changeRoomStatus()" 
      @send-message="handleSendMessage"
      @get-chat-member="getchatMember_"
      @add-member="addMember_"
      @get-room-member="getRoomMember_"
      @remove-member="removeMember"
      :chatMembers="chatMember" 
      :roomMembers="roomMember" 
      :selectedChat="selectedChat" 
      :messages="messages" 
      :isSecret="isSecret"  
      :chats="chats"
      :isRoomOwner="selectedChat.ownerId==uid"
    />
  </div>
</template>

<script>
import ChatList from '../components/ChatList.vue';
import ChatArea from '../components/ChatArea.vue';
import GuideBar from '../components/GuideBar.vue';
import {getChatsList,getChatMessages,removeMember,getRoomMember,addMember,getChatMember,dissolveRoom,sendMessage} from '../api';

export default {
name: 'ChatPage',
components: {
  GuideBar,
  ChatList,
  ChatArea,
},
data() {
  return {
    uid: null,
    userImg: '',
    username: '',
    chats: [],          //聊天列表
    //当前聊天室信息
    selectedChat: null,  //保存了聊天信息{id,name,}
    messages: [],
    isSecret: false,
    isDissolved: false,
    roomId: 0,
    chatMember: [], //加入会话但是没加入项目的成员
    roomMember: [], //加入项目的成员
    pollingInterval: null, // 长轮询的定时器
  };
},
created() {
  // 从本地存储中获取用户信息
  try {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      const user = JSON.parse(storedUser);
      this.uid = user.UID;
      this.username = user.UName;
    }
    this.userImg = localStorage.getItem('userImage');
    this.getChatsList_(""); //获取聊天列表
  } catch {
    alert("获取本地个人信息失败");
  }
},
methods: {
  //将项目成员移除
  async removeMember(id){
    try {
      let ok=false;
      ok=await removeMember(this.selectedChat.id,id);
      if(ok)alert("移除成功！");
      else alert("移除失败！");
      this.getRoomMember_();
    } catch (error) {
      alert("移除失败！");
    } 
  },
  //从后端请求项目成员
  async getRoomMember_(){
    try{
      const response = await getRoomMember(this.selectedChat.id);
      this.roomMember=response;
    }catch{
      alert("获取房间成员列表失败！");
    }
  },
  //将会话成员添加到项目中
  async addMember_(id){
    try {
      let ok=false;
      ok=await addMember(this.selectedChat.id,id);
      if(ok){
        alert("添加成功！");
        this.getchatMember_();
      }
      else alert("添加失败！");
    } catch (error) {
      alert("添加失败！");
    } 
  },
  //从后端请求在会话中但不在项目中的成员
  async getchatMember_(){
    try{
      const response = await getChatMember(this.selectedChat.id);
      this.chatMember=response;
    }catch{
      alert("获取会话成员列表失败！");
    }
  },
  //获取聊天列表
  async getChatsList_(searchWord) {
    try {
      const data=await getChatsList(this.uid,searchWord);
      this.chats=data;                     
    } catch {
      alert("聊天列表获取失败");
    }
  },
  //更改房间状态:改成解散房间
  changeRoomStatus() {
    if (confirm("你确定解散房间吗？")) {
        if (this.sendRoomStatus("解散")) {
          this.isdissolved = true;
          alert("解散房间成功！");
          //解散后的逻辑(需补充)
          this.selectedChat = null;
          this.getChatsList_("");
        }
      }
  },
  //向后端传输状态更新信息
  async sendRoomStatus() {
    try {
      let ok=true;
      ok=await dissolveRoom(this.selectedChat.id);
      return ok;
    } catch (error) {
      return false;
    }
  },
  //列表选中房间，加载房间历史聊天记录
  async onChatSelected(chatInfo) {
    this.stopPolling() ;
    this.selectedChat = chatInfo;
    try {
      // 获取历史聊天信息
      const data=await getChatMessages(chatInfo.id,0);
      this.messages=data;
      // 开启长轮询
      this.startPolling();
    } catch (error) {
      alert('获取聊天记录失败！');
    }
  },
  //处理发送信息
  async handleSendMessage(newMessage) {
    const message = {
      id: this.messages.length + 1,
      text: newMessage,
      senderId: this.uid,
      senderName: this.username,
      senderAvatarSrc: this.userImg,
      isImage: false,
      imageSrc: ''
  };

    // 发送新消息请求到后端
    try{
      let ok=false;
      ok=await sendMessage(this.selectedChat.id,message);
      if(!ok){
        alert("发送消息失败！");
      }
    }catch{
      alert("发送消息失败！");
    }
  },
  // 开启长轮询
    startPolling() {
    if (this.pollingInterval) {
      clearInterval(this.pollingInterval);
    }

    this.pollingInterval = setInterval(async () => {
      try {
        //alert("messages: "+JSON.stringify(this.messages));
        const lastMessageId = this.messages.length;
        const data=await getChatMessages(this.selectedChat.id,lastMessageId);
        if (data.length > 0) {
          this.messages.push(...data);
        }
      } catch (error) {
        console.error('Error polling for new messages:', error);
      }
    }, 1000); // 每1秒轮询一次 
  },
  // 停止长轮询
  stopPolling() {
    if (this.pollingInterval) {
      clearInterval(this.pollingInterval);
      this.pollingInterval = null;
    }
  }
},
beforeUnmount() {
  this.stopPolling(); // 组件销毁前停止长轮询
}
};
</script>

<style scoped>
.chat-container {
  max-width: 1200px;
  display: flex;
  width: 80%;
  margin: 0 auto;
  border: 1px solid #ccc;
  height: 800px;
  margin: 140px auto 100px; 
}
</style>