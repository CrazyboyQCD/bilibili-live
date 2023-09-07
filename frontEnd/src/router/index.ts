import { createRouter, createWebHistory } from 'vue-router'
import RoomMonitor from '@/views/RoomMonitor.vue'

const routes = [
  {
    path: '/',
    name: 'RoomMonitor',
    component: RoomMonitor
  },
  {
    path: '/areainfo',
    name: 'AreaInfo',
    component: () => import('../views/AreaInfo.vue')
  },
  {
    path: '/livestatistics/:anchor',
    name: 'LiveStatistics',
    component: () => import('../views/LiveStatistics.vue')
  },
  {
    path: '/livewatch',
    name: 'LiveWatch',
    component: () => import('../views/LiveWatch.vue')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
