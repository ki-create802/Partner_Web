<template>
    <div class="topbar">
        <GuideBar />
    </div>
    <div class="user-info">
        <InformationBox />
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
                <!-- 加载状态 -->
                <div v-if="isLoading" class="loading">加载中...</div>
                <!-- 错误提示 -->
                <div v-if="error" class="error">{{ error }}</div>
                <!-- 等搭中内容 -->
                <div v-if="currentTab === 'waiting'" class="wait-content">
                    <FindListItem 
                        v-for="(chat,index) in pendingChats" 
                        :key="index" 
                        :item="chat" />
                </div>
                <!-- 历史内容 -->
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
import InformationBox from '../components/InformationBox.vue';
import GuideBar from '@/components/GuideBar.vue';
import FindListItem from '@/components/FindListItem.vue';
import HistoryItem from '@/components/HistoryItem.vue';
import axios from 'axios';
import {level} from "@/api";

export default {
    name: 'HomePage',
    data() {
        return {
            currentTab: 'waiting', // 默认显示“等搭中”
            pendingChats: [],
            historyChats: [],
            isLoading: false,
            error: null,
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
    components: {
        GuideBar,
        InformationBox,
        FindListItem,
        HistoryItem
    },
    methods: {
        async fetchPendingChats() {
            this.isLoading = true;
            this.error = null;
            try {
                const user = JSON.parse(localStorage.getItem('user'));
                const userID = user.UID;
                const response = await axios.post(`http://localhost:8082/chat/getPendingChats`, {
                    userID: userID
                });
                this.pendingChats = response.data.data.pendingChats;
            } catch (err) {
                this.error = '获取等搭中聊天室失败，请重试';
                console.error(err);
            } finally {
                this.isLoading = false;
            }
        },
        async fetchHistoryChats() {
            this.isLoading = true;
            this.error = null;
            try {
                const user = JSON.parse(localStorage.getItem('user'));
                const userID = user.UID;
                const response = await axios.post(`http://localhost:8082/chat/getHistoryChats`, {
                    userID: userID
                });
                this.historyChats = response.data.data.historyChats;
            } catch (err) {
                this.error = '获取历史聊天室失败，请重试';
                console.error(err);
            } finally {
                this.isLoading = false;
            }
        },
        async load(){
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
                const response = await level(UserID);  //自己alert
                
                this.n = response+1;
            } catch (error) {
                console.error('Error loading user information:', error);
                alert("请求粉丝信息失败，请稍后重试！");
            }
        }
    },
    watch: {
        currentTab(newTab) {
            if (newTab === 'waiting') {
                this.fetchPendingChats();
            } else if (newTab === 'history') {
                this.fetchHistoryChats();
            }
        },
    },
    mounted() {
        this.fetchPendingChats();
        this.load();
    },
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