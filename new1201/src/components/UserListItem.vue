<template>
  <div class="userListContainer">
    <div class="userItem">
      <button class="userImageButton" @click="goToPersonalPage">
        <img class="userImage" :src="`http://localhost:8082/avatars/${user.UID}.jpg`" alt="User Image" />
      </button>
      <div class="userInfo">
        <a class="userName">{{ user.UName }}</a>
        <a class="userRemark">个性签名：{{ user.URemark }}</a>
      </div>
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
    methods: { // 将方法放在 methods 中
      goToPersonalPage() {
        this.$router.push({
          name: 'PersonalPage2',
          query: {
            UID: this.user.UID, // 传递用户 ID
          },
        });
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

/* 头像按钮样式 */
.userImageButton {
  padding: 0;
  border: none;
  background: none;
  cursor: pointer;
  margin-right: 20px;
}

/* 添加悬停效果 */
.userImageButton:hover .userImage {
  transform: scale(1.1); /* 悬停时放大 */
}

.userListContainer {
  width: 70%; /* 设置宽度为页面宽度的 70% */
  margin: 0 auto; /* 居中显示 */
}

  </style>