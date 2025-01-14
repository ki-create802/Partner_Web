<!-- src/components/UserOptions.vue -->
<template>
    <div class="user-options">
      <div class="option" @click="handleLogout">退出登录</div>
      <div class="option" @click="handleChangePassword">修改密码</div>
      <div class="option" @click="toPersonalPage">个人主页</div>
    </div>
  </template>
  
  <script>
import router from '@/router';
  export default {
    name: 'OptionsOfGuideBar',
    methods: {
      handleLogout() {
        // 处理退出登录逻辑
        localStorage.clear();
        router.push("/");
        console.log('退出登录');
      },
      handleChangePassword() {
        // 处理修改密码逻辑
        router.push("ModifyPWPage");
        console.log('修改密码');
      },
      toPersonalPage(){
        router.push("PersonalPage");
      }
    }
  }
  </script>
  
  <style scoped>
  .user-options {
    position: absolute;
    top: 100%;
    right: 0;
    background-color: white;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    padding: 8px;
    width: 100px;
    z-index: 998;
  }
  
  .option {
    padding: 8px;
    cursor: pointer;
  }
  
  .option:hover {
    background-color: #f2f2f2;
  }

  .form-group {
    margin-bottom: 15px;
  }
  </style>