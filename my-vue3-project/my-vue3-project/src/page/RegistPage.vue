<template>
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
  export default {
    name: "UserSignUp",
    data() {
        return {
            email: "",
            password: "",
            confirmPassword: "",
            errorMessage: "",
        };
    },
    methods: {
      handleSignUp() {
        if (!this.email || !this.password || !this.confirmPassword) {
            this.errorMessage = "Please fill in all fields.";
            return;
        }
        if (this.password !== this.confirmPassword) {
            this.errorMessage = "Passwords do not match.";
            return;
        }
  
        alert("Sign-up successful!");
        this.errorMessage = "";
      },
    },
  };
  </script>
  
<style scoped>
  /* Add styles similar to Login.vue */
    .signup-container {
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
  