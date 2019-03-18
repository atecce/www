<template>
  <div>

    <Header></Header>

    <transition name="fade" mode="out-in">
      <router-view></router-view>
    </transition>

    <section class="hero">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            <router-link to="/canon">canon</router-link>
          </h1>
        </div>
      </div>
    </section>

  </div>
</template>

<style>
.fade-enter-active,
.fade-leave-active {
  transition-duration: 0.3s;
  transition-property: opacity;
  transition-timing-function: ease;
}

.fade-enter,
.fade-leave-active {
  opacity: 0
}
</style>


<script>
import Header from './Header.vue'
import Chart from './Chart.vue'

export default {
  components: {
    Header,
    Chart
  },
  created: function () {

    // TODO event tracking
    var xhr = new XMLHttpRequest()
    xhr.overrideMimeType("text/plain; charset=x-user-defined")
    xhr.open("GET", "https://auth.atec.pub/")
    xhr.send()

    // load MusicKit
    document.addEventListener('musickitloaded', function() {
      var xhr = new XMLHttpRequest()
      xhr.overrideMimeType("text/plain; charset=x-user-defined")
      xhr.open("GET", "https://auth.atec.pub/music")
      xhr.addEventListener("load", function() {
        window.MusicKit.configure({
          developerToken: this.responseText,
          app: {
            name: 'atec.pub',
            build: '0.1.0'
          }
        })
      })
      xhr.send()
    })

    // set playlist
    localStorage.setItem("focus", "pl.u-XmgMhDYx4Jml")
  }
}
</script>