<template>
    <div class="topbar">
        <GuideBar />
    </div>
    <div class="user-info">
        <InformationBox />
    </div>
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
import GuideBar from '../components/GuideBar.vue';
import InformationBox from '../components/InformationBox.vue';

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
.more-info {
    display: flex;
    gap: 20px;
}
.user-info {
    margin-top: 0;
}
.achievement {
    height: 500px;
    width: 250px;
    background: #ccc;
    margin-top: 15px;
    margin-left: 8px;
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
    margin-left: 10px;
}
.tabs {
    display: flex;
}
.tabs button {
    padding: 2px 5px;
    margin: 0 5px;
    font-size: 13px;
    cursor: pointer;
    border: none;
    background-color: #f4f4f4;
    transition: background-color 0.3s;
}
.tabs button.active {
    background-color: #007bff;
    color: white;
}
.tab-content {
    width: 1000px;
    height: 575px;
    border: 1.5px solid #000000;
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