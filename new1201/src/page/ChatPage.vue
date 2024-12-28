<template>
  <GuideBar />
  <h1>聊天室界面</h1>
  <div class="chat-container">
    <ChatList 
      @chat-selected="onChatSelected" 
      @search-chats="getChatsList"
      :chats="chats"
    />
    <ChatArea v-if="selectedChat"
      @set-dissolve="changeRoomStatus('解散')" 
      @set-secrect="changeRoomStatus('私密')" 
      @set-public="changeRoomStatus('公开')" 
      @send-message="handleSendMessage"
      @get-chat-member="getchatMember"
      @add-member="addMember"
      @get-room-member="getRoomMember"
      @remove-member="removeMember"
      :chatMembers="chatMember" 
      :roomMembers="roomMember" 
      :selectedChat="selectedChat" 
      :messages="messages" 
      :isSecret="isSecret"  
      :chats="chats"
    />
  </div>
</template>

<script>
import ChatList from '../components/ChatList.vue';
import ChatArea from '../components/ChatArea.vue';
import GuideBar from '../components/GuideBar.vue';
import axios from 'axios';
import {getChatsList,getChatMessages} from '../api';

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
    //聊天列表
    chats: [],
    //当前聊天室信息
    selectedChat: null,  //保存了聊天信息{id,name,ownerImage}
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
      await axios.post('http://localhost:3000/api/removeMember', {
        roomid: this.selectedChat.id,
        RoomMemberId: id,
      });
    } catch (error) {
      alert("移除失败！");
    } 
    alert("移除成功！");
    this.getRoomMember();
  },
  //从后端请求项目成员
  async getRoomMember(){
    try{
      const response = await axios.get(`http://localhost:3000/api/getRoomMember`, {
        params: {
          roomid: this.selectedChat.id,
        }
      });
      this.roomMember=response.data;
    }catch{
      alert("获取房间成员列表失败！");
    }
  },
  //将会话成员添加到项目中
  async addMember(id){
    try {
      await axios.post('http://localhost:3000/api/addMember', {
        roomid: this.selectedChat.id,
        chatMemberId: id,
      });
    } catch (error) {
      alert("添加失败！");
    } 
    alert("添加成功！");
    this.getchatMember();
  },
  //从后端请求在会话中但不在项目中的成员
  async getchatMember(){
    try{
      const response = await axios.get(`http://localhost:3000/api/getChatMember`, {
        params: {
          roomid: this.selectedChat.id,
        }
      });
      this.chatMember=response.data;
    }catch{
      alert("获取会话成员列表失败！");
    }
  },
  //获取聊天列表
  async getChatsList_(searchWord) {
    try {
      const data=await getChatsList(this.uid,searchWord);
      alert(JSON.stringify(data));
      this.chats=data;
    } catch {
      alert("聊天列表获取失败");
    }
  },
  //更改房间状态
  changeRoomStatus(changeInfo) {
    if (changeInfo == "解散") {
      if (confirm("你确定解散房间吗？")) {
        if (this.sendRoomStatus("解散")) {
          this.isdissolved = true;
          alert("解散房间成功！");
          //解散后的逻辑(需补充)
          this.selectedChat = null;
          this.getChatsList("");
        }
      }
    } else if (changeInfo == "私密") {
      if (confirm("你确定将房间设置为私密吗？ps：私密后只有搭子才能看见以及进入房间。")) {
        if (this.sendRoomStatus("私密")) {
          this.isSecret = true;
          alert("设置成功！");
        }
      }
    } else if (changeInfo == "公开") {
      if (confirm("你确定将房间设置为公开吗？ps：公开后所有人都能发现你的房间。")) {
        if (this.sendRoomStatus("公开")) {
          this.isSecret = false;
          alert("设置成功！");
        }
      }
    }
  },
  //向后端传输状态更新信息
  async sendRoomStatus(roomInfo) {
    try {
      const response = await axios.get(`http://localhost:3000/api/setRoomStatus`, {
        params: {
          roomid: this.selectedChat.id,
          setStatus: roomInfo
        }
      });
      if (!response.data.success) {
        return false;
      }
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
  handleSendMessage(newMessage) {
    const message = {
      id: this.messages.length + 1,
      text: newMessage,
      senderId: this.uid,
      senderName: this.username,
      senderAvatarSrc: this.userImg,
      isImage: false,
      imageSrc: ''
    };
    const newMessageData = JSON.stringify(message);

    // 发送新消息请求到后端
    axios.get('http://localhost:3000/api/sendchats/', {
      params: {
        roomid: this.selectedChat.id,
        newmessage: newMessageData
      }
    }).then(response => {
      console.log('Message sent successfully:', response.data);
    }).catch(error => {
      alert("发送消息失败！" + error);
    });
  },
  // 开启长轮询
    startPolling() {
    if (this.pollingInterval) {
      clearInterval(this.pollingInterval);
    }

    this.pollingInterval = setInterval(async () => {
      try {
        const lastMessageId = this.messages.length;
        const data=await getChatMessages(this.selectedChat.id,lastMessageId);
        if (data.length > 0) {
          this.messages.push(...data);
          //对列表消息进行去重
          this.messages = this.messages.filter((item, index, self) => self.findIndex(t => t.id === item.id) === index);
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
}
</style>