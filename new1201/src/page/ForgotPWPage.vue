<template>
  <div class="bgimg"></div>
    <div class="forgot-password-container">
      <!-- 判断是否显示验证码及邮箱表单 -->
      <div v-if="!showResetPasswordModal">
        <h1>重新设置密码</h1>
        <p>输入您的邮箱以获取验证码</p>
  
        <!-- 邮箱输入框及发送验证码按钮 -->
        <form @submit.prevent="handleVerifyCode">
          <div class="email-container">
            <input
              type="email"
              v-model="email"
              placeholder="请输入您的邮箱"
              required
            />
            <!-- <button type="button" @click="sendVerificationCode" :disabled="isCodeSent">
              {{ isCodeSent ? '${countdown}s后重新发送' : "发送验证码" }}
            </button> -->
            <button type="button" @click="sendVerificationCode" :disabled="isCodeSent">
              {{ isCodeSent ? `${countdown}s后重新发送` : "发送验证码" }}
            </button>
  
            
          </div>
  
          <p class="login-link">
              <router-link to="/">点击此处返回登录</router-link>
            </p>
  
          <!-- 验证码输入框 -->
          <div v-if="isCodeSent" class="form-group">
          <!--div  class="form-group"--> 
            <input
              type="text"
              v-model="verificationCode"
              placeholder="输入验证码"
              required
            />
          </div>
  
          <p v-if="isCodeSent">验证码已经发送到您的邮箱</p>
          <button type="submit" v-if="isCodeSent">确认</button>
        </form>
      </div>
  
      <!-- 重设密码弹窗 -->
      <div v-if="showResetPasswordModal" class="reset-password-modal">
        <h2>请重新设置您的密码</h2>
        <form @submit.prevent="resetPassword">
          <div class="form-group">
            <label for="newPassword">新的密码：</label>
            <input
              type="password"
              id="newPassword"
              v-model="newPassword"
              placeholder="请输入密码"
              required
            />
          </div>
          <div class="form-group">
            <label for="confirmPassword">确认密码：</label>
            <input
              type="password"
              id="confirmPassword"
              v-model="confirmPassword"
              placeholder="请再次确认密码"
              required
            />
          </div>
          <button type="submit">确认</button>
        </form>
      </div>
    </div>
  </template>
  
  <script>
  import { checkCode } from '@/api';
  import {forgetPW} from '../api';
  import {resetPW} from '../api';
  export default {
    name: "ForgotPWPage",
    data() {
      return {
        email: "",
        isCodeSent: false,
        countdown: 0, // 倒计时初始值
        verificationCode: "",
        showResetPasswordModal: false,
        newPassword: "",
        confirmPassword: "",
      };
    },
    methods: {
      async sendVerificationCode() {
        if (this.email) {
          let ok=false;
          ok =await forgetPW(this.email);
          if(!ok)return;
          alert("验证码已经发送到邮箱");
          this.isCodeSent = true; // 禁用按钮
          this.countdown = 60; // 设置倒计时初始值
          this.startCountdown(); // 开始倒计时
        } else {
          alert("请输入邮箱");
        }
      },
        // if (this.email) {
        //   this.isCodeSent = true; // 模拟验证码发送逻辑
        //   alert("Verification code sent to your email.");
        // } else {
        //   alert("Please enter a valid email address.");
        // }
        // // 触发发送验证码逻辑
        // console.log(`Sending verification code to ${this.email}`);
        
        // // 开始倒计时
        // this.isCodeSent = true;
        // this.startCountdown();
      
      startCountdown() {
        const interval = setInterval(() => {
          this.countdown -= 1; // 每秒减少倒计时
          if (this.countdown <= 0) {
            clearInterval(interval); // 倒计时结束，清除计时器
            this.isCodeSent = false; // 恢复按钮可用状态
          }
        }, 1000); // 每秒执行一次
      },
      async handleVerifyCode() {
        try{
          const ok=await checkCode(this.email,this.verificationCode)
          if (ok) {
            this.showResetPasswordModal = true;
            this.isCodeSent = false; // 确保隐藏验证码和邮箱表单
          } else {
            alert("验证码错误");
          }
        }catch{
          console("验证错误");
        }
      },
      async resetPassword() {
        if (this.newPassword === this.confirmPassword) {
          let ok=false;
          ok=await resetPW(this.email,this.newPassword);
          if(ok){
            alert("密码重置成功");
            this.showResetPasswordModal = false; 
          }
        } else {
          alert("两次输入的密码不一样。");
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .bgimg{
        /* width: 150px;
        height: 60px; */
        width: 100%;
        height: 100%;
        position: fixed; 
        top: 0;
        left: 0;
        background-image: url('~@/assets/login_bg.JPG');
        background-size: cover;
        background-position: center; /* 居中对齐背景图 */
        z-index: -1; /* 设置层级，使其在所有内容下方 */
    }
  .forgot-password-container {
    max-width: 400px;
    margin: 100px auto;
    padding: 40px;
    border: 1px solid #ddd;
    border-radius: 8px;
    /* background-color: #f9f9f9;
    text-align: center; */
    background-color: rgba(255, 255, 255, 0.5); /* 半透明白色背景 */
    text-align: center;
    position: relative;
    z-index: 1;
  }
  
  h1 {
    margin-bottom: 20px;
  }
  
  p {
    margin-bottom: 20px;
    font-size: 14px;
    color: #555;
  }
  
  .email-container {
    display: flex;
    gap: 10px;
    align-items: center;
    margin-bottom: 15px;
  }
  
  input[type="email"],
  input[type="text"]{
    flex: 1;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 14px;
  }
  
  button {
    padding: 8px 12px;
    background-color: #007bff;
    color: #fff;
    border: none;
    border-radius: 4px;
    font-size: 14px;
    cursor: pointer;
  }
  
  button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }
  
  button:hover:enabled {
    background-color: #0056b3;
  }
  
  .reset-password-modal {
    margin-top: 20px;
    padding: 40px;
    border: 1px solid #ddd;
    border-radius: 8px;
    /* background-color: #fefefe; */
    background-color: rgba(255, 255, 255, 0.5);
    text-align: left;
  }
  
  .reset-password-modal h2 {
    text-align: center;
  }
  
  .form-group {
      margin-bottom: 15px;
      /* text-align: left; */
  }

  label {
      display: block;
      margin-bottom: 5px;
      font-weight: bold;
  }
  input[type="password"] {
    width: 100%;
    padding: 8px;
    margin-bottom: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 14px;
  }
  </style>
  