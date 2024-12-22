<template>
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
                        alert("前端返回信息"+JSON.stringify(data)+"登录成功！");
                        // 将用户信息存储在本地
                        localStorage.setItem('user', JSON.stringify(data));
                        localStorage.setItem('userImage', data.UImage); // 存储用户头像URL
                        this.$router.push("/HomePage"); // 登录成功后跳转到主页
                    } else {
                        this.errorMessage = data.message;
                    }
                } catch (error) {
                    this.errorMessage = "网络错误，请稍后再试。";
                }
            }
    
        }    
    };
</script>

<style scoped>
    /* Same styles as before */
    .login-container {
        max-width: 400px;
        margin: 100px auto;
        padding: 40px;
        border: 1px solid #ddd;
        border-radius: 8px;
        background-color: #f9f9f9;
        text-align: center;
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
