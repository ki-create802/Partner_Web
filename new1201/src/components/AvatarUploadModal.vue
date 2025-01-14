<template>
    <!-- 弹窗 -->
    <div v-if="isModalOpen" class="modal">
      <div class="modal-content">
        <!-- 关闭按钮 -->
        <span class="close" @click="closeModal">&times;</span>
        <h2>上传头像</h2>
        <!-- 文件选择器 -->
        <input type="file" @change="handleFileChange" accept="image/*" />
        <!-- 图片预览 -->
        <div v-if="imageUrl" class="image-preview">
          <img :src="imageUrl" alt="头像预览" />
        </div>
        <!-- 上传按钮 -->
        <button @click="uploadImage" :disabled="!imageUrl">上传</button>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  // 控制弹窗显示
  const isModalOpen = ref(true); // 默认打开弹窗
  // 存储图片的 URL
  const imageUrl = ref(null);
  // 存储选中的文件
  const selectedFile = ref(null);
  
  // 关闭弹窗
  const closeModal = () => {
    isModalOpen.value = false;
    imageUrl.value = null; // 清空预览
    selectedFile.value = null; // 清空文件
  };
  
  // 处理文件选择
  const handleFileChange = (event) => {
    const file = event.target.files[0];
    if (file) {
      selectedFile.value = file;
      // 生成图片的 URL 用于预览
      imageUrl.value = URL.createObjectURL(file);
    }
  };
  
  // 上传图片（这里只是一个示例，实际需要调用 API）
  const uploadImage = () => {
    if (selectedFile.value) {
      console.log('上传的文件:', selectedFile.value);
      // 这里可以调用上传 API
      alert('头像上传成功！');
      closeModal();
    }
  };
  </script>
  
  <style scoped>
  /* 弹窗样式 */
  .modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .modal-content {
    background-color: white;
    padding: 20px;
    border-radius: 8px;
    width: 300px;
    text-align: center;
  }
  
  .close {
    float: right;
    cursor: pointer;
    font-size: 24px;
  }
  
  .image-preview img {
    max-width: 100%;
    max-height: 200px;
    margin-top: 10px;
    border-radius: 8px;
  }
  
  button {
    margin-top: 10px;
    padding: 8px 16px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }
  </style>