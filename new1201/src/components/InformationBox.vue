<template>
  <div class="background-layer"></div>
  <div class="information">
    <div class="avatar-container">
      <div class="avatar">
        <img :src="avatarUrl" alt="头像" id="avatar-img">
      </div>
      <input type="file" class="avatar-upload" accept="image/*" @change="onAvatarChange" ref="avatarUpload" />
      <button class="avatar-edit-btn" @click="triggerAvatarUpload">上传头像</button>
    </div>
    <div class="user-info">
      <div class="username">
        <span id="username-text">{{ username }}</span>
        <button @click="editusername" class="edit-btn">编辑</button>
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
import { getFansNum,editNewName,uploadAvatar } from '@/api.js';

export default {
  name: 'InformationBox',
  data() {
    const user = JSON.parse(localStorage.getItem('user'));
    //alert("本地存储user："+localStorage.getItem('user'));
    const UserID = user.UID;
    alert("本地存储："+JSON.stringify(user));
    return{
      UserID,
      avatarUrl: UserID ? `http://localhost:8082/avatars/${UserID}.jpg` : require("@/assets/avatar.png"),
      username: JSON.parse(localStorage.getItem('user')).UName||'默认用户名',
      followersCount: 2000,
      signature: JSON.parse(localStorage.getItem('user')).URemark ||'默认个性签名',
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
      // 调用上传头像函数
      uploadAvatar(this.UserID, file)
        .then((success) => {
          if (success) {
            alert("头像上传成功！");
          } else {
            alert("头像上传失败，请重试。");
          }
        })
        .catch((error) => {
          console.error("上传失败：", error);
        });
    },
    async editusername() {
      const newusername = prompt('请输入新的昵称', this.username);
      if(!newusername||newusername.trim=="")return;
      //向后端发送修改昵称请求
      try{
        let ok=false;
        ok=await editNewName(this.UserID,newusername);
        if(!ok){
          alert("修改昵称失败！");
          return;
        }
      }catch{
        alert("修改昵称失败！");
        return;
      }
      let user=JSON.parse(localStorage.getItem('user'));
      user.UName=newusername;
      localStorage.setItem('user', JSON.stringify(user));
      this.username=user.UName;
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
    async getFansNum_(){
      try{
        const fanNum=await getFansNum();
        this.followersCount=fanNum;
      }catch{
        alert("请求粉丝数失败");
      }
    }
  },
  mounted() {
    this.getFansNum_();
  }
}
</script>

<style>
.information {
  display: flex;
  padding: 25px;
  border: 1px solid #ddd;
  border-radius: 8px;
  width: 100%;
  background-color: rgba(255, 255, 255, 0.1);  /* 设置白色半透明背景 */
}

.avatar-container {
  position: relative;
  margin-bottom: 5px;
  margin-left: 25px;
  margin-right: 35px;
  background-color: rgba(255, 255, 255, 0.1);  /* 设置白色半透明背景 */

}

.avatar {
  width: 95px;
  height: 95px;
  border-radius: 50%;
  /* background: #3578c9; */
  overflow: hidden;
  margin-bottom: 10px;
  margin-top: 15px;
  background-color: rgba(255, 255, 255, 0.1);  /* 设置白色半透明背景 */

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
  font-size: 16px;
}

.user-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 30px;
  margin-left: 10px;
  margin-right: 20px;
}

.username {
  font-size: 23px;
  margin-bottom: 25px;
}

.followers {
  font-size: 15px;
}

.username {
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
  height: 70px;
  padding: 10px;
  margin-top: 50px;
  border-radius: 8px;
  border: 1px solid transparent;  /* 初始边框透明 */
  resize: none;
  font-size: 17px;
  background: none;
  transition: all 0.3s ease;  /* 添加过渡效果 */
}

.signature:focus {
  border-color: white;  /* 聚焦时显示边框 */
  background-color: rgba(255, 255, 255, 0.3);  /* 聚焦时背景更不透明 */
  outline: none;  /* 去掉默认的焦点轮廓 */
}

.signature:hover {
  border-color: white;  /* 鼠标悬停时显示边框 */
  background-color: rgba(255, 255, 255, 0.3);  /* 鼠标悬停时背景更不透明 */
}
</style>