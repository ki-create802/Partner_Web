<!-- src/components/ACalendar.vue -->
<template>
  <div class="calendar">
    <div class="header">
      <button @click="prevMonth">&lt;</button>
      <span>{{ currentMonthName }} {{ currentYear }}</span>
      <button @click="nextMonth">&gt;</button>
    </div>
    <div class="days">
      <div v-for="day in daysOfWeek" :key="day" class="day-name">{{ day }}</div>
    </div>
    <div class="dates">
      <div v-for="date in dates" :key="date.date" :class="['date', { selected: isSelected(date.date), today: isToday(date.date) }]">
        <button @click="selectDate(date.date)">{{ date.date }}</button>
      </div>
    </div>
    <ScheduleList :scheduleItems="scheduleItems" />
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
      scheduleItems: [
        { date: '10.01', content: '默认行程' },
        { date: '10.05', content: '默认行程' },
        { date: '10.10', content: '默认行程' },
        { date: '10.01', content: '默认行程' },
        { date: '10.05', content: '默认行程' },
        { date: '10.10', content: '默认行程' }

      ]
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
  border-radius: 4px;
  padding: 10px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.days {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  margin-bottom: 10px;
}

.day-name {
  text-align: center;
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
}

.date.selected button {
  background-color: #007bff;
  color: white;
  border-radius: 4px;
}

.date.today button {
  background-color: #ffc107;
  color: white;
  border-radius: 4px;
}
</style>