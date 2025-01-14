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
            <div class="medal1">
                <img src="../assets/medal1.png" class="img_m1" />
                <span class="m1"></span>
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
                    
                    <FindListItem />
                </div>
            </div>
        </div>
    </div>

</template>

<script>
import InformationBox from '../components/InformationBox.vue';
import GuideBar from '@/components/GuideBar.vue';
import FindListItem from '@/components/FindListItem.vue';
import axios from 'axios';

export default {
    name: 'HomePage',
    data() {
        return {
            currentTab: 'waiting', // 默认显示“等搭中”
            pendingChats: [],
            historyChats: [],
            isLoading: false,
            error: null,
        }
    },
    components: {
        GuideBar,
        InformationBox,
        FindListItem,
    },
    methods: {
        async fetchPendingChats() {
            this.isLoading = true;
            this.error = null;
            try {
                const user = JSON.parse(localStorage.getItem('user'));
                const userID = user.UID;
                console.log(userID);
                const response = await axios.post(`http://localhost:8082/chat/getPendingChats`, {
                    userID: userID
                });
                alert("接收到的response"+JSON.stringify(response));
                this.pendingChats=response.data.data.pendingChats;
                alert("接收到的9999:"+JSON.stringify(this.pendingChats));
                // this.pendingChats = response.data.data.pendingChats.map(chat => ({
                //     ...chat,
                //     memberList: chat.memberList || [] // 处理空的 memberList
                // }));
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
                console.log(userID);
                const response = await axios.post(`http://localhost:8082/chat/getHistoryChats`, {
                    userID: userID
                });
                alert("接收到的消息："+JSON.stringify(response.data.data.pendingchats));
                this.historyChats = response.data.data.historyChats.map(chat => ({
                    ...chat,
                    memberList: chat.memberList || [] // 处理空的 memberList
                }));
            } catch (err) {
                this.error = '获取历史聊天室失败，请重试';
                console.error(err);
            } finally {
                this.isLoading = false;
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
    },
}
</script>

<style scoped>
.background-layer {
    position: absolute;     /* 使用绝对定位 */
    top: 130px;                 /* 定位到顶部 */
    left: 0;                /* 从左边开始 */
    width: 100%;           /* 占满全屏宽度 */
    height: 32%;           
    background-image: url('../assets/person_background.jpg');  
    background-size: cover;  /* 背景图自动适应尺寸 */
    background-position: center center;  /* 背景图居中 */
    background-repeat: no-repeat;  /* 不重复背景图 */
    background-size: cover;  /* 背景图自动适应尺寸 */
    background-position: center center;  /* 背景图居中 */
    background-repeat: no-repeat;  /* 不重复背景图 */
    z-index: 0;                   
}
.more-info {
    margin-top: 8%;
    display: flex;
    gap: 30px;
    background-color: #f1f2f3;
}
.user-info {
    margin-top: 11.5%;
    z-index: 1;
    position: relative;
}
.achievement {
    height: 700px;
    width: 320px;
    background: white;
    margin-top: 15px;
    margin-left: 50px;
}
.medal1 {
    margin-left: 15px;
}
.img_m1 {
    height: 50px;
    margin-top: 20px;
}
.tabWidget {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    padding: 15px;
    margin-left: 35px;
}
.tabs {
    display: flex;
}
.tabs button {
    padding: 5px 8px;
    margin: 2px 5px;
    font-size: 16px;
    cursor: pointer;
    border: none;
    background-color: white;
    transition: background-color 0.3s;
}
.tabs button.active {
    background-color: #007bff;
    color: white;
}
.tab-content {
    width: 1300px;
    height: 650px;
    border: 1px solid #f1f2f3;
    background-color: white;
}
.wait-content {
    padding: 10px;
    margin: 10px;
    background: #d2030398;
}
.his-content {
    padding: 10px;
    margin: 10px;
    background: #d2030398;
}
</style>