import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/myProfile',
      name: 'myProfile',
      component: () => import('../components/myProfile/MyProfile.vue'),
      meta: {
        title: 'MyProfile' // tab title on browser
      }
    },
    {
      path: '/',
      name: 'TheHome',
      component: () => import('../components/TheHome.vue'),
      meta: {
        title: 'Home' // tab title on browser
      }
    },
    {
      path: '/addProduct',
      name: 'addProduct',
      component: () => import('../components/addProduct/AddProduct.vue'),
      meta: {
        title: 'Add Product'
      }
    }
  ]
})

// Update document title on route change
router.afterEach((to) => {
  document.title = (to.meta.title as string) || 'Default Title'
})

export default router
