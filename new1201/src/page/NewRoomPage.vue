<template>
    <GuideBar />
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
                <option value="exercise">运动</option>
                <option value="fun">娱乐</option>
                <option value="car">拼车</option>
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
export default{
    name:'NewRoomPage',
    components:{
        GuideBar,
        ElDatePicker,
    },
    data(){
        return{
            roomname:"",
            roomType:"",
            people_num:"",
            selectedDate:"",
            //selectedDate:[],
            errorMessage: "", // 错误提示信息
        };
    },
    setup(){
    },
    methods: {
        async handleSubmit() {
            if (this.roomname && this.roomType && this.people_num && this.selectedDate) {
                alert(
                `聊天室 "${this.roomname}" 创建成功！类型: ${this.roomType}, 时间: ${this.selectedDate}`
                );
                this.$router.push("/ChatPage");
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

    .createroom-container {
        /* max-width: 600px; */
        width:50vw;
        /* height: 50vh; */
        margin: 50px auto;
        padding: 40px;
        border: 1px solid #ddd;
        border-radius: 8px;
        background-color: #f9f9f9;
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
