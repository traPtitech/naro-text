import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './pages/HomePage.vue'
import NotFound from './pages/NotFound.vue'
import PingPage from './pages/PingPage.vue' //[!code ++]

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/ping', name: 'ping', component: PingPage }, //[!code ++]
  { path: '/:path(.*)*', component: NotFound }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
