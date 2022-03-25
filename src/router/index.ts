import {createRouter, createWebHistory} from 'vue-router'
import type {RouteRecordRaw} from 'vue-router';
import HomeView from '../views/HomeView.vue'
import projectRoutes from "@/generated/projectRoutes";

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  }
]
routes.push(...projectRoutes)

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
