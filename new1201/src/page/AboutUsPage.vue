<template>
    <GuideBar />
    <div class="welcome-wrap">
        <div class="welcome-text">
            <p>欢迎来到约吧！</p>
            <p>连接与共享，让生活更便捷、更有趣</p>
        </div>
    </div>
    <div class="introduction-text">
        <div class="text1">
            <h1>关于我们</h1>
            <p>我们的平台致力于通过连接与共享，为用户提供便捷、高效、有趣的生活服务。</p>
            <p>我们希望通过技术手段，打破人与人之间的隔阂，让用户能够轻松找到志同道合的伙伴，共同享受生活的乐趣。</p>
            <p>无论是约车出行、约球运动，还是寻找各种搭子服务，我们都希望能够为用户提供一个安全、可靠、便捷的平台。</p>
            <p>我们的目标是连接您与同样兴趣的人，让您的每一次出行、每一场运动都更加有趣和高效。</p>
            <p>我们相信，通过我们的努力，用户的生活将变得更加丰富多彩。</p>
        </div>
    </div>
    <div class="img-connection">
        <img src="@/assets/connection.jpg" style="width: 90%; height: 500px;">
    </div>
    <div class="service-wrap">
        <h1>我们提供的服务</h1>
        <div class="service-image">
            <div v-for="(item, index) in items" :key="index" class="image-item">
                <img :src="item.img" alt="图片" style="height: 500px;" />
                <p v-html="item.text"></p>
            </div>
        </div>
    </div>
    <div class="aboutus-wrap">
        <h1>我们的团队</h1>
        <p>我们组建了一支由学院技术巨匠组成的团队。在这里，我们的目标是打造创新的数字生活体验，不断颠覆传统。</p>
    </div>
</template>

<script>
import GuideBar from '@/components/GuideBar.vue';
export default{
    name:'AboutUsPage',
    components:{
      GuideBar,
    },
    data(){
      return {
        items: [
          { 
            img: require('@/assets/drive.jpg'), 
            text: '您可以发布或搜索拼车需求，系统会根据出发地、目的地、时间<br>等信息智能匹配合适的拼车伙伴，这种方便快捷的拼车服务<br>能帮助您轻松找到同行伙伴，共享出行成本，减少交通压力。' 
          },
          { img: require('@/assets/football.jpg'), 
            text: '您可以发布约球活动，系统会根据您的运动偏好和地理位置智能<br>匹配合适的球友，无论是篮球、足球还是其他体育活动，<br>都能快速找到志同道合的队友，约球更方便。' 
          },
          { img: require('@/assets/interests.jpg'), 
            text: '找搭子服务旨在为您提供多样化的社交和活动组织平台，无论是<br>一起骑车、登山，还是摄影等其他日常活动，您都可以根据<br>自己的爱好和需求，找到志同道合的伙伴，提升生活质量。' 
          }
        ]
      };
    }
}
</script>

<style scoped>
.welcome-wrap {
  /* 设置容器为全屏，使用背景图片 */
  width: 100vw;
  height: 600px;
  background-image: url('@/assets/partner.jpg'); 
  background-size: cover; /*背景图片覆盖整个容器 */
  background-position: center; /* 背景图片居中 */
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
}
.welcome-text {
  font-size: 3rem;
  color: white;
  font-weight: bold;
  text-shadow: 3px 3px 6px rgba(0, 0, 0, 0.5); /* 文字阴影 */
}
.introduction-text {
  margin-top: 100px;
  margin-bottom: 50px;
}
.introduction-text .text1 p {
  font-size: 20px;
}
.service-wrap {
  margin-top: 100px;
}
.service-wrap h1 {
  margin-bottom: 30px;
}
.service-image {
  display: flex;
  justify-content: space-evenly;
}
.image-item p {
  font-size: 18px; 
  color: #333;  
  text-align: center; 
}
.aboutus-wrap {
  margin-top: 80px;
}
.aboutus-wrap p {
  font-size: 20px;
}
</style>