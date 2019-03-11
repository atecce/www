<template>
  <div>
    <Header></Header>
    <Chart></Chart>
  </div>
</template>

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