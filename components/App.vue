<template>
  <div>

    <Header></Header>

    <transition name="fade" mode="out-in">

      <router-view></router-view>

      <section class="hero">
        <div class="hero-body">
          <div class="container">
            <p>
              Hai! My name's Alex. I write software to make a living, and you can find more about me profesionally <a href="/resume.pdf">here</a>.
              <br><br>
              But really it's the closest thing I know there is to having super powers, and I'd be doing it anyway.
              I know of no other skill where you can write your careful thought down and watch it solve problems, and that forms the core of my interest in it.
              <br><br>
              In particular I'm fascinated how text serves as a universal interface to make meaningful things, 
              and that perspective spans from the way the culture of Unix creates practical tools for engineering systems that are almost ubiquitously adopted 
              (and indeed, have become invisible to most people), to academic research in natural language processing 
              to explain how exactly it is that I'm communicating with you through this sentence right now.
              If you're into that sort of thing as well, you can see some personal stuff I've been messing around with <router-link to="/canon">here</router-link>.
            </p>
          </div>
        </div>
      </section>

    </transition>
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