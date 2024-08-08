import { createRouter, createWebHistory } from 'vue-router';
import home from './Components/HomePage.vue';
import register from './Components/Reg.vue';

const routes = [
    { path: '/', component: home },
    { path: '/register', component: register },
];

const router = createRouter({
    history: createWebHistory(), // Используйте режим history для чистых URL
    routes,
});

export default router;
