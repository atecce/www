import Vue from 'vue'
import VueRouter from 'vue-router'

import App from './components/App.vue'
import Canon from './components/Canon.vue'

Vue.use(VueRouter)

new Vue({
  render: h => h(App),
  router: new VueRouter({
    routes: [
      { path: '/canon', component: Canon }
    ]
  }),
}).$mount('#app')