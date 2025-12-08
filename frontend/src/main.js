import './style.css'

import App from './App.vue'
import { createApp } from 'vue'
import VueKonva from 'vue-konva';
import { createPinia } from 'pinia'


import { router } from './router.js'

const pinia = createPinia()

createApp(App).use(pinia).use(router).use(VueKonva).mount('#app')
