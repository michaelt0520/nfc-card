import { createRouter, createWebHistory } from 'vue-router'

const lazyLoadRoute = pageName => {
  return () => import(/* webpackChunkName: "[request]" */ `@/views/${pageName}`)
}

const routes = [
  { path: '/', name: 'home', component: lazyLoadRoute('Home') },
  { path: '/about', name: 'about', component: lazyLoadRoute('About') },
  { path: '/profile', name: 'profile', component: lazyLoadRoute('Profile') },
  { path: '/personal/:code', name: 'personal', component: lazyLoadRoute('Personal') },
  { path: '/personal/:code/edit', name: 'edit', component: lazyLoadRoute('Edit') },
  { path: '/admin/dashboard', name: 'dashboard', component: lazyLoadRoute('Dashboard'), isPrivate: true },
  { path: '/signin', name: 'signin', component: lazyLoadRoute('Signin'), isPrivate: false },
  { path: '/signup', name: 'signup', component: lazyLoadRoute('Signup') },
  { path: '/404', name: 'error-404', component: lazyLoadRoute('Error404') },
  { path: '/:pathMatch(.*)*', component: lazyLoadRoute('Error404') }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

const isAuthenticated = localStorage.getItem('is_authenticated')

router.beforeEach((to, from, next) => {
  if (to.name === 'dashboard' && !isAuthenticated) next({ name: 'signin' })

  if (to.name === 'signin' && isAuthenticated) next({ name: 'home' })

  next()
})

export default router
