import { createRouter, createWebHistory } from 'vue-router';

// 这里将用到的路由组件文件引入
import LoginPage from '../page/LoginPage.vue';
import RegistPage from '../page/RegistPage.vue';
import HomePage from '../page/HomePage.vue';
import FindPage from '../page/FindPage.vue';
import ChatPage from '../page/ChatPage.vue';
import PersonalPage from '../page/PersonalPage.vue';
import AboutUsPage from '@/page/AboutUsPage.vue';
import NewRoomPage from '@/page/NewRoomPage.vue';
import ForgotPWPage from '@/page/ForgotPWPage.vue';

const routes = [
  {
    // 根目录
    path: '/',
    name: 'LoginPage',    // 自定义的路由名
    component: LoginPage  // 用到的组件名
  },
  {
    path: '/RegistPage',
    name: 'RegistPage',
    component: RegistPage
  },
  {
    path: '/HomePage',
    name: 'HomePage',
    component: HomePage,
    meta: { requiresAuth: true } // 需要认证
  },
  {
    path: '/FindPage',
    name: 'FindPage',
    component: FindPage,
    meta: { requiresAuth: true } // 需要认证
  },
  {
    path: '/ChatPage',
    name: 'ChatPage',
    component: ChatPage,
    meta: { requiresAuth: true } // 需要认证
  },
  {
    path: '/PersonalPage',
    name: 'PersonalPage',
    component: PersonalPage,
    meta: { requiresAuth: true } // 需要认证
  },
  {
    path: '/NewRoomPage',
    name: 'NewRoomPage',
    component: NewRoomPage,
    meta: { requiresAuth: true } // 需要认证
  },
  {
    path: '/AboutUsPage',
    name: 'AboutUsPage',
    component: AboutUsPage
  },
  {
    path: '/ForgotPWPage',
    name: 'ForgotPWPage',
    component: ForgotPWPage
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

// 全局前置守卫
router.beforeEach((to, from, next) => {
  const isAuthenticated = checkAuthentication(); // 自定义函数，检查用户是否已登录

  if (to.matched.some(record => record.meta.requiresAuth)) {
    // 如果目标路由需要认证
    if (!isAuthenticated) {
      // 如果用户未登录，重定向到登录页面
      next('/');
    } else {
      // 如果用户已登录，继续导航
      next();
    }
  } else if (to.name === 'LoginPage') {
    // 如果用户访问登录页面，清除本地存储
    localStorage.clear();
    next();
  } else {
    // 如果目标路由不需要认证，继续导航
    next();
  }
});
function checkAuthentication() {
  // 这里你可以实现具体的认证逻辑
  // 例如检查本地存储或会话存储中是否有用户信息
  return localStorage.getItem('user') !== null;
  //return 0;
}

export default router;