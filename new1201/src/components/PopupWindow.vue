<!-- PopupWindow.vue -->
<template>
    <div class="popup-window">
      <div class="popup-content">
        <button class="close-button" @click="closePopup">×</button>
        <slot></slot>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'PopupWindow',
    methods: {
      closePopup() {
        this.$emit('close');
      },
    },
  };
  </script>
  
  <style scoped>
  .popup-window {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000; /* 确保弹出窗口在最上层 */
  }
  
  .popup-content {
    background: white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    position: relative; /* 使关闭按钮相对于这个容器定位 */
    width: 80%; /* 调整弹出窗口的宽度 */
    max-width: 500px; /* 设置最大宽度 */
  }
  
  .close-button {
    position: absolute;
    top: 10px; /* 调整到右上角 */
    right: 10px; /* 调整到右上角 */
    background-color: transparent;
    border: none;
    font-size: 48px;
    cursor: pointer;
    color: #d4d4d4; /* 设置按钮文字颜色 */
  }
  
  .close-button:hover {
    color: #484848; /* 鼠标悬停时改变颜色 */
  }
  </style>