import Vue from 'vue'
import VueRouter from 'vue-router'

import App from './components/App.vue'
import Canon from './components/Canon.vue'
import IO from './components/IO.vue'

Vue.use(VueRouter)

new Vue({
  render: h => h(App),
  router: new VueRouter({
    mode: 'history',
    routes: [
      { path: '/canon', component: Canon },
      { path: '/io', component: IO }
    ]
  }),
}).$mount('#app')