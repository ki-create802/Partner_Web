<template>
    <body>
        <button @click="toHomePage">首页</button>
        <button @click="toFindPage">找搭子</button>
        <button @click="toChatPage">聊天室</button>
        <button class="sayhi" @click="toPersonalPage">你好！{{ userID}}</button>
    </body>
</template>

<script>
import router from '@/router';
import { inject } from 'vue';

export default {
    name: 'GuideBar',
    setup(){
        const userID=inject("userID");
        return{
            userID,
        };
    },
    methods:{
        toHomePage(){
            router.push("HomePage");
        },
        toFindPage(){
            router.push("FindPage");
        },
        toChatPage(){
            router.push("ChatPage");
        },
        toPersonalPage(){
            router.push("PersonalPage");
        }

    }
}
</script>
    


<style scoped>
    body{
        margin-top: 5px;
        display: flex;
        justify-content: space-between; /* 子元素之间的间距 */
        align-items: center; /* 子元素垂直居中 */
        background-color: #96b7e9;
    }
    button{
        width: 100%;
        height: 80px;
        color:#ffffff;
        background-color: transparent;
        border: none;
        font-size: 200%;
    }
    button:hover{
        background-color: #c5d2e6;
    }
    .sayhi{
        font-size: 100%;
    }
</style>