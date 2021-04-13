import { createRouter, createWebHistory } from 'vue-router'

const lazyLoadRoute = pageName => {
  return () => import(/* webpackChunkName: "[request]" */ `@/views/${pageName}`)
}

const routes = [
  { path: '/', name: 'home', component: lazyLoadRoute('Home') },
  { path: '/about', name: 'about', component: lazyLoadRoute('About') },
  { path: '/profile', name: 'profile', component: lazyLoadRoute('Profile') },
  { path: '/personal/:code', name: 'card', component: lazyLoadRoute('Card') },
  { path: '/admin/dashboard', name: 'dashboard', component: lazyLoadRoute('Dashboard') },
  { path: '/signin', name: 'signin', component: lazyLoadRoute('Signin') },
  { path: '/signup', name: 'signup', component: lazyLoadRoute('Signup') },
  { path: '/404', name: 'error-404', component: lazyLoadRoute('Error404') },
  { path: '/:pathMatch(.*)*', component: lazyLoadRoute('Error404') }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
