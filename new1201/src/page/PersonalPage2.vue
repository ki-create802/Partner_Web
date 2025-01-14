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
import InformationBox2 from '../components/InformationBox2.vue';
import {otherInform} from "@/api";
import {fansss} from "@/api";
import { whetherfollow } from '@/api.js';


export default {
    name: 'HomePage',
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

        }
    },
    created() {
        this.UID = this.$route.query.UID; // 获取传递的 UID
        console.log('接收到的 UID:', this.UID);
        // 你可以在这里根据 UID 发起请求，获取用户详细信息
    },
    mounted() {
        this.loadInform(); // 页面加载时调用
    },

    components: {
        GuideBar,
        InformationBox2,
    },
    methods:{
        async loadInform(){
            //请求个人信息
            try {
                const response = await otherInform(this.UID);
                //alert("888")
                //alert("收到的赋值前response："+JSON.stringify(response))
                this.userInfo = response;
                //alert("收到的response："+JSON.stringify(this.userInfo))
            } catch (error) {
                console.error('Error loading user information:', error);
                //alert("请求个人信息失败，请稍后重试！");
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


        }
    }
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
    display: flex;
    gap: 30px;
    background-color: #f1f2f3;
}
.user-info {
    margin-top: 90px;
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