import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import { createPinia } from 'pinia';

//导入Naive UI组件库
import naive from "naive-ui";

const app = createApp(App);

app.use(router);
app.use(createPinia());
app.use(naive);

app.mount('#app');
