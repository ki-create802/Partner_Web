import { createRouter, createWebHistory } from 'vue-router';

//这里将用到的路由组件文件引入
import LoginPage from '../page/LoginPage.vue';
import RegistPage from '../page/RegistPage.vue';
import HomePage from '../page/HomePage.vue';
import FindPage from '../page/FindPage.vue';
import ChatPage from '../page/ChatPage.vue';
import PersonalPage from '../page/PersonalPage.vue';
import AboutUsPage from '@/page/AboutUsPage.vue';

const routes = [
  {
    //根目录
    path: '/',
    name: 'LoginPage',    //自定义的路由名
    component: LoginPage  //用到的组件名
  },
  {
    path: '/RegistPage',
    name: 'RegistPage',
    component: RegistPage
  },
  {
    path: '/HomePage',
    name: 'HomePage',
    component: HomePage
  },
  {
    path: '/FindPage',
    name: 'FindPage',
    component: FindPage
  },
  {
    path: '/ChatPage',
    name: 'ChatPage',
    component: ChatPage
  },
  {
    path: '/PersonalPage',
    name: 'PersonalPage',
    component: PersonalPage
  },
  {
    path: '/AboutUsPage',
    name: 'AboutUsPage',
    component: AboutUsPage
  },

];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;