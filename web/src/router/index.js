import { createRouter, createWebHistory } from 'vue-router'

const lazyLoadRoute = pageName => {
  return () => import(/* webpackChunkName: "[request]" */ `@/views/${pageName}`)
}

const routes = [
  { path: '/', name: 'home', component: lazyLoadRoute('Home') },
  { path: '/about', name: 'about', component: lazyLoadRoute('About') },
  { path: '/profile', name: 'profile', component: lazyLoadRoute('Profile') },
  { path: '/c/:code', name: 'card', component: lazyLoadRoute('Card') },
  { path: '/setting', name: 'setting', component: lazyLoadRoute('Setting') },
  { path: '/admin', name: 'admin', component: lazyLoadRoute('Admin'), isPrivate: true },
  { path: '/company', name: 'company', component: lazyLoadRoute('Company'), isPrivate: true },
  { path: '/signin', name: 'signin', component: lazyLoadRoute('Signin'), isPrivate: false },
  { path: '/signup', name: 'signup', component: lazyLoadRoute('Signup') },
  { path: '/404', name: 'error-404', component: lazyLoadRoute('Error404') },
  { path: '/:pathMatch(.*)*', component: lazyLoadRoute('Error404') }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const currentUser = localStorage.getItem('user') && JSON.parse(localStorage.getItem('user'))
  const isAuthenticated = localStorage.getItem('is_authenticated')

  if (to.name === 'admin' && !isAuthenticated) next({ name: 'home' })

  if (to.name === 'company' && !(currentUser.role === 'Company Manager' && currentUser.type === 'Buniness')) next({ name: 'home' })

  if (to.name === 'setting' && !isAuthenticated) next({ name: 'home' })

  next()
})

export default router
