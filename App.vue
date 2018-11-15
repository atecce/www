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
    </div>
  </div>
</template>

<script>
export default {
  created: function () {
    var xhr = new XMLHttpRequest()
    xhr.overrideMimeType("text/plain; charset=x-user-defined")
    xhr.open("GET", "https://auth.atec.pub/")
    xhr.send()

    document.addEventListener('musickitloaded', function() {

      var xhr = new XMLHttpRequest()
      xhr.overrideMimeType("text/plain; charset=x-user-defined")
      xhr.open("GET", "https://auth.atec.pub/music")
      xhr.addEventListener("load", function() {

        MusicKit.configure({
          developerToken: this.responseText,
          app: {
            name: 'atec.pub',
            build: '0.1.0'
          }
        })

        let music = MusicKit.getInstance()

        music.play()

      })
      xhr.send()

    })
  }
}
</script>