<!-- src/components/GuideBar.vue -->
<template>
    <body>
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
            <div class="userInfo" @mouseenter="showOptions = true" @mouseleave="showOptions = false">
              <img :src="LogoImage" class="userImg" />
              <OptionsOfGuideBar v-if="showOptions" />
            </div>
          </div>
        </div>
      </div>
    </body>
  </template>
  
  <script>
  import router from '@/router';
  import { inject } from 'vue';
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
      const userID = inject("userID");
      return {
        userID,
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
    }
  }
  </script>
  
  <style scoped>
  .mybackground {
    background: linear-gradient(to bottom, rgba(91, 181, 255, 0.2), rgba(255, 255, 255, 0.2));
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
  </style>