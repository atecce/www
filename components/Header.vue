<template>
  <div class="container">
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <router-link class="navbar-item icon-flipped" to="/">
          <i class="fas fa-slash"></i>
        </router-link>
        <a class="navbar-item" href="https://keybase.io/atec" target="_blank">
          <i class="fab fa-keybase"></i>
        </a>
        <a class="navbar-item" href="https://itunes.apple.com/profile/atec" target="_blank">
          <i class="fab fa-apple"></i>
        </a>
        <a class = "navbar-item" href="https://stackoverflow.com/story/atec" target="_blank">
          <i class="fab fa-stack-overflow"></i>
        </a>
        <a class = "navbar-item" href="https://github.com/atecce" target="_blank">
          <i class="fab fa-github"></i>
        </a>
        <a class="navbar-item" href="https://linkedin.com/in/atecce" target="_blank">
          <i class="fab fa-linkedin"></i>
        </a>
        <a class="navbar-item" target="_blank" href="/resume.pdf">
          <i class="fas fa-file-alt"></i>
        </a>
        <a class="navbar-item" href="mailto:root@atec.pub">
          <i class="fas fa-envelope"></i>
        </a>
      </div>
    </nav>
  </div>
</template>

<style>
.icon-flipped {
  transform: scaleX(-.75);
  -moz-transform: scaleX(-.75);
  -webkit-transform: scaleX(-.75);
  -ms-transform: scaleX(-.75);
}
</style>
