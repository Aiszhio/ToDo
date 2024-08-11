import { createRouter, createWebHistory } from 'vue-router';
import home from './Components/HomePage.vue';
import register from '@/Components/Reg.vue';
import Entrance from "@/Components/Entrance.vue";

const routes = [
    { path: '/', component: home },
    { path: '/register', component: register },
    { path: '/entrance', component: Entrance }
];

const router = createRouter({
    history: createWebHistory(), // Используйте режим history для чистых URL
    routes,
});

export default router;
