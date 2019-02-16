<template>
  <div>
    <div class="container">
      <nav class="navbar" role="navigation" aria-label="main navigation">
        <div class="navbar-brand">
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
        </div>
      </nav>
    </div>
    <svg width="500" height="270">
      <g style="transform: translate(0, 10px)">
        <path :d="line" />
      </g>
    </svg>
  </div>
</template>

<script>
import * as d3 from 'd3';

export default {
  data: function () {
    return {
      data: [99, 71, 78, 25, 36, 92],
      line: '',
    }
  },
  methods: {
    getScales() {

      // set ranges of x and y axis
      const x = d3.scaleLinear().range([0, 430]);
      const y = d3.scaleLinear().range([210, 0]);

      // scale left and top axis
      d3.axisLeft().scale(x);
      d3.axisTop().scale(y);

      // ???
      x.domain(d3.extent(this.data, (d, i) => i));
      y.domain([0, d3.max(this.data, d => d)]);

      return { x, y };
    },

    calculatePath() {
      const scale = this.getScales();

      // ???
      const path = d3.line()
        .x((d, i) => scale.x(i))
        .y(d => scale.y(d));

      this.line = path(this.data);
    },

    temp() {
      var xhr = new XMLHttpRequest()
      xhr.open("GET", "http://localhost:8080/Maslin,%20T.%20Paul/Occurrence%20of%20the%20Garter%20Snake,%20Thamnophis%20sirtalis,%20in%20the%20Great%20Plains%20and%20Rocky%20Mountains.json")
      xhr.send()
      xhr.addEventListener("load", function() {
        // eslint-disable-next-line
        console.log(this.responseText)
      })
    }
  },

  mounted: function () {
    this.calculatePath();
  },
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

    localStorage.setItem("focus", "pl.u-XmgMhDYx4Jml")

    this.temp()
  }
}
</script>

<style>
svg {
  margin: 25px;
  fill: none;
  stroke: #76BF8A;
  stroke-width: 3px;
}
</style>
