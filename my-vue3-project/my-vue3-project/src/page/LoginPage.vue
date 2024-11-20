<template>
    <h1>登录界面</h1>
    <input v-model="userID" placeholder="输入userID" />
    <button @click="toHomePage">登录</button>
    <button @click="toRegistPage">注册</button>
    <h5>输入的id：{{ userID }}</h5>
</template>

<script>
import router from '@/router';
import { inject } from 'vue';
export default{
    name:'LoginPage',
    methods:{
        toHomePage(){
            router.push("HomePage");
        },
        toRegistPage(){
            router.push("RegistPage");
        }
    },
    setup(){
        const userID=inject('userID');
        return{
            userID
        };
    }
}
</script>
