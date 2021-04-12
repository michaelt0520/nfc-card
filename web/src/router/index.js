import { createRouter, createWebHistory } from 'vue-router'

const lazyLoadRoute = pageName => {
  return () => import(/* webpackChunkName: "[request]" */ `@/views/${pageName}`)
}

const routes = [
  { path: '/', name: 'Home', component: lazyLoadRoute('Home') },
  { path: '/about', name: 'About', component: lazyLoadRoute('About') },
  { path: '/profile', name: 'Profile', component: lazyLoadRoute('Profile') },
  { path: '/personal/:code', name: 'Card', component: lazyLoadRoute('Card') },
  { path: '/admin/dashboard', name: 'Dashboard', component: lazyLoadRoute('Dashboard') },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
