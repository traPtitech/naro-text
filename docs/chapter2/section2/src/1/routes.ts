export const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/:path(.*)*', component: NotFound }
]
