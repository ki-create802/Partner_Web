<template>
    <div class="userItem">
      <img class="userImage" :src="`http://localhost:8082/avatars/${user.UID}.jpg`" alt="User Image" />
      <div class="userInfo">
        <a class="userName">{{ user.UName }}</a>
        <a class="userRemark">个性签名：{{ user.URemark }}</a>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'UserListItem',
    props: {
      user: {
        type: Object,
        required: true,
      },
    },
  }
  </script>
  
  <style scoped>
  .userItem {
    display: flex;
    align-items: center;
    padding: 15px;
    border-bottom: 1px solid #e0e0e0;
    background-color: #ffffff;
    transition: background-color 0.3s ease;
  }
  
  .userItem:hover {
    background-color: #f5f5f5;
  }
  
  .userImage {
    width: 60px;
    height: 60px;
    border-radius: 50%;
    margin-right: 20px;
    border: 2px solid #e0e0e0;
  }
  
  .userInfo {
    display: flex;
    flex-direction: column;
  }
  
  .userName {
    font-size: 18px;
    font-weight: bold;
    color: #333333;
    margin-bottom: 5px;
  }
  
  .userRemark {
    font-size: 14px;
    color: #666666;
  }
  </style>