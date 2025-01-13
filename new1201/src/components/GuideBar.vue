<!-- src/components/GuideBar.vue -->
<template>
    <div>
      <div class="mybackground">
        <div class="searchbody">
          <div class="left">
            <img :src="LogoImage" class="logoImg"/>
          </div>
          <div class="right">
            <button @click="toHomePage">首页</button>
            <button @click="toFindPage">找搭子</button>
            <button @click="toChatPage">聊天室</button>
            <button @click="toAboutUsPage">关于我们</button>
            <div class="newARoom">
              <button @click="toNewRoomPage" title="新建房间">+</button>
            </div>
            <div class="userInfo" @mouseenter="showOptions = true" @mouseleave="showOptions = false">
              <img :src="UserAvater" class="userImg" />
              <OptionsOfGuideBar v-if="showOptions" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import router from '@/router';
  import LogoImage from "@/assets/logo.png";
  import OptionsOfGuideBar from './OptionsOfGuideBar.vue';
  
  export default {
    name: 'GuideBar',
    components: {
        OptionsOfGuideBar
    },
    data() {
      return {
        LogoImage: LogoImage,
        showOptions: false
      }
    },
    setup() {
      let UserID = 0;
      try {
          const storedUser = localStorage.getItem('user');
          if (storedUser) {
              const user = JSON.parse(storedUser);
              UserID = user.UID;
          }
      } catch {
          alert("获取本地个人信息失败");
      }
      const UserAvater=`http://localhost:8082/avatars/${UserID}.jpg`;
      return {
        userID:UserID,
        UserAvater
      };
    },
    methods: {
      toHomePage() {
        router.push("HomePage");
      },
      toFindPage() {
        router.push("FindPage");
      },
      toChatPage() {
        router.push("ChatPage");
      },
      toPersonalPage() {
        router.push("PersonalPage");
      },
      toAboutUsPage() {
        router.push("AboutUsPage");
      },
      toNewRoomPage() {
        router.push("NewRoomPage");
      },
    }
  }
  </script>
  
  <style scoped>
  .mybackground {
    background: linear-gradient(to bottom, rgb(9, 100, 175), rgba(255, 255, 255, 0));
    padding: 20px;
  }
  
  .searchbody {
    margin-top: 10px;
    display: flex;
    justify-content: space-between; /* 子元素之间的间距 */
    align-items: center; /* 子元素垂直居中 */
    background-color: #f2f6fc;
    border-radius: 50px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }
  
  .logoImg {
    margin-left: 30px;
    height: 50px;
  }
  
  .right {
    display: flex;
    width: 75%;
  }
  
  .userInfo {
    position: relative;
    display: flex;
    align-items: center; /* 垂直居中 */
    margin-right: 25px;
  }
  
  .userImg {
    margin-left: 30px;
    height: 50px;
    border-radius: 50%; /* 设置为圆形 */
  }
  
  button {
    width: 100%;
    height: 80px;
    color: #000000;
    background-color: transparent;
    border: none;
    font-size: 175%;
    font-family: '宋体', SimSun, serif;
  }
  
  button:hover {
    color: #9696969d;
  }
  
  .sayhi {
    font-size: 100%;
  }

  .newARoom {
    display: flex;
    align-items: center;
    justify-content: center;

  }

  .newARoom button {
    color: white; /* 调整按钮颜色 */
    font-size: 24px; /* 调整字体大小 */
    background-color: #000000;
    font-weight:bolder;
    height: 30px; /* 调整高度 */
    width: 30px; /* 调整宽度 */
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-left: 10px; /* 调整间距 */
  }
  .newARoom button:hover {
    background-color: #0000009c;
  }
  </style>