import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Login",
      component: () => import("@/views/Login.vue")
    },
    {
      path: "/layout",
      name: "Layout",
      component: () => import("@/views/Layout.vue"),
      children: [
        {
          path: "/home",
          name: "Home",
          component: () => import("@/views/Home.vue")
        },
        {
          path: "/profile",
          name: "Profile",
          component: () => import("@/views/Profile.vue")
        },
      ]
    },
  ]
})

export default router
