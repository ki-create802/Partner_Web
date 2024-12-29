<!-- src/components/ACalendar.vue -->
<template>
  <div class="calendar">
    <div class="header">
      <button @click="prevMonth" class="prevbtn"></button>
      <span class="monthyear">{{ currentMonthName }} {{ currentYear }} </span>
      <button @click="nextMonth" class="nextbtn"></button>
    </div>
    <div class="days">
      <div v-for="day in daysOfWeek" :key="day" class="day-name">{{ day }}</div>
    </div>
    <div class="dates">
      <div v-for="date in dates" :key="date.date" :class="['date', { selected: isSelected(date.date), today: isToday(date.date) }]">
        <button @click="selectDate(date.date)">{{ date.date }}</button>
      </div>
    </div>
    <ScheduleList :scheduleItems="scheduleItems" class="list"/>
  </div>
</template>

<script>
import ScheduleList from './ScheduleList.vue';
import axios from 'axios';

export default {
  name: 'ACalendar',
  components: {
    ScheduleList
  },
  data() {
    return {
      currentDate: new Date(),
      selectedDate: null,
      daysOfWeek: ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'],
      scheduleItems: [],
    };
  },
  computed: {
    currentYear() {
      return this.currentDate.getFullYear();
    },
    currentMonth() {
      return this.currentDate.getMonth();
    },
    currentMonthName() {
      return this.currentDate.toLocaleString('default', { month: 'long' });
    },
    dates() {
      const dates = [];
      const firstDay = new Date(this.currentYear, this.currentMonth, 1);
      const lastDay = new Date(this.currentYear, this.currentMonth + 1, 0);
      const startDay = firstDay.getDay();
      const totalDays = lastDay.getDate();

      for (let i = 0; i < startDay; i++) {
        dates.push({ date: '', disabled: true });
      }

      for (let i = 1; i <= totalDays; i++) {
        dates.push({ date: i, disabled: false });
      }

      return dates;
    }
  },
  created() {
    this.fetchScheduleItems();
  },
  methods: {
    
    //获取个人日程信息  
    async fetchScheduleItems() {
      // 从本地存储中获取用户信息
      try {
        const user = JSON.parse(localStorage.getItem('user'));
        if (!user) {
          console.error('User information not found in local storage.');
          return;
        }
        const uid = user.UID;
      //向后端发送请求个人日程请求
      
        const response = await axios.get('http://localhost:3000/api/schedule', 
        {
          params:{
            Uid: uid,
          },
        });
        this.scheduleItems = response.data;
        this.$emit('update-schedule-items', this.scheduleItems);
      } catch (error) {
        console.error('There was an error fetching the schedule items!', error);
      }
    },
    prevMonth() {
      this.currentDate = new Date(this.currentYear, this.currentMonth - 1, 1);
    },
    nextMonth() {
      this.currentDate = new Date(this.currentYear, this.currentMonth + 1, 1);
    },
    selectDate(date) {
      this.selectedDate = new Date(this.currentYear, this.currentMonth, date);
      this.$emit('date-selected', this.selectedDate);
    },
    isSelected(date) {
      if (!this.selectedDate) return false;
      return (
        this.selectedDate.getFullYear() === this.currentYear &&
        this.selectedDate.getMonth() === this.currentMonth &&
        this.selectedDate.getDate() === date
      );
    },
    isToday(date) {
      const today = new Date();
      return (
        today.getFullYear() === this.currentYear &&
        today.getMonth() === this.currentMonth &&
        today.getDate() === date
      );
    }
  }
};
</script>

<style scoped>
.calendar {
  width: 300px;
  border: 1px solid #ccc;
  border-radius: 15px;
  padding: 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2); /* 阴影效果 */
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  border-radius: 10px;
  margin-top:20px;
  margin-left:20px;
  margin-right:20px;
}

.days {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  margin-bottom: 10px;
}

.day-name {
  text-align: center;
  color: darkgreen; /* 墨绿色 */
  font-size: 1.2em; /* 较大 */
  font-weight:900;
}

.dates {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 5px;
}

.date {
  text-align: center;
  
}

.date button {
  width: 100%;
  padding: 5px;
  border: none;
  background-color: transparent;
  cursor: pointer;
  border-radius: 10px;
  color: rgb(0, 0, 0); /* 墨绿色 */
  font-size: 1em; /* 较大 */
  font-weight:900;
}

.date.selected button {
  background-color: #007bff;
  color: white;
  border-radius: 10px;
}

.date.today button {
  background-color: #ffc107;
  color: white;
  border-radius: 10px;
}

.list{
  margin-left:10px;
  margin-right: 10px;
}

.monthyear{
  color: darkgreen; /* 墨绿色 */
  font-size: 1.5em; /* 较大 */
  font-weight:900;
}

.prevbtn{
  width: 40px; /* 按钮宽度 */
  height: 40px; /* 按钮高度 */
  border: none; /* 移除边框 */
  background-color: transparent; /* 透明背景 */
  background-size: cover; /* 背景图片覆盖按钮 */
  cursor: pointer; /* 鼠标指针为手型 */
  background-image: url('@/assets/prev-icon.jpg'); /* 设置背景图片 */
}

.nextbtn{
  width: 40px; /* 按钮宽度 */
  height: 40px; /* 按钮高度 */
  border: none; /* 移除边框 */
  background-color: transparent; /* 透明背景 */
  background-size: cover; /* 背景图片覆盖按钮 */
  cursor: pointer; /* 鼠标指针为手型 */
  background-image: url('@/assets/next-icon.jpg'); /* 设置背景图片 */
}

</style>