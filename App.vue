<template>
  <div>
    <head>
      <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.1/css/bulma.css">
      <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.3.1/css/all.css" 
            integrity="sha384-mzrmE5qonljUremFsqc01SB46JvROS7bZs3IO2EmfFsd15uHvIt+Y8vEf7N7fWAU" crossorigin="anonymous">

      <!-- <script src="https://js-cdn.music.apple.com/musickit/v1/musickit.js"></script> -->

    </head>
    <body>
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

        <button class="button fas fa-backward" onclick="backward()"></button>
        <button class="button fas fa-play" onclick="play()"></button>
        <button class="button fas fa-pause" onclick="pause()"></button>
        <button class="button fas fa-forward" onclick="forward()"></button>
        <button class="button" onclick="setQueue()">Set Queue</button>
      </div>
      <!-- <script src="common.js"></script> -->
    </body>
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      music: null
    }
  },
  computed: {

    token: function () {
      return this.$http.get("https://auth.atec.pub/music")
    },

    // player
    // TODO refactor next two to toggle with backing state
    // play: function () {
    //   this.music.play()
    // },
    // pause: function () {
    //   this.music.pause()
    // },
    // forward: function () {
    //   this.music.skipToNextItem()
    // },
    // backward: function () {
    //   this.music.skipToPreviousItem()
    // },
    // setQueue: function () {
    //   this.music.setQueue({ 
    //     url: 'https://itunes.apple.com/us/album/the-organ-hearts/429020349' 
    //   }).then(function() {
    //     this.music.play()
    //   })
    // },
  },

  created: function () {
    this.onThemeChange = () => {
      this.theme = this.$localStorage.get('theme');
      document.body.className = this.theme;
    };
    this.onThemeChange();

    // Events
    this.onAlert = (alert) => {
      this.alert.details = alert;
      this.alert.countdown = 5;
    };

    let initializeMusicKit = () => {
      // Configure MusicKit
      window.MusicKit.configure({
        developerToken: this.token,
        app: {
          name: 'atec.pub',
          build: '0.1.0'
        },
      });

      initialize();
    };

    let initialize = () => {
      let musicKit = window.MusicKit.getInstance();

      if (!musicKit.isAuthorized && this.$route.meta.isLibrary) {
        this.$router.push({ path: '/', replace: true });
      }

      this.musicKit = musicKit;

      this.isAuthorized = this.musicKit.isAuthorized;
      this.musicKit.addEventListener(window.MusicKit.Events.authorizationStatusDidChange, this.onAuthorizationStatusDidChange);

      // Create callback functions
      this.musicKit.addEventListener(window.MusicKit.Events.mediaPlaybackError, this.mediaPlaybackError);

      this.mediaItemDidChange = (event) => {
        // Show a notification
        if (('Notification' in window) && event.item && this.$localStorage.get('showPlaybackNotifications')) {
          window.Notification.requestPermission();

          if (window.Notification.permission === 'granted') {
            if (this.notification) {
              this.notification.close();
            }
          }
        }
      };
      this.musicKit.addEventListener(window.MusicKit.Events.mediaItemDidChange, this.mediaItemDidChange);
    };

    // Check for MusicKit
    if (window.MusicKit) {
      try {
        this.musicKit = window.MusicKit.getInstance();
        initialize();
      } catch (err) {
        initializeMusicKit();
      }
    } else {
      document.addEventListener('musickitloaded', () => {
        initializeMusicKit();
      });
    }
  }
}
</script>