<template>
  <GuideBar />
  <div class="mybody">
    <!-- 轮播图 -->
    <!-- <div>
      <CarouselImage :images="images" class="CarouselImage"/>
    </div> -->
    <!-- 一个图片 -->
    <img src="@/assets/image.png"  class="topimg" />

    <div class="lr">
      <div class="left">   
        
        <!-- 搜索框 -->
        <SearchBox class="searchbox" @search="handleSearch" :initialScope="selectedScope" />
        <!-- 模块 -->
        <Module />
        <!-- 火文 -->
        <div class="hot">
          <img src="@/assets/fire.png"  class="hot-icon" />
          <div class="hotlist">
            <FindListItem v-for="(item, index) in hotlist" :key="index" :item="item" class="hotitem" />
          </div>
        </div>
      </div>



      <div class="right">

        <div class="welcome">
          <h2>
            欢迎回来,
            <a class="lg-fg-blue light" href="/user/1078083" target="_blank">{{username}}</a>
          </h2>
          <div class="lg-index-calendar lg-fg-green dark"> 
            <span class="lg-punch-small">{{month}}</span>
            <span class="lg-punch-big">{{day}}</span>
            <span class="lg-punch-small">{{weekday}}</span>
          </div>
          <div class="am-g">
            <div class="am-u-sm-12 lg-small">
              <!-- 距...还剩...天 -->
              <div v-for="(item, index) in scheduleItems" :key="index">
                距 <strong>{{ item.content }}</strong> 还有 <strong>{{ getDaysUntil(item.date) }}天</strong><br>
              </div>
            </div>
            <strong style="display: block; margin-top: 20px;" v-if="showFortune" class="ffortune">{{ fortune }}</strong>
          </div>
        </div>

        <h1>我的行程</h1>
        <ACalendar class="calendar" ref="child" @update-schedule-items="updateScheduleItems" />

      </div>
    </div>

  </div>
</template>

<script>
import GuideBar from '@/components/GuideBar.vue';
import ACalendar from '@/components/ACalendar.vue';
import Module from '@/components/MoDule1.vue';
import SearchBox from '@/components/SearchBox.vue';
import FindListItem from '@/components/FindListItem.vue';
import { fetchHotData,search } from '@/api.js';

export default {
  name: 'HomePage',
  components: {
    GuideBar,
    ACalendar,
    // CarouselImage,
    SearchBox,
    FindListItem,
    Module
  },
  data() {
    return {
      images: [
        require('@/assets/轮播1.jpg'),
        require('@/assets/轮播2.jpg'),
        require('@/assets/轮播3.jpg'),
        require('@/assets/轮播4.jpg'),
        require('@/assets/轮播5.jpg'),
        require('@/assets/轮播6.jpg')
      ],
      hotlist: [] ,// 初始化 hotlist 数组
      month:"",
      day:"",
      weekeday:"",
      fortune: null, // 用来保存随机生成的运势
      showButton:true,
      showFortune:false,
      scheduleItems:[],
      username:"123",
    };
  },
  created() {
    this.fetchHotData_(); // 在组件创建时获取热门数据
  },

  mounted() {
    this.updateDate();
    this.loadFromLocalStorage();
  },


  methods: {
    updateScheduleItems(newScheduleItems) {
      this.scheduleItems = newScheduleItems;
    },
    async handleSearch(searchParams) {
      try {
        const { query, scope } = searchParams;
        const searchResults = await search(query, scope);
        this.$router.push({
          name: 'FindPage',
          query: {
            q: query,
            scope: scope,
            results: JSON.stringify(searchResults)
          }
        });
      } catch (error) {
        console.error('搜索请求失败:', error);
      }
    },
    async fetchHotData_() {
      try {
        const data = await fetchHotData();
        this.hotlist = data;
      } catch (error) {
        console.error('获取热门数据失败:', error);
      }
    },
    updateDate() {
      const date = new Date();
      const months = [
        '一月', '二月', '三月', '四月', '五月', '六月',
        '七月', '八月', '九月', '十月', '十一月', '十二月'
      ];
      const weekdays = [
        '星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'
      ];

      this.month = months[date.getMonth()];
      this.day = this.formatDay(date.getDate());
      this.weekday = weekdays[date.getDay()];
    },
    formatDay(day) {
      return day < 10 ? `0${day}` : day;
    },
    
    generateFortune() {
      // 运势列表
      const fortunes = [
        '今日运势：大吉！一切顺利。',
        '今日运势：中吉，谨慎行事。',
        '今日运势：小心，可能会有一些挑战。',
        '今日运势：啊哦，今天可能遇到一些阻碍。',
        '今日运势：平平，平淡的一天。',
      ];
      
      // 随机选择一条运势
      const randomIndex = Math.floor(Math.random() * fortunes.length);
      this.fortune = fortunes[randomIndex];
    },
    disappear(){
      this.showButton = false;  // 点击后隐藏按钮
      this.showFortune = true;      // 显示“开始”
    },

    getDaysUntil(targetDateStr) {  //距...行程还有多少天，用于洛谷
      // 将目标日期字符串转换为 Date 对象
      const targetDate = new Date(targetDateStr.replace(/\./g, "-")); 
      // 获取今天的日期 
      const today = new Date();
      // 计算日期差异（以毫秒为单位）
      const timeDifference = targetDate - today;
      //alert("日期差="+JSON.stringify(targetDateStr)+" - "+JSON.stringify(today));
      //alert("日期差："+JSON.stringify(timeDifference));
      // 将毫秒转换为天数
      const daysLeft = Math.ceil(timeDifference / (1000 * 60 * 60 * 24));
      return daysLeft;
    },
    loadFromLocalStorage() {
      console.log(111);
      // 从 localStorage 中获取值并更新 storedValue
      const user = JSON.parse(localStorage.getItem('user'));
      if (user) {
        console.log(222);
        this.username = user.UName;  // 如果存在，则赋值
        console.log(this.username)
      } else {
        console.log(333);
        this.username = 'No data found';  // 如果没有值，则设置默认值
      }
    }
  }
}
</script>



<style scoped>
.mybody {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  /*width: 1200px; /* 设置容器的宽度 */
  margin: 0 auto; /* 水平居中 */
}

.lr {
  display: flex; /* 启用 Flexbox 布局 */
  flex-direction: row; /* 子元素水平排列 */
  justify-content: space-between; /* 子元素之间的间距 */
  width: 100%; /* 占满父容器的宽度 */
  width: 1200px; /* 设置容器的宽度 */
  margin: 0 auto; /* 水平居中 */
}

.left {
  flex: 0 0 60%; /* 占据 60% 的宽度 */
  background-color: transparent;
  display: flex;
  flex-direction: column;
  align-items: center; /* 水平居中 */
  padding: 10px; /* 添加一些内边距 */
}

.right {
  flex: 0 0 30%; /* 占据 40% 的宽度 */
  display: flex;
  flex-direction: column;
  align-items: center; /* 水平居中 */
  /*box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);*/
}

.welcome {
  margin-top: 20%;
  margin-bottom: 30%;
}

.lg-punch-small, .lg-punch-big {
      display: inline-block; /* 使元素在同一行显示 */
      vertical-align: middle; /* 垂直居中对齐 */
}

.lg-punch-small {
  color: darkgreen; /* 墨绿色 */
  font-size: 2.0em; /* 较小 */
  writing-mode: vertical-rl; /* 竖着排列 */
  text-orientation: upright; /* 文字竖直排列 */
}


.lg-punch-big {

  color: darkgreen; /* 墨绿色 */
  font-size: 7em; /* 较大 */
  font-weight:900;
}

.sign{
  background-color: #5ba4f6;
  margin-top: 20px;
}

.ffortune{
  color: darkgreen; /* 墨绿色 */
  font-size: 1.5em; /* 较大 */
  font-weight:900;
}


.calendar {
  margin-top:0%;
  max-width: 400px;
  width: 90%;
}

.searchbox {
  padding: 10px;
  margin-top: 20px;
  margin-bottom: 20px;
}


.sections {
  display: grid;
  grid-template-columns: repeat(3, 1fr); /* 三列 */
  grid-template-rows: repeat(2, auto); /* 两行 */
  gap: 40px; /* 列和行之间的间距 */
  width: 100%; /* 使网格容器占满父容器的宽度 */
  margin-top: 20px; /* 添加一些上边距 */
}

.sections button {
  width: 100%; /* 使按钮占满网格单元格的宽度 */
  height: 80px;
  padding: 10px;
  border: none;
  background-color: #ffffff76;
  cursor: pointer;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  border-radius: 20px;
}

.sections button:hover {
  background-color: #f4feffaa;
}

.hot {
  display: flex;
  flex-direction: column; /* 纵向排列 */
  align-items: flex-start; /* 靠左对齐 */
  width: 100%; /* 使容器占满父容器的宽度 */
  margin-top: 20px; /* 添加一些上边距 */
}

.hotitem{
  margin-bottom: 10px;
}

.hotlist {
  width: 100%;
}

.hot-icon {
  width: 50px; /* 图标宽度 */
  height: 50px; /* 图标高度 */
  margin-right: 8px; /* 图标与右侧内容的间距 */
}

.topimg{
  width:80%;
  margin: 0 auto; /* 水平居中 */
}


</style>