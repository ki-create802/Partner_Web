<template>
    <GuideBar />
    <h1>{{userID}}的个人主页</h1>
</template>

<script>
import GuideBar from '@/components/GuideBar.vue';
import { inject } from 'vue';

export default{
    name:'PersonalPage',
    components:{
        GuideBar,
    },
    setup(){
        const userID=inject("userID");
        return{
            userID
        };
    }
}
</script>
