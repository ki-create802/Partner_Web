<template>
  <div class="information">
    <div class="avatar-container">
      <div class="avatar">
        <img :src="`http://localhost:8082/avatars/${userInfo.UID}.jpg`" alt="头像" id="avatar-img">
      </div>
      <button class="follow-btn" @click="triggerFollow" :class="{ 'followed': isFollowed }"> {{ isFollowed ? '取消关注' : '+关注' }} </button>
    </div>
    <div class="user-info">
      <div class="username">
        <span id="username-text">{{ userInfo.UName }}</span>
      </div>
      <div class="followers">
        <span id="followers-text">粉丝数: {{ fan }}</span>
      </div>
    </div>
    <div class="information-right">
      {{ userInfo.URemark }}
    </div>
  </div>
</template>

<script>
//   import axios from 'axios';
import { getFansNum } from '@/api.js';
import { followadd } from '@/api.js';
import { unfollow } from '@/api.js';


export default {
  name: 'InformationBox',
  props: {
    userInfo: {
      type: Object,
      required: true,
      default: () => ({ // 提供默认值
        UID: '',
        UName: '',
        Password: '',
        Email: '',
        URemark: '',
        UImage: '',
      })
    },
    fan: {
      type: Number, // 根据实际情况调整类型
      required: true,
      //default: 0 // 默认值
    },
    isFollowed: {
      type: Boolean,
      required: true,

    }
  },

  data() {
    return{
      avatarUrl:require("@/assets/avatar.png"), // 默认头像
      followersCount: 2000,
    };
  },
  methods: {
    async triggerFollow() {//alert
      try {
        //获取个人id
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

        //发送关注或取消关注的请求
        if (this.isFollowed) {
          // 取消关注
          await unfollow(UserID, this.userInfo.UID);
        } else {
          // 关注
          await followadd(UserID, this.userInfo.UID);
        }
        // 关注成功后，触发事件通知父组件
        this.$emit('follow-success');
        
      } 
      catch (error) {
        console.error('Error loading user information:', error);
        alert("关注失败，请稍后重试！");
      }
    },
    async getFansNum_(){
      try{
        const fanNum=await getFansNum();
        this.followersCount=fanNum;
      }catch{
        alert("请求粉丝数失败");
      }
    }
  },
  mounted() {
    this.getFansNum_();
  }
}
</script>

<style>
.information {
  display: flex;
  padding: 25px;
  border: 1px solid #ddd;
  border-radius: 8px;
  width: 100%;
  background-color: rgba(255, 255, 255, 0.33);  /* 设置白色半透明背景 */
}

.avatar-container {
  position: relative;
  margin-bottom: 5px;
  margin-left: 25px;
  margin-right: 35px;
}

.avatar {
  width: 95px;
  height: 95px;
  border-radius: 50%;
  /* background: #3578c9; */
  overflow: hidden;
  margin-bottom: 10px;
  margin-top: 15px;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-upload {
  display: none;
}

.follow-btn {
  margin-top: 10px;
  padding: 5px 10px;
  background: rgba(245, 221, 5, 0.845);
  border: none;
  color: black;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

.user-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 30px;
  margin-left: 10px;
  margin-right: 20px;
}

.username {
  font-size: 23px;
  margin-bottom: 25px;
}

.followers {
  font-size: 15px;
}

.username {
  font-weight: bold;
}

.edit-btn {
  margin-left: 10px;
  font-size: 14px;
  color: #4e90f0;
  background: none;
  border: none;
  cursor: pointer;
}

.edit-btn:hover {
  text-decoration: underline;
}

.information-right {
  flex: 2; /* 占据剩余空间 */
  padding: 10px; /* 减少内边距 */
  font-size: 14px; /* 减小字体大小 */
  line-height: 1.2; /* 减小行高 */
  height: 100px; /* 固定高度 */
  overflow: hidden; /* 超出部分隐藏 */
  text-align: left; /* 文本左对齐 */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1); /* 添加阴影 */
  border-radius: 8px; /* 圆角 */
  background-color: rgba(255, 255, 255, 0.432); /* 半透明背景 */
  margin-top:20px;
}

.signature {
  width: 90%;
  height: 70px;
  padding: 10px;
  margin-top: 50px;
  border-radius: 8px;
  border: 1px solid transparent;  /* 初始边框透明 */
  resize: none;
  font-size: 17px;
  background: none;
  transition: all 0.3s ease;  /* 添加过渡效果 */
}

.signature:focus {
  border-color: white;  /* 聚焦时显示边框 */
  background-color: rgba(255, 255, 255, 0.3);  /* 聚焦时背景更不透明 */
  outline: none;  /* 去掉默认的焦点轮廓 */
}

.signature:hover {
  border-color: white;  /* 鼠标悬停时显示边框 */
  background-color: rgba(255, 255, 255, 0.3);  /* 鼠标悬停时背景更不透明 */
}

.follow-btn.followed {
  background-color: #ccc; /* 已关注按钮样式 */
  color: #666;
}

</style>