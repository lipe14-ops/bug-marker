import { createMemoryHistory, createRouter } from 'vue-router'
import Mark from './pages/Mark.vue'
import Projects from './pages/Projects.vue'
import Login from './pages/Login.vue'
import CreateUser from './pages/CreateUser.vue'

const routes = [
  { path: '/', component: Login },
  { name: 'login', path: '/login', component: Login },
  { name: 'projects', path: '/projects/:ownerID', component: Projects, props: true },
  { name: 'project', path: '/project/:projectID', component: Mark, props: true },
  { name: 'signin', path: '/signin', component: CreateUser },
]

export const router = createRouter({
  history: createMemoryHistory(),
  routes,
})