<template>
    <div class="bgimg"></div>
    <div class="login-container">
        <h1>登录</h1>
        <!-- 默认的页面刷新行为被 .prevent 修饰符阻止 -->
        <!-- required表示这是一个必填字段，未填写时提交表单会提示错误 -->
        <form @submit.prevent="handleLogin">
            <div class="form-group">
                <label for="email">账号:</label>
                <input
                    type="email"
                    id="email"
                    v-model="email"
                    placeholder="请输入账号"
                    required
                />
            </div>
    
            <div class="form-group">
                <label for="password">密码:</label>
                <input
                    type="password"
                    id="password"
                    v-model="password"
                    placeholder="请输入密码"
                    required
                />
            </div>
  
            <button type="submit">登录</button>
        </form>

        <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    
        <p class="signup-link">
            还没有账号？ <router-link to="/RegistPage">点击此处注册</router-link>
        </p>

        <p>
            <router-link to="/ForgotPWPage">忘记密码?</router-link>
        </p>
    </div>
</template>
  
<script>
import { login } from '@/api.js';
    export default {
        name: "LoginPage",
        data() {
            return {
                email: "",
                password: "",
                errorMessage: "",
            };
        },
        setup(){
        },
        methods: {
            async handleLogin() {
                if (!this.email || !this.password) {
                    this.errorMessage = "账号或密码不能为空！";
                    return;
                }
        
                try {
                    const data=await login(this.email,this.password);
                    if (data) {
                        // 将用户信息存储在本地
                        localStorage.setItem('user', JSON.stringify(data));
                        this.$router.push("/HomePage"); // 登录成功后跳转到主页
                    } else {
                        this.errorMessage = data.message;
                        // this.errorMessage = "账号或密码错误";
                    }
                } catch (error) {
                    this.errorMessage = "账号或密码错误";
                }
            }
    
        }    
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
    .login-container {
        
        max-width: 400px;
        margin: 100px auto;
        padding: 40px;
        border: 1px solid #ddd;
        border-radius: 8px;
        /* background-color: #f9f9f9; */
        /* background-color: rgba(255, 255, 255, 0.5)!important;  */
        background-color: rgba(255, 255, 255, 0.5); /* 半透明白色背景 */
        text-align: center;
        position: relative;
        z-index: 1;
    }
    body body {
        /* background-color: red; */
        background-color: rgba(255, 255, 255, 0.5);
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
