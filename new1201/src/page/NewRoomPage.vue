<template>
    <GuideBar />
    <div class="bgimg"></div>
    <div class = "createroom-container">
        <form @submit.prevent="handleSubmit">
            <div class="form-group">
                <label for = "room_name">聊天室名称:</label>
                <input
                type = "room_name"
                id = "room_name"
                v-model = "roomname"
                placeholder="请输入聊天室名称"
                required
                />
            </div>
            <!-- 聊天室类型 -->
            <div class="form-group">
                <label for="room_type">活动板块:</label>
                <select id="room_type" v-model="roomType" required>
                <option disabled value="">请选择板块</option>
                <option value=1>运动</option>
                <option value=3>娱乐</option>
                <option value=2>拼车</option>
                <option value=4>拼单</option>
                <option value=5>约拍</option>
                <option value=6>其他</option>
                </select>
            </div>

            <div class="form-group">
                <label for = "people_num">活动人数:</label>
                <input
                    type = "people_num"
                    id="people_num"
                    v-model="people_num"
                    placeholder="请输入活动人数"
                    @input="validateNumberInput"
                    required
                />
            </div>

            <!-- 年月选择 -->
            <div class="form-group">
                <label for="room_date">活动日期:</label>
                <!-- <el-date-picker
                    v-model="selectedDate"
                    type="date"
                    placeholder="请选择日期"
                    format="YYYY-MM-DD"
                /> -->
                <el-date-picker
                    v-model="selectedDate"
                    type="date"
                    placeholder="请选择日期"
                    format="YYYY-MM-DD"
                    :clearable="true"
                    :editable="true"
                />
            </div>

            <div class="form-group">
                <label for = "remark">活动简介:</label>
                <textarea
                    id="remark"
                    v-model="remark"
                    placeholder="请输入活动的详细内容"
                    rows="4"
                    required
                />
            </div>

            <!-- 提交按钮 -->
            <button type="submit">创建聊天室</button>
        </form>
    </div>

</template>

<script>
import GuideBar from '../components/GuideBar.vue';
import { ElDatePicker } from "element-plus";
import { newRoom } from '@/api';
export default{
    name:'NewRoomPage',
    components:{
        GuideBar,
        ElDatePicker,
    },
    data(){
        return{
            roomname:"",
            roomType:0,
            people_num:"",
            selectedDate:"",
            //selectedDate:[],
            errorMessage: "", // 错误提示信息
        };
    },
    setup(){
    },
    // methods: {
    //     async handleSubmit() {
    //         if (this.roomname && this.roomType && this.people_num && this.selectedDate) {
    //             alert(
    //             `聊天室 "${this.roomname}" 创建成功！类型: ${this.roomType}, 时间: ${this.selectedDate}`
    //             );
    //             this.$router.push("/ChatPage");
    //         } else {
    //             alert("请填写所有字段！");
    //         }
    //     },
    //     validateNumberInput() {
    //         // 通过正则移除非数字字符
    //         this.people_num = this.people_num.replace(/\D/g, "");
    //     },
    // },
    methods: {
        formatDate(date) {
            if (!date) return ''; // 如果日期为空，返回空字符串
            const year = date.getFullYear();
            const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从 0 开始，需要 +1
            const day = String(date.getDate()).padStart(2, '0');
            return `${year}-${month}-${day}`;
        },
        async handleSubmit() {
            if (this.roomname && this.roomType && this.people_num && this.formatDate(this.selectedDate)) {
                let uid=0;
                try {
                    const storedUser = localStorage.getItem('user');
                    if (storedUser) {
                        const user = JSON.parse(storedUser);
                        uid = user.UID;
                    }
                } catch {
                    alert("获取本地个人信息失败");
                }
                let ok=false;
                ok=await newRoom(this.roomname,this.roomType,uid,this.people_num,this.formatDate(this.selectedDate),this.remark);
                if(ok)this.$router.push("/ChatPage");
                
            } else {
                alert("请填写所有字段！");
            }
        },
        validateNumberInput() {
            // 通过正则移除非数字字符
            this.people_num = this.people_num.replace(/\D/g, "");
        },
    },
};
</script>

<style scoped>

    .bgimg{
        /* width: 150px;
        height: 60px; */
        width: 100%;
        height: 100%;
        position: fixed; 
        top: 0;
        left: 0;
        background-image: url('~@/assets/login_bg.JPG');
        background-size: cover;
        background-position: center; /* 居中对齐背景图 */
        z-index: -1; /* 设置层级，使其在所有内容下方 */
    }
    .createroom-container {
        /* max-width: 600px; */
        width:50vw;
        /* height: 50vh; */
        margin: 50px auto;
        padding: 40px;
        border: 1px solid #ddd;
        border-radius: 8px;
        /* background-color: #f9f9f9; */
        background-color: rgba(255, 255, 255, 0.5); /* 半透明白色背景 */
        text-align: center;

        display:flex;
        flex-direction: column; /* 垂直方向排列子项 */
        justify-content: center;/* 子元素居中对齐 */
        align-items: center; /*垂直居中*/
        
    }

    .form-group {
        display: flex; /* 使用 Flexbox 布局 */
        align-items: center; /* 垂直对齐 */
        gap: 10px; /* 控制间距 */
        margin-bottom: 15px;
        width: 400px;
        
    }



    label {
        flex-shrink: 0; /* 确保 label 不会缩小 */
        font-weight: bold; /* 让文字更醒目 */
        min-width: 100px; /* 控制 label 的宽度，确保对齐 */
    }

    input,
    select{
        flex: 1; /* 输入框占用剩余空间 */
        padding: 8px;
        border: 1px solid #ccc;
        border-radius: 4px;
        font-size: 14px;
    }

    textarea {
        flex: 1; /* 占用剩余空间 */
        padding: 8px;
        border: 1px solid #ccc;
        border-radius: 4px;
        font-size: 14px;
        resize: vertical; /* 允许用户垂直调整大小 */
        width: 100%; /* 填满父容器宽度 */
        height: auto; /* 自动根据内容调整高度 */
    }

    button {
        width: 40%;
        padding: 10px;
        background-color: #007bff;
        color: white;
        border: none;
        border-radius: 4px;/*边框圆角*/
        font-size: 16px;
        cursor: pointer;
        /* justify-content: flex-end; 子元素靠底部对齐 */
        position:relative;
        bottom:0;
    }

    button:hover {
        background-color: #0056b3;
    }
</style>
