<template>
  <div class="message-container">
    <div v-if="!isCurrentUser" class="message left">
      <div class="senderInfo">
        <img class="userImg" :src="`http://localhost:8082/avatars/${message.UID}.jpg`"/>
        <strong>{{ message.UName }}</strong>
      </div>
      <div class="message-content">
         {{ message.RContent }}
      </div>
    </div>
    <div v-else class="message right">
      <div class="message-content">
         {{ message.RContent }}
      </div>
      <div class="senderInfo">
        <img class="userImg" :src="`http://localhost:8082/avatars/${message.UID}.jpg`"/>
        <strong>{{ message.UName }}</strong>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AMessage',
  props: {
    message: {
      type: Object,
      required: true,
    },
    isCurrentUser: {
      type: Boolean,
      default: false,
    },
  },
};
</script>

<style scoped>
.message-container {
  margin-bottom: 10px;
}

.message {
  display: flex;
  align-items: center;
}

.left {
  justify-content: flex-start; /* 消息靠左显示 */
}

.right {
  justify-content: flex-end; /* 消息靠右显示 */
}

.message-content {
  background-color: rgb(255, 230, 235);
  border-radius: 10%;
  text-align: left; /* 确保消息内容靠左显示 */
  margin-right: 20px; 
  margin-left: 20px;
  padding: 10px;
  font-size: 24px;
}

.right .message-content {
  text-align: right; /* 当前用户的消息内容靠右显示 */
}

.senderInfo {
  display: flex;
  flex-direction: column; /* 使子元素竖直排列 */
  align-items: center; /* 使子元素在垂直方向上居中对齐 */
}

.userImg {
  width: 40px; /* 设置图像的宽度 */
  height: 40px; /* 设置图像的高度 */
  border-radius: 50%; /* 使图像显示为圆形 */
  vertical-align: middle; /* 使图像垂直居中 */
}

</style>