import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

import App from './components/App.vue'
import Canon from './components/Canon.vue'

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: App },
    { path: '/canon', component: Canon }
  ]
})

new Vue({
  el: '#app',
  router,
  template: App,
  components: { Canon }
})