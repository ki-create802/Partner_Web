<!-- src/components/ScheduleList.vue -->
<template>
  <div class="schedule-list">
    <div v-for="(item, index) in scheduleItems" :key="index" class="schedule-item">
      <div class="wrap">
        <img src="@/assets/clock.jpg"  class="clock" />
        <div class="schedule-date" style="text-align: left;" >{{ item.date }}</div>
      </div>
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
  margin-bottom: 5px;
  color: rgb(161, 98, 153); /* 墨绿色 */
  font-size: 0.9em; /* 较大 */
  font-weight:900;
}

.schedule-content {
  font-size: 0.9em;
  font-weight:900;
}

.clock{
  width: 40px; /* 按钮宽度 */
  height: 40px; /* 按钮高度 */
}

.wrap{
  display: flex; /* 启用 Flexbox 布局 */
  flex-direction: row; /* 子元素水平排列（默认值，可省略） */
  gap: 10px; /* 子元素之间的间距 */
}

</style>