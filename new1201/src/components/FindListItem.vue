<template>
    <body>
        <div class="body">
            <left>
                <a>房间号：{{ this.roomID }}</a>
                <a>{{ this.roomOwnerName }}的聊天室</a>
                <a class="roomIntro">简介：{{ this.roomIntro }}</a>
            </left>
            <right>
                <img :src="defaultImage" />
                <img :src="defaultImage" />
                <img :src="defaultImage" />
            </right>
        </div>
    </body>
</template>

<script>
import defaultImage from "@/assets/14efc6b3_E380544_1da91e8b.png";
export default {
    name: 'FindListItem',
    data(){
        return{
            defaultImage:defaultImage,
        };
    },
    props:{
        roomID:{
            type:Number,
            default:0,
        },
        roomIntro:{
            type:String,
            default:'这是默认简介'
        },
        roomOwnerID:{
            type:Number,
            default:0
        },
        roomOwnerName:{
            type:String,
            default:'默认楼主名'
        }
    }
}
</script>
    


<style scoped>
body{
    height: 100px;
    background-color:#ffffff8e ;
    border-radius: 16px;
    padding: 15px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);

}
.body{
    display: flex;
    flex-direction: row;
    justify-content: space-between; /* 子元素之间的间距 */
}
left{
    height: 100px;
    max-width: 70%;
    background-color:transparent;
    display: flex;
    flex-direction: column;
    align-items: flex-start; /* 左对齐 */
}
right{
    background-color:transparent;
    display: flex;
    flex-direction: row;
    align-items: center;
}
img{
    width: 50px; /* 设置图片的宽度 */
    height: 50px; /* 设置图片的高度 */
    border-radius: 50%;
    background-color: aquamarine;

}
.roomIntro{
    font-size: 16px;
    color: #ffffff;
    text-align: left;
}
</style>