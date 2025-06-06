import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(""),
  routes: [
    {
      path: '/',
      name: 'home',
      component:()=> import("@/view/home/index.vue")
    },
    {
      path: '/login',
      name: 'login',
      component:()=> import("@/view/user/login.vue")
    },
    {
      path: '/register',
      name: 'register',
      component:()=> import("@/view/user/register.vue")
    }
  ]
})

export default router
