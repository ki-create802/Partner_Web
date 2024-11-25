import { createApp } from 'vue';
import App from './App.vue';
import ElementPlus from 'element-plus' // 导入ElementPlus
import router from './router';

const app = createApp(App);
app.use(router);
app.use(ElementPlus); // 使用Element UI
app.mount('#app');