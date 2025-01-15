<template>
    <div class="topbar">
        <GuideBar />
    </div>
    <div class="user-info" >
        <InformationBox2 :userInfo="userInfo" :fan="fan" :isFollowed="isFollowed" @follow-success="loadInform"  />
    </div>
    <!-- 添加背景图层 -->
    <div class="background-layer"></div>
    <div class="more-info">
        <div class="achievement">
           <!-- 动态渲染前 n 个 item -->
            <div v-for="(item, index) in items.slice(0, n)" :key="index" class="achievement-item">
                <img :src="item.icon" alt="icon" class="achievement-icon" />
                <span class="achievement-text">{{ item.text }}</span>
            </div>
        </div>
        <div class="tabWidget">
            <div class="tabs">
                <button 
                    :class="{ active: currentTab === 'waiting' }" 
                    @click="currentTab = 'waiting'">
                    等搭中
                </button>
                <button 
                    :class="{ active: currentTab === 'history' }" 
                    @click="currentTab = 'history'">
                    历史
                </button>
            </div>
            <div class="tab-content">
                <div v-if="currentTab === 'waiting'" class="wait-content">
                    <FindListItem 
                        v-for="(chat,index) in pendingChats" 
                        :key="index" 
                        :item="chat" />
                </div>
                <div v-if="currentTab === 'history'" class="his-content">
                    <HistoryItem
                        v-for="(chat,index) in historyChats" 
                        :key="index" 
                        :item="chat" />
                </div>
            </div>
        </div>
    </div>

</template>

<script>
import GuideBar from '../components/GuideBar.vue';
import InformationBox2 from '../components/InformationBox2.vue';
import {otherInform} from "@/api";
import {fansss} from "@/api";
import { whetherfollow,getHistory,getPending } from '@/api.js';
import FindListItem from '@/components/FindListItem.vue';
import {level} from "@/api";
import HistoryItem from '@/components/HistoryItem.vue';


export default {
    name: 'personalPage2',
    data() {
        return {
            currentTab: 'waiting', // 默认显示“等搭中”,
            userInfo: {
                UID: '',
                UName: '',
                Password: '',
                Email: '',
                URemark: '',
                UImage: ''
            },
            fan:null,
            isFollowed:false, //关注状态
            UID:0,
            pendingChats: [],
            historyChats: [],

            n: 3, // 控制显示的块数量
            items: [ // 定义每个块的图标和文本
                { icon: require('@/assets/medal1.png'), text: 'Lv1.搭子探索者' },
                { icon: require('@/assets/medal2.png'), text: 'Lv2.初级搭子体验者' },
                { icon: require('@/assets/medal3.png'), text: 'Lv3.中级搭子爱好者' },
                { icon: require('@/assets/medal4.png'), text: 'Lv4.高级搭子行家' },
                { icon: require('@/assets/medal5.png'), text: 'Lv5.搭子大师' },
            ],
        }
    },
    created() {
        this.UID = this.$route.query.UID; // 获取传递的 UID
       
    },
    mounted() {
        this.loadInform(); // 页面加载时调用alert
    },

    components: {
        GuideBar,
        InformationBox2,
        FindListItem,
        HistoryItem
    }, 
    methods:{
        async fetchPendingChats() {
            try{
                const data=await getPending(this.UID);
                this.pendingChats=data;
            }catch{
                alert("获取等搭中……信息失败");
            }
        },
        async fetchHistoryChats() {
            try{
                const data=await getHistory(this.UID);
                this.historyChats=data;
            }catch{
                alert("获取信息失败");
            }
        },
        async loadInform(){
            this.fetchPendingChats();
            this.fetchHistoryChats();
            //请求个人信息
            try {
                const response = await otherInform(this.UID);
                this.userInfo = response;
            } catch (error) {
                console.error('Error loading user information:', error);
            }

            //请求关注人数
            try {
                //alert("fan888")
                const response = await fansss(this.UID);  //别人
                //alert("收到的赋值前response："+JSON.stringify(response))
                this.fan = response;
                //alert("收到的response："+JSON.stringify(this.fan))
            } catch (error) {
                console.error('Error loading user information:', error);
                alert("请求粉丝信息失败，请稍后重试！");
            }

            //请求关注状态
            try {
                //获取个人id
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
                //alert("开始请求关注状态")
                const followStatus = await whetherfollow(UserID, this.UID); // 获取关注状态
                this.isFollowed = followStatus; // 更新关注状态
                //alert("isFollow:" +this.isFollowed);
            } catch (error) {
                console.error('Error loading follow status:', error);
            }


            //请求徽章
            try {
                const response = await level(Number(this.UID));  //人家
                this.n = response+1;
            } catch (error) {
                console.error('Error loading user information:', error);
                alert("请求粉丝信息失败，请稍后重试！");
            }
            


        }
    }
}
</script>

<style scoped>
/* 背景图层 */
.background-layer {
    position: absolute;
    top: 130px;
    left: 0;
    width: 100%;
    height: 32%;
    background-image: url('../assets/person_background.jpg');
    background-size: cover;
    background-position: center center;
    background-repeat: no-repeat;
    opacity: 0.8; /* 降低背景图透明度 */
    z-index: 0;
}

/* 整体布局 */
.more-info {
    display: flex;
    gap: 30px;
    background-color: #ffffea;
    border-radius: 15px; /* 圆角 */
    margin: 20px auto; /* 居中 */
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1); /* 添加阴影 */
}

.user-info {
    margin-top: 11.5%;
    z-index: 1;
    position: relative;
}

/* 成就部分 */
.achievement {
    height: 700px;
    width: 320px;
    background: white;
    margin-top: 15px;
    margin-left: 50px;
    border-radius: 15px; /* 圆角 */
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1); /* 添加阴影 */
    padding: 20px; /* 内边距 */
}

.medal1 {
    margin-left: 15px;
}

.img_m1 {
    height: 50px;
    margin-top: 20px;
}

/* 选项卡部分 */
.tabWidget {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    padding: 20px;
    margin-left: 35px;
    background-color: rgba(255, 255, 255, 0.8);
    border-radius: 15px; /* 圆角 */
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1); /* 添加阴影 */
    flex-grow: 1; /* 填充剩余空间 */
}

.tabs {
    display: flex;
    gap: 10px; /* 按钮间距 */
    margin-bottom: 20px; /* 与内容间距 */
}

.tabs button {
    padding: 10px 20px;
    font-size: 16px;
    cursor: pointer;
    border: none;
    background-color: #f1f1f1;
    border-radius: 8px; /* 圆角 */
    transition: background-color 0.3s, transform 0.2s;
}

.tabs button:hover {
    background-color: #007bff;
    color: white;
    transform: translateY(-2px); /* 悬停时上移 */
}

.tabs button.active {
    background-color: #007bff;
    color: white;
}

/* 内容区域 */
.tab-content {
    width: 100%;
    height: 650px;
    border: 1px solid #f1f2f3;
    background-color: rgba(255, 255, 255, 0.8);
    border-radius: 15px; /* 圆角 */
    padding: 20px; /* 内边距 */
    overflow-y: auto; /* 允许滚动 */
}

.wait-content, .his-content {
    display: flex;
    flex-direction: column;
    gap: 15px; /* 调整间距大小 */
}

/* FindListItem 组件样式 */
.find-list-item {
    background-color: white;
    padding: 15px;
    border-radius: 10px; /* 圆角 */
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1); /* 添加阴影 */
    transition: transform 0.2s, box-shadow 0.2s;
}

.find-list-item:hover {
    transform: translateY(-2px); /* 悬停时上移 */
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2); /* 悬停时阴影加深 */
}

.achievement {
    height: 720px;
    width: 19%;
    background-image: url('../assets/medal_background.png');
    background-size: cover;
    background-position: center center;
    margin-top: 40px;
    margin-left: 50px;
    display: flex; /* 新增：使用 flex 布局 */
    flex-direction: column; /* 新增：垂直排列 */
    gap: 10px; /* 新增：设置 item 之间的间距 */
    padding: 20px; /* 新增：内边距 */
}

.achievement-item {
    display: flex; /* 新增：使用 flex 布局 */
    align-items: center; /* 新增：垂直居中 */
    gap: 10px; /* 新增：图标和文字之间的间距 */
}

.achievement-icon {
    width: 40px; /* 新增：图标宽度 */
    height: 40px; /* 新增：图标高度 */
}

.achievement-text {
    font-size: 14px; /* 新增：文字大小 */
    color: #333; /* 新增：文字颜色 */
}

</style>