import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/myProfile',
      name: 'myProfile',
      component: () => import('../components/myProfile/MyProfile.vue')
    },
    {
      path: '/',
      name: 'TheHome',
      component: () => import('../components/TheHome.vue')
    }
  ]
})

export default router
