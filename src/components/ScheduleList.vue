<!-- src/components/ScheduleList.vue -->
<template>
    <div class="schedule-list">
      <div v-for="(item, index) in scheduleItems" :key="index" class="schedule-item">
        <div class="schedule-date" style="text-align: left;">{{ item.date }}</div>
        <div class="schedule-content" style="font-family: '宋体', SimSun, serif;text-align: left;">{{ item.content }}</div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'ScheduleList',
    props: {
      scheduleItems: {
        type: Array,
        required: true
      }
    }
  };
  </script>
  
  <style scoped>
  .schedule-list {
    margin-top: 20px;
    font-size: 150%;
  }
  
  .schedule-item {
    display: flex;
    flex-direction: column;
    margin-bottom: 10px;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  .schedule-date {
    font-weight: bold;
    margin-bottom: 5px;
  }
  
  .schedule-content {
    font-size: 0.9em;
  }
  </style>