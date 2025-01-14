<template>
    <div class="totalcontainer" v-if="item">
        <div class="body">
            <left>
                <!-- 图标 -->
                <img :src="`http://localhost:8082/avatars/${this.item.roomOwnerID}.jpg`" class="icon" />
                <div class="room-info">
                    <!-- 房间名称和房间号 -->
                    <div class="roomname">{{ item.roomName }} （{{ item.roomOwnerName }}的聊天室）</div>
                    <!-- 房间简介 -->
                    <div class="roomIntro">{{ item.roomIntro }}</div>
                </div>
            </left>
            <right>
                <img v-for="(i, index) in item.memberList" :key="index" :src="`http://localhost:8082/avatars/${i.memberID}.jpg`"/>
                <button class="join-button" @click="showModal = true">加入</button>
            </right>
        </div>
        <!-- 使用弹窗组件 -->
        <ConfirmModal
            :visible="showModal"
            :roomOwnerName="item.roomOwnerName"
            @confirm="handleJoin"
            @cancel="showModal = false"
        />
    </div>
</template>

<script>
import defaultImage from "@/assets/14efc6b3_E380544_1da91e8b.png";
import ConfirmModal from "@/components/PopWindowJoin.vue"; // 引入弹窗组件
import {joinChat} from "@/api"
export default {
    name: 'FindListItem',
    components: {
        ConfirmModal, // 注册弹窗组件
    },
    data() {
        return {
            defaultImage: defaultImage,
            showModal: false, // 控制弹窗显示
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
    },
    methods: {
      async handleJoin() {
        //本地获取uid
        let UserID = 0;
        try {
            const storedUser = localStorage.getItem('user');
            if (storedUser) {
                const user = JSON.parse(storedUser);
                UserID = user.UID;
            }
        } catch {
            alert("获取本地个人信息失败");
        }
        //请求加入房间
        let ok=false;
        ok=await joinChat(this.item.roomID,UserID);
        if(!ok){
            alert("加入房间失败！");
            return;
        }
        this.showModal = false; // 关闭弹窗
        // 跳转到 ChatPage
        this.$router.push({
          name: 'ChatPage',
          query: {
            roomID: this.item.roomID, // 传递房间 ID
          },
        });
      },
    },
}
</script>

 
<style scoped>
.totalcontainer {
  height: 100px;
  border-radius: 16px;
  padding: 0px;
  box-shadow: 0 4px 4px rgb(194, 179, 152);
  background-color: rgba(250, 235, 215, 0.61);
  padding-right: 15px;
  padding-left: 15px;
  background-position: center;
}

.body {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

left {
  height: 100px;
  max-width: 70%;
  background-color: transparent;
  display: flex;
  flex-direction: row;
  align-items: center; /* 垂直居中 */
  justify-content: flex-start; /* 水平靠左 */
}

/* 图标样式 */
.icon {
  width: 60px; /* 图标宽度 */
  height: 60px; /* 图标高度 */
  border-radius: 90%; /* 圆形图标 */
  margin-right: 10px; /* 图标与右侧内容的间距 */
}

/* 房间信息容器 */
.room-info {
  display: flex;
  flex-direction: column;
  justify-content: left; /* 垂直居中 */
}

/* 房间名称和房间号 */
.roomname {
  color: #e6a4b4;
  font-size: 1em;
  font-weight: 900;
  text-align: left;
}

/* 房间简介 */
.roomIntro {
  font-size: 0.9em;
  color: #e6a4b4;
  text-align: left;
  margin-top: 5px; /* 简介与房间名称的间距 */
}

right {
  background-color: transparent;
  display: flex;
  flex-direction: row;
  align-items: center;
}

/* 右侧图片样式 */
img {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background-color: aquamarine;
  margin-left: 5px;
}

.join-button {
background-color: #edd0de; /* 按钮背景颜色 */
color: white; /* 按钮文字颜色 */
border: none; /* 去掉边框 */
border-radius: 8px; /* 圆角 */
padding: 8px 16px; /* 内边距 */
font-size: 0.9em; /* 字体大小 */
font-weight: 600; /* 字体粗细 */
cursor: pointer; /* 鼠标悬停时显示手型 */
margin-left: 10px; /* 与右侧图片的间距 */
transition: background-color 0.3s ease; /* 背景颜色过渡效果 */
}

.join-button:hover {
background-color: #4a90e2; /* 鼠标悬停时的背景颜色 */
}

.ellipsis {
font-size: 2.5em; /* 字体大小 */
color: #666; /* 字体颜色 */
margin-left: 5px; /* 与头像的间距 */
vertical-align: middle; /* 垂直居中 */
}

</style>