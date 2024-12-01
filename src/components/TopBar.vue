<template>
    <div class="top-rectangle"></div>
    <el-menu class="topbar" router :default-active="this.$route.fullPath" mode="horizontal" :ellipsis="false">
        <div class="logo">
            <img src="../assets/logo.jpg" alt="Logo" class="logo">
        </div>
        <div class="navbar">
            <ul class="nav-items">
                <li v-for="(item, index) in noSubMenuItems" :key="index" class="nosub-item">
                    <div v-if="item.type === 'text'">
                        <a href="#" >{{ item.label }}</a>
                    </div>
                    <div v-if="item.type === 'image'">
                        <img style="width: 25px; height: auto;" :src="require(`../assets/${item.icon}`)" alt="icon" />
                    </div>
                </li>
                <li v-for="(item, index) in hasSubMenuItems" :key="index" :index="item.index" class="hassub-item">
                    <div>
                        <img style="width: 25px; height: auto;" :src="require(`../assets/${item.icon}`)"/>
                        <button @click="toggleDropdown(index)" class="dropdown-toggle">
                            <i :class="{'rotate': item.isDropdownVisible}">&#9660;</i> <!-- 向下箭头 -->
                        </button>
                        <div v-show="item.isDropdownVisible" class="dropdown">
                            <ul>
                                <li v-for="(subItem, subIndex) in subMenuItems(item.index)" 
                                    :key="subIndex" 
                                    :index="subItem.index" >>
                                    {{ subItem.label }}
                                </li>
                            </ul>
                        </div>
                    </div>
                </li>
            </ul>
        </div>
    </el-menu>
</template>


<script >
import { defineComponent, computed, reactive } from 'vue';
// //导入菜单选项配置文件
import config from '../config/config.json';

export default defineComponent({
    name: 'TopBar',
    setup() {
        // 使用 reactive 确保 menuItems 是响应式的
        const menuItems = reactive(config.menuItems);

        const hasSubMenuItems = computed(() => { //返回所有mainMenu为'hasSub'的菜单项
            return menuItems.filter(item => item.mainMenu === 'hasSub');
        });
        const noSubMenuItems = computed(() => { //返回所有mainMenu为'noSub'的菜单项
            return menuItems.filter(item => item.mainMenu === 'noSub');
        });
        const subMenuItems = (mainMenuIndex) => { //返回与指定mainMenuIndex匹配的子菜单项。
            return menuItems.filter(item => item.mainMenu === mainMenuIndex)
        };
        // 切换下拉框显示状态
        const toggleDropdown = (index) => {
            const item = hasSubMenuItems.value[index];
            console.log(item.index);
            item.isDropdownVisible = !item.isDropdownVisible; // 切换显示状态
            console.log(item.isDropdownVisible);
        };

        return {
            menuItems,
            hasSubMenuItems,
            noSubMenuItems,
            subMenuItems,
            toggleDropdown
        }
    }
    
});
</script>

<style lang="css" scoped>
.navbar-text{
    display: flex;
    flex-direction: row;
}
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
    height: 70px;
    width: 77%;
    position: relative; 
    left: 20%;   /* 使其从左边开始，20%的空白，使其居中 */
    /* 设置元素的边框圆角半径 */
    border-radius: 30px;  
    left: 10%;
    right: 10%;
    box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
    display: flex;
    justify-content: space-between;
}
.logo {
    width: 140px;
    height: auto;
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
    margin-right: 20px;
    margin-top: 10px;
}
a {
    text-decoration: none;  /* 去除链接的下划线 */
    color: inherit;         /* 继承父元素的文本颜色 */
}
.dropdown {
    position: absolute;
    top: 100%;
    background-color: #fff;
    border: 1px solid #ccc;
    padding: 10px;
    /* display: block; */
    width: 70px; 
    border-radius: 5px; /* 圆角边框 */
}
.dropdown ul {
    list-style: none; /* 移除列表前的默认符号 */
    padding: 0; 
    margin: 0;
}
/* 每一项（li） */
.dropdown li {
    font-size: 14px; 
    font-family: 'Arial', sans-serif; 
    font-weight: normal; 
    padding: 4px 0; 
    cursor: pointer; /* 设置鼠标悬停时的指针 */
    transition: background-color 0.3s; /* 添加过渡效果 */
}
/* 鼠标悬停时，改变背景颜色 */
.dropdown li:hover {
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