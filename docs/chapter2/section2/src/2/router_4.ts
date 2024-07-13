import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './pages/HomePage.vue'
import NotFound from './pages/NotFound.vue'
import PingPage from './pages/PingPage.vue'
import LoginPage from './pages/LoginPage.vue'
import CityPage from './pages/CityPage.vue'

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/ping', name: 'ping', component: PingPage },
  { path: '/login', name: 'login', component: LoginPage },
  { path: '/city/:cityName', name: 'city', component: CityPage, props: true },
  { path: '/:path(.*)*', component: NotFound }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

/*[!code ++]*/ router.beforeEach(async (to) => {
  /*[!code ++]*/ if (to.path === '/login') {
    return true //[!code ++]
  } //[!code ++]
  const res = await fetch('/api/me') //[!code ++]
  if (res.ok) return true //[!code ++]
  return '/login' //[!code ++]
}) //[!code ++]

export default router
