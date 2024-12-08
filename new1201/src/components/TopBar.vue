<template>
    <div class="top-rectangle"></div>
    <div class="topbar">
        <div class="logo">
            <img src="../assets/logo.jpg" alt="Logo" class="logo">
        </div>
        <div class="navbar">
            <ul class="nav-items">
                <button @click="toHomePage">首页</button>
                <button @click="toFindPage">求搭列表</button>
                <button @click="toChatPage">我的聊天</button>
                <button @click="toAboutUsPage">关于我们</button>
                <div class="newARoom">
                    <img @click="toNewRoomPage" src='@/assets/create_groups.png' style="height: 40px;">
                </div>
                <div>
                    <img style="width: 40px;" src='@/assets/personal_setting.png'/>
                    <button @click="toggleDropdown" class="dropdown-toggle">
                        <i :class="{'rotate': isDropdownVisible}">&#9660;</i> <!-- 向下箭头 -->
                    </button>
                    <div v-show="isDropdownVisible" class="dropdown">
                        <ul>
                            <button @click="toPersonalPage">个人主页</button>
                            <button @click="ChangePass">修改密码</button>
                            <button @click="PersonExport">退出登录</button>
                        </ul>
                    </div>
                </div>
            </ul>
        </div>
    </div>
</template>


<script >
import { defineComponent, ref } from 'vue';
import router from '@/router';

export default defineComponent({
    name: 'TopBar',
    setup() {
        // 使用 ref 来定义响应式数据
        const isDropdownVisible = ref(false);
        // 切换下拉框显示状态
        const toggleDropdown = () => {
            isDropdownVisible.value = !isDropdownVisible.value; // 切换显示状态
            console.log(isDropdownVisible.value);
        };

        return {
            isDropdownVisible,
            toggleDropdown
        }
    },
    methods: {
        toHomePage() {
            router.push("HomePage");
        },
        toFindPage() {
            router.push("FindPage");
        },
        toChatPage() {
            router.push("ChatPage");
        },
        toPersonalPage() {
            router.push("PersonalPage");
        },
        toAboutUsPage() {
            router.push("AboutUsPage");
        },
        toNewRoomPage() {
            router.push("NewRoomPage");
        },
    }
    
});
</script>

<style lang="css" scoped>
.top-rectangle {
    /* position: fixed; */
    top: 0;
    height: 50px;
    width: 100%;
    background: -webkit-linear-gradient(top, hsla(207, 63%, 70%, 0.37), rgba(255, 255, 255,0)) no-repeat;
    z-index: 998;
    /* 确保长方形位于其他内容之上 */
}
.topbar {
    /* position: fixed; */
    top: -20px;
    height: 90px;
    width: 80%;
    position: relative; 
    /* 设置元素的边框圆角半径 */
    border-radius: 30px;  
    /* 使其从左边开始，10%的空白，使其居中 */
    left: 10%;
    box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
    display: flex;
    justify-content: space-between;
}
.logo {
    width: 180px;
    margin-left: 10px;
    margin-right: 20px;
}
.navbar{
    align-items: center;
    display: flex;
}
.nav-items {
    list-style: none;
    display: flex;
    gap: 50px;
    align-items: center; /* 垂直居中对齐 */
    margin-right: 20px;
    margin-top: 15px;
}
.nav-items button {
    background: none; 
    border: none;
    font-size: 20px;
    cursor: pointer; /* 设置鼠标悬停时的指针 */
    transition: background-color 0.3s; /* 添加过渡效果 */
}
/* 鼠标悬停时，改变背景颜色 */
.nav-items button:hover {
    color: #007BFF; /* 鼠标悬停时的字体颜色 */
    font-weight: bold; /* 鼠标悬停时加粗字体 */
}
.dropdown {
    position: absolute;
    top: 100%;
    background-color: #fff;
    border: 1px solid #ccc;
    padding: 10px;
    width: 70px; 
    border-radius: 5px; /* 圆角边框 */
}
.dropdown ul {
    list-style: none; /* 移除列表前的默认符号 */
    padding: 0; 
    margin: 0;
}
/* 每一项（li） */
.dropdown button {
    font-size: 16px; 
    font-family: 'Arial', sans-serif; 
    font-weight: normal;
    background: none; 
    border: none;
    padding: 4px 0; 
    cursor: pointer; /* 设置鼠标悬停时的指针 */
    transition: background-color 0.3s; /* 添加过渡效果 */
}
/* 鼠标悬停时，改变背景颜色 */
.dropdown button:hover {
    background-color: #f0f0f0; /* 设置鼠标悬停时的背景色 */
    color: #007BFF; /* 鼠标悬停时的字体颜色 */
    font-weight: bold; /* 鼠标悬停时加粗字体 */
}
/* 控制 dropdown-toggle 按钮样式 */
.dropdown-toggle {
    background: none;
    border: none;
    cursor: pointer; /* 鼠标样式为指针 */
    padding: 3px; 
    font-size: 14px;
    color: #333; 
    transition: transform 0.3s; /* 过渡效果 */
}
/* 当下拉框显示时，旋转箭头 */
.dropdown-toggle i {
    transition: transform 0.3s ease;
}
.dropdown-toggle i.rotate {
    transform: rotate(180deg); /* 旋转180度 */
}
/* 鼠标悬停时改变箭头颜色 */
.dropdown-toggle:hover i {
    color: #007BFF;
}
</style>