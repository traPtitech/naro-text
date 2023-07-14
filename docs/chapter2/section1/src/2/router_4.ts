import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './pages/HomePage.vue'
import NotFound from './pages/NotFound.vue'
import AxiosPage from './pages/AxiosPage.vue'
import LoginPage from './pages/LoginPage.vue'
import CityPage from './pages/CityPage.vue'
import axios from 'axios' //[!code ++]

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/axios', name: 'axios', component: AxiosPage },
  { path: '/login', name: 'login', component: LoginPage },
  { path: '/city/:cityName', name: 'city', component: CityPage, props: true },
  { path: '/:path(.*)', component: NotFound }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

/*[!code ++]*/ router.beforeEach(async (to) => {
  /*[!code ++]*/ try {
    await axios.get('/api/whoami') //[!code ++]
  } /*[!code ++]*/ catch (_) {
    /*[!code ++]*/ if (to.path === '/login') {
      return true //[!code ++]
    } //[!code ++]
    return '/login' //[!code ++]
  } //[!code ++]
  return true //[!code ++]
}) //[!code ++]

export default router
