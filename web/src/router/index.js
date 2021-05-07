import { createRouter, createWebHistory } from 'vue-router'

const lazyLoadRoute = pageName => {
  return () => import(/* webpackChunkName: "[request]" */ `@/views/${pageName}`)
}

const routes = [
  { path: '/', name: 'home', component: lazyLoadRoute('Home'), meta: { title: 'Dashboard' } },
  { path: '/c/:code', name: 'card', component: lazyLoadRoute('Card'), meta: { title: 'Card' } },
  { path: '/setting', name: 'setting', component: lazyLoadRoute('Setting'), meta: { title: 'Setting' } },
  { path: '/admin', name: 'admin', component: lazyLoadRoute('Admin'), meta: { title: 'Admin' } },
  { path: '/company', name: 'company', component: lazyLoadRoute('Company'), meta: { title: 'Company' } },
  { path: '/404', name: 'error-404', component: lazyLoadRoute('Error404'), meta: { title: '404' } },
  { path: '/:pathMatch(.*)*', component: lazyLoadRoute('Error404'), meta: { title: '404' } }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title

  const currentUser = localStorage.getItem('user') && JSON.parse(localStorage.getItem('user'))
  const isAuthenticated = localStorage.getItem('is_authenticated')

  if (to.name === 'admin' && (!isAuthenticated || currentUser.role !== 4)) next({ name: 'home' })

  if (to.name === 'company' && !(isAuthenticated && currentUser.role === 3 && currentUser.type === 2)) next({ name: 'home' })

  if (to.name === 'setting' && !isAuthenticated) next({ name: 'home' })

  next()
})

export default router
