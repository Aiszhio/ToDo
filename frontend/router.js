import { createRouter, createWebHistory } from 'vue-router';
import home from './Components/HomePage.vue';
import register from './Components/Reg.vue';
import Entrance from "./Components/Entrance.vue";
import cabinet from "./Components/Cabinet.vue";
import help from "./Components/Help.vue";
import notes from "./Components/Notes.vue";

const routes = [
    { path: '/', component: home },
    { path: '/register', component: register },
    { path: '/entrance', component: Entrance },
    { path: '/cabinet', component: cabinet },
    { path: '/help', component: help },
    { path: "/notes", component: notes }
];

const router = createRouter({
    history: createWebHistory(), // Используйте режим history для чистых URL
    routes,
});

export default router;
