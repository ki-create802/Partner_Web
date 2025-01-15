<template>
  <div v-if="visible" class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <h3>确认加入房间</h3>
      </div>
      <div class="modal-body">
        <p>确认加入房间 {{ roomID }} 吗？</p>
      </div>
      <div class="modal-footer">
        <button class="modal-button confirm" @click="handleConfirm">确认</button>
        <button class="modal-button cancel" @click="handleCancel">取消</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CustomModal',
  props: {
    visible: {
      type: Boolean,
      required: true,
    },
    roomID: {
      type: String,
      required: true,
    },
  },
  methods: {
    handleConfirm() {
      this.$emit('confirm');
    },
    handleCancel() {
      this.$emit('cancel');
    },
  },
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* 半透明黑色背景 */
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000; /* 确保弹窗在最上层 */
}

.modal-content {
  border-radius: 10px;
  width: 400px;
  padding: 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  background-image: linear-gradient(rgba(255, 255, 255, 0.5), rgba(255, 255, 255, 0.5)), url('@/assets/popupwindow1.jpg'); /* 叠加半透明层 */
  background-size: cover;
  background-position: center;
}

.modal-header {
  text-align: center;
  margin-bottom: 20px;
}

.modal-header h3 {
  margin: 0;
  font-size: 1.5em;
  color: #333;
}

.modal-body {
  text-align: center;
  margin-bottom: 20px;
}

.modal-body p {
  margin: 0;
  font-size: 1.2em;
  color: #555;
}

.modal-footer {
  display: flex;
  justify-content: space-around;
}

.modal-button {
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  font-size: 1em;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.modal-button.confirm {
  background-color: #5ba4f6;
  color: white;
}

.modal-button.confirm:hover {
  background-color: #4a90e2;
}

.modal-button.cancel {
  background-color: #e6a4b4;
  color: white;
}

.modal-button.cancel:hover {
  background-color: #d48a9a;
}
</style>