<template>
    <div class="totalcontainer" v-if="item">
        <div class="body">
            <left>
                <a>房间号：{{ item.roomID }}</a>
                <a>{{ item.roomOwnerName }}的聊天室</a>
                <a class="roomIntro">{{ item.roomIntro }}</a>
            </left>
            <right>
                <img v-for="(image, index) in item.images" :key="index" :src="'http://localhost:3000/'+image" />
            </right>
        </div>
    </div>
</template>

<script>
import defaultImage from "@/assets/14efc6b3_E380544_1da91e8b.png";

export default {
    name: 'FindListItem',
    data() {
        return {
            defaultImage: defaultImage,
        };
    },
    props: {
        item: {
            type: Object,
            required: true,
            default: () => ({
                roomID: '',
                roomName: '',
                roomIntro: '',
                roomOwnerID:0,
                roomOwnerName:'默认房主',
                memberList:[],
                images:[]
            })
        }
    }
}
</script>

<style scoped>
.totalcontainer {
    height: 100px;
    background-color: #ffffff48;
    border-radius: 16px;
    padding: 15px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.body {
    display: flex;
    flex-direction: row;
    justify-content: space-between; /* 子元素之间的间距 */
}

left {
    height: 100px;
    max-width: 70%;
    background-color: transparent;
    display: flex;
    flex-direction: column;
    align-items: flex-start; /* 左对齐 */
}

right {
    background-color: transparent;
    display: flex;
    flex-direction: row;
    align-items: center;
}

img {
    width: 50px; /* 设置图片的宽度 */
    height: 50px; /* 设置图片的高度 */
    border-radius: 50%;
    background-color: aquamarine;
    margin-left: 5px; /* 图片之间的间距 */
}

.roomIntro {
    font-size: 16px;
    color: #a7a7a7;
    text-align: left;
}
</style>