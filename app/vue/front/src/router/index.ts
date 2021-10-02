import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Index',
    component: () => import('../views/Index.vue')
  },
  {
    path: '/maker/list',
    name: 'MakerList',
    component: () => import('../views/Maker/List.vue')
  },
  {
    path: '/maker/detail/:path',
    name: 'MakerDetail',
    component: () => import('../views/Maker/Detail.vue')
  },
  {
    path: '/bbs',
    name: 'Bbs',
    component: () => import('../views/Bbs/Index.vue')
  },
  {
    path: '*',
    name: 'NotFound',
    component: () => import('../views/NotFound.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
