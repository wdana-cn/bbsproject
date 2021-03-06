import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import VueParticles from 'vue-particles'
Vue.use(VueParticles)

import ElementUI from 'element-ui'
import { Notification } from 'element-ui'
import 'assets/element-variables.scss'

Vue.use(ElementUI)

// require('@/mock/mock.js')

Vue.config.productionTip = false
Vue.config.silent = true
Vue.prototype.$notify = Notification;


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
