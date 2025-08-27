import { createRouter, createWebHistory } from 'vue-router'
import IndexView from '../components/IndexView.vue'
import NewIncidentView from '../components/NewIncidentView.vue'

const routes = [
  { path: '/', component: IndexView },
  { path: '/new', component: NewIncidentView }
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
