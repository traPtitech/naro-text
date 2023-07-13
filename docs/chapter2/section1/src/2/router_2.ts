import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './pages/HomePage.vue'
import NotFound from './pages/NotFound.vue'
import AxiosPage from './pages/AxiosPage.vue'
import LoginPage from './pages/LoginPage.vue' //[!code ++]

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/axios', name: 'axios', component: AxiosPage },
  { path: '/login', name: 'login', component: LoginPage }, //[!code ++]
  { path: '/:path(.*)', component: NotFound }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
