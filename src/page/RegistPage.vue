<template>
    <h1>注册界面</h1>
    <button @click="toLoginPage">返回登录</button>
</template>

<script>
import router from '@/router';

export default{
    name:'RegistPage',
    methods:{
        toLoginPage(){
            router.push("/");
        }
    }
}
</script>
