<template>
  <div>
    <div class="container">
      <nav class="navbar" role="navigation" aria-label="main navigation">
        <div class="navbar-brand">
          <a class = "navbar-item" href="https://keybase.io/atec" target="_blank">
            <i class="fab fa-keybase"></i>
          </a>
          <a class = "navbar-item" href="https://github.com/atecce" target="_blank">
            <i class="fab fa-github"></i>
          </a>
          <a class = "navbar-item" href="https://linkedin.com/in/atecce" target="_blank">
            <i class="fab fa-linkedin"></i>
          </a>
          <a class="navbar-item" target="_blank" href="/resume.html">
            <i class="fas fa-file-alt"></i>
          </a>
        </div>
      </nav>

      <h2 class="subtitle">Dogmatic opinions. Pragmatic decisions</h2>
      <h2 class="subtitle">Stay curious</h2>

      <button class="button fas fa-backward" v-on:click="backward()"></button>
      <button class="button fas fa-play" v-on:click="play()"></button>
      <button class="button fas fa-pause" v-on:click="pause()"></button>
      <button class="button fas fa-forward" v-on:click="forward()"></button>
      <button class="button" v-on:click="setQueue()">Set Queue</button>
    </div>
    <!-- <script src="common.js"></script> -->
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      music: null,
      token: null
    }
  },

  methods: {

    // player
    // TODO refactor next two to toggle with backing state
    play: function () {
      window.MusicKit.getInstance().play()
    },
    pause: function () {
      this.music.pause()
    },
    forward: function () {
      this.music.skipToNextItem()
    },
    backward: function () {
      this.music.skipToPreviousItem()
    },
    setQueue: function () {
      this.music.setQueue({
        album: '429020349'
      })
    },
  },

  created: function () {

    document.addEventListener('musickitloaded', function () {

      var xhr = new XMLHttpRequest()
      xhr.overrideMimeType("text/plain; charset=x-user-defined")
      xhr.open("GET", "https://auth.atec.pub/music", false)
      xhr.send()
      this.token = xhr.responseText

      // eslint-disable-next-line
      console.log("configuring MusicKit")

      // eslint-disable-next-line
      console.log(this.token)

      window.MusicKit.configure({
        developerToken: this.token,
        app: {
          name: 'atec.pub',
          build: '0.1.0'
        },
      })

      this.music = window.MusicKit.getInstance()

      // eslint-disable-next-line
      console.log(this.music)
    })
  }
}
</script>