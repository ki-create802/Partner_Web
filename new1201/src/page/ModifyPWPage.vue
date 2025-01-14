<template>
    <GuideBar/>
    <div class="bgimg"></div>
    <div class = "modify-pw-container">
        <h1>修改密码</h1>
        <form @submit.prevent="resetPassword">
          <div class="form-group">
            <label for="newPassword">新的密码</label>
            <input
              type="password"
              id="newPassword"
              v-model="newPassword"
              placeholder="请输入密码"
              required
            />
          </div>
          <div class="form-group">
            <label for="confirmPassword">请确认密码</label>
            <input
              type="password"
              id="confirmPassword"
              v-model="confirmPassword"
              placeholder="请再次确认密码"
              required
            />
          </div>
          <button type="submit" @click="editPW_">Reset Password</button>
        </form>
    </div>
</template>

<script>
import {editPW} from '@/api'
import GuideBar from '../components/GuideBar.vue';
export default{
    name: "ModifyPWPage",
    components:{
        GuideBar,
    },
    data(){
        return{
            newPassword:"",
            confirmPassword:"",
        }
    },
    methods:{
      async editPW_(){
        if(this.newPassword!=this.confirmPassword){
          alert("请确认两次输入密码一致！");
          return;
        }
        let UserID = 0;
        try {
            const storedUser = localStorage.getItem('user');
            if (storedUser) {
                const user = JSON.parse(storedUser);
                UserID = user.UID;
            }
        } catch {
            alert("获取本地个人信息失败");
        }
        try{
          let ok=false;
          ok=await editPW(UserID,this.newPassword);
          if(ok)alert("修改密码成功！");
          else alert("修改密码失败！");
        }catch{
          alert("修改密码失败！");
        }
      }
    }
}
</script>

<style scoped>
.bgimg{
        /* width: 150px;
        height: 60px; */
        width: 100%;
        height: 100%;
        position: fixed; 
        top: 0;
        left: 0;
        background-image: url('~@/assets/login_bg.JPG');
        background-size: cover;
        background-position: center; /* 居中对齐背景图 */
        z-index: -1; /* 设置层级，使其在所有内容下方 */
    }
.modify-pw-container{
    max-width: 400px;
    margin: 180px auto;
    padding: 40px;
    border: 1px solid #ddd;
    border-radius: 8px;
    /* background-color: #f9f9f9; */
    background-color: rgba(255, 255, 255, 0.5);
    text-align: center;
}
h1 {
  margin-bottom: 20px;
}
button {
  padding: 8px 12px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
}
.form-group {
    margin-bottom: 15px;
    text-align: left;
}

label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
}

input {
    width: 100%;
    padding: 8px;
    margin-bottom: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 14px;
}

</style>