<template>
  <div class="information">
    <div class="avatar-container">
      <div class="avatar">
        <img :src="avatarUrl" alt="头像" id="avatar-img">
      </div>
      <input type="file" class="avatar-upload" accept="image/*" @change="onAvatarChange" ref="avatarUpload" />
      <button class="avatar-edit-btn" @click="triggerAvatarUpload">上传头像</button>
    </div>
    <div class="user-info">
      <div class="nickname">
        <span id="nickname-text">{{ nickname }}</span>
        <button @click="editNickname" class="edit-btn">编辑</button>
      </div>
      <div class="followers">
        <span id="followers-text">粉丝数: {{ followersCount }}</span>
      </div>
    </div>
    <div class="information-right">
      <textarea 
        class="signature" 
        placeholder="编辑个性签名..." 
        v-model="signature"
        @blur="saveSignature"> <!--在失去焦点时保存签名-->
      </textarea>  
      
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'InformationBox',
  data() {
    return{
      avatarUrl: require("@/assets/avatar.png"), // 默认头像
      nickname: '张三',
      followersCount: 2000,
      signature: '这是我的个性签名',
    };
  },
  methods: {
    triggerAvatarUpload() {
      this.$refs.avatarUpload.click(); // 触发文件选择框
    },
    onAvatarChange(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.avatarUrl = e.target.result; // 更新头像
        };
        reader.readAsDataURL(file);
      }
    },
    editNickname() {
      const newNickname = prompt('请输入新的昵称', this.nickname);
      if (newNickname) {
        this.nickname = newNickname; // 更新昵称
        localStorage.setItem('userNickname', this.nickname); // 保存到 localStorage

        // 假设你的 API 路径是 /update-nickname
        axios.post('/update-nickname', { nickname: this.nickname })
        .then(response => {
          console.log('昵称更新成功:', response.data);
        })
        .catch(error => {
          console.error('更新昵称失败:', error);
        });
      }
    },
    saveSignature() {
      localStorage.setItem('userSignature', this.signature); // 保存到 localStorage

      // 假设你的 API 路径是 /update-signature
      axios.post('/update-signature', { signature: this.signature })
      .then(response => {
        console.log('签名更新成功:', response.data);
      })
      .catch(error => {
        console.error('保存签名失败:', error);
      });
    },
    //初始化时从 localStorage 获取数据
    // mounted() {
    //   const savedNickname = localStorage.getItem('userNickname');
    //   if (savedNickname) {
    //     this.nickname = savedNickname;
    //   }
    //   const savedSignature = localStorage.getItem('userSignature');
    //   if (savedSignature) {
    //     this.signature = savedSignature;
    //   }
    //   const savedAvatar = localStorage.getItem('userAvatar');
    //   if (savedAvatar) {
    //     this.avatarUrl = savedAvatar;
    //   }
    // }
  }
}
</script>

<style>
.information {
  display: flex;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 8px;
  width: 100%;
  background-color: #9bbec27e;
}

.avatar-container {
  position: relative;
  margin-bottom: 5px;
  margin-left: 10px;
  margin-right: 20px;
}

.avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  /* background: #3578c9; */
  overflow: hidden;
  margin-bottom: 10px;
  margin-top: 15px;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-upload {
  display: none;
}

.avatar-edit-btn {
  margin-top: 10px;
  padding: 5px 10px;
  background: none;
  border: none;
  color: #3578c9;
  border-radius: 4px;
  cursor: pointer;
}

.user-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 40px;
  margin-left: 10px;
  margin-right: 20px;
}

.nickname {
  font-size: 18px;
  margin-bottom: 15px;
}

.followers {
  font-size: 12px;
}

.nickname {
  font-weight: bold;
}

.edit-btn {
  margin-left: 10px;
  font-size: 14px;
  color: #4e90f0;
  background: none;
  border: none;
  cursor: pointer;
}

.edit-btn:hover {
  text-decoration: underline;
}

.information-right {
  flex: 2;
  padding-left: 20px;
}

.signature {
  width: 90%;
  height: 65px;
  padding: 10px;
  margin-top: 50px;
  border-radius: 8px;
  border: 1px solid #494848;
  resize: none;
  font-size: 14px;
  background: none;
}

.signature:focus {
  border-color: #4e90f0;
  outline: none;
}
</style>