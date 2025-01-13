<template>
    <div class="information">
      <div class="avatar-container">
        <div class="avatar">
          <img :src="avatarUrl" alt="头像" id="avatar-img">
        </div>
        <button class="follow-btn" @click="triggerFollow">+关注</button>
      </div>
      <div class="user-info">
        <div class="username">
          <span id="username-text">{{ username }}</span>
        </div>
        <div class="followers">
          <span id="followers-text">粉丝数: {{ followersCount }}</span>
        </div>
      </div>
      <div class="information-right">
        <textarea 
          class="signature" 
          placeholder="编辑个性签名..." 
          v-model="signature"
          readonly>   <!--添加不可更改属性-->   
        </textarea>  
      </div>
    </div>
  </template>
  
  <script>
//   import axios from 'axios';
  import { getFansNum } from '@/api.js';
  
  export default {
    name: 'InformationBox',
    data() {
      return{
        avatarUrl:require("@/assets/avatar.png"), // 默认头像
        username: JSON.parse(localStorage.getItem('user')).UName||'默认用户名',
        followersCount: 2000,
        signature: JSON.parse(localStorage.getItem('user')).URemark ||'默认个性签名',
      };
    },
    methods: {
      triggerFollow() {
       
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
    background-color: rgba(255, 255, 255, 0.1);  /* 设置白色半透明背景 */
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
    flex: 2;
    padding-left: 20px;
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
  </style>