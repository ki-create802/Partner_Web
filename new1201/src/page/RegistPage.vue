<template>
    <div class="bgimg"></div>
    <div class="signup-container">
        <h1>注册您的账号</h1>
        <form @submit.prevent="handleSignUp">
            <div class="form-group">
                <label for="email">邮箱:</label>
                <input
                    type="email"
                    id="email"
                    v-model="email"
                    placeholder="请输入您的邮箱"
                    required
                />
            </div>

            <div class = "form-group">
                <label for = "username">用户名：</label>
                <input
                    type="username"
                    id="username"
                    v-model="username"
                    placeholder="请输入您的用户名"
                    required
                />
            </div>
  
            <div class="form-group">
                <label for="password">密码:</label>
                <input
                    type="password"
                    id="password"
                    v-model="password"
                    placeholder="请输入您的密码"
                    required
                />
            </div>
  
            <div class="form-group">
                <label for="confirm-password">确认密码:</label>
                <input
                    type="password"
                    id="confirm-password"
                    v-model="confirmPassword"
                    placeholder="请再次确认您的密码"
                    required
                />
            </div>
    
            <button type="submit">注册</button>
        </form>
        <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
        <p class="login-link">
            已经有账号？ <router-link to="/">点击此处登录</router-link>
        </p>
    </div>
</template>
  
  <script>
  import {signUp} from '@/api';
  export default {
    name: "UserSignUp",
    data() {
        return {
            email: "",
            username:"",
            password: "",
            confirmPassword: "",
            errorMessage: "",
        };
    },
    methods: {
      async handleSignUp() {
        if (!this.email || !this.username || !this.password || !this.confirmPassword) {
            this.errorMessage = "请输入所有字段";
            return;
        }
        if (this.password !== this.confirmPassword) {
            this.errorMessage = "两次输入的密码不匹配";
            return;
        }
        try{
            let ok=await signUp(this.username,this.email,this.password);
            if(ok){
                alert("注册成功");
                this.errorMessage = "";
            }
        }catch{
            this.errorMessage = "注册失败";
        }
      },
    },
  };
  </script>
  
<style scoped>
  /* Add styles similar to Login.vue */
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
    .signup-container {
        max-width: 400px;
        margin: 180px auto;
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
    
    .form-group {
        margin-bottom: 15px;
        text-align: left;
    }
        
    label {
        display: block;
        margin-bottom: 5px;
        font-weight: bold;
    }
    
    input {
        width: 100%;
        padding: 8px;
        margin-bottom: 10px;
        border: 1px solid #ccc;
        border-radius: 4px;
        font-size: 14px;
    }
  
    button {
        padding: 10px 20px;
        background-color: #007bff;
        color: #fff;
        border: none;
        border-radius: 4px;
        font-size: 16px;
        cursor: pointer;
    }
    
    button:hover {
        background-color: #0056b3;
    }
  
    .error-message {
        color: red;
        margin-top: 10px;
    }
</style>
  