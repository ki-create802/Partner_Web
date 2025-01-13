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
                <div v-if="currentTab === 'waiting'" class="wait-content">
                    <!-- 等搭中内容 -->
                    <p>等搭中的聊天框内容。</p>
                </div>
                <div v-if="currentTab === 'history'" class="his-content">
                    <!-- 历史内容 -->
                    <p>历史的聊天框内容。</p>
                </div>
            </div>
        </div>
    </div>

</template>

<script>
import TopBar from '../components/TopBar.vue';
import InformationBox from '../components/InformationBox.vue';
import GuideBar from '@/components/GuideBar.vue';

export default {
    name: 'HomePage',
    data() {
        return {
            currentTab: 'waiting' // 默认显示“等搭中”
        }
    },
    components: {
        GuideBar,
        InformationBox,
    },
}
</script>

<style scoped>
.background-layer {
    position: absolute;     /* 使用绝对定位 */
    top: 130px;                 /* 定位到顶部 */
    left: 0;                /* 从左边开始 */
    width: 100vw;           /* 占满全屏宽度 */
    height: 38%;           
    background-image: url('../assets/person_background.jpg');  
    background-size: cover;  /* 背景图自动适应尺寸 */
    background-position: center center;  /* 背景图居中 */
    background-repeat: no-repeat;  /* 不重复背景图 */
    background-size: cover;  /* 背景图自动适应尺寸 */
    background-position: center center;  /* 背景图居中 */
    background-repeat: no-repeat;  /* 不重复背景图 */
    z-index: -1;             /* 确保背景图位于所有内容的下方 */        
}
.topbar {
    width: 100vw;
}
.more-info {
    /* position: absolute; */
    display: flex;
    gap: 30px;
    background-color: #f1f2f3;
    width: 100vw;
}
.user-info {
    width: 100vw;
    margin-top: 90px;
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