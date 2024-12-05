<template>
    <GuideBar />
    <h1>聊天室界面</h1>
    <div class="chat-container">
      <ChatList @chat-selected="onChatSelected" />
      <ChatArea @room-status-changed="roomStatusChange" :selectedChat="selectedChat" :messages="messages" :isSecret="isSecret"  v-if="selectedChat" @send-message="handleSendMessage"/>
    </div>
  </template>
  
  <script>
  import ChatList from '../components/ChatList.vue';
  import ChatArea from '../components/ChatArea.vue';
  import GuideBar from '../components/GuideBar.vue';
  import axios from 'axios';

  export default {
    name: 'ChatPage',
    components: {
      GuideBar,
      ChatList,
      ChatArea,
    },
    data() {
      return {
        uid:null,
        userImg:'',
        username:'',
        selectedChat: null,
        messages: [],
        isSecret:false,
        isDissolved:false
      };
    },
    created() {
        // 从本地存储中获取用户信息
        const storedUser = localStorage.getItem('user');
        if (storedUser) {
            const user = JSON.parse(storedUser);
            this.uid = user.UID;
            this.username = user.UName;
        }
        this.userImg=localStorage.getItem('userImage');
    },
    methods: {
      roomStatusChange(data){
        this.isSecret=data.isSecret;
        this.isDissolved=data.isdissolved
        //将房间状态更新信息传给后端
      },
      onChatSelected(data) {
        this.selectedChat = data.chat;
        this.messages = data.messages;
      },
      //处理发送信息
      handleSendMessage(newMessage) {
        const message = {
            id: this.messages.length + 1,
            text: newMessage,
            senderId: this.uid,
            senderName:this.username,
            senderAvatarSrc:this.userImg,
            isImage:false, 
            imageSrc:''
        };
        const newMessageData = JSON.stringify(message);

        // 发送新消息请求到后端
        axios.get('http://localhost:3000/api/sendchats/', { 
            params :{
                roomid: this.selectedChat.id,
                newmessage: newMessageData
            }
        }).then(response => {
            console.log('Message sent successfully:', response.data);  
            this.messages.push(message);// 更新本地消息列表
        }).catch(error => {
           alert("发送消息失败！"+error);
        });
      }
    },
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