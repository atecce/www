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
    <div class="chart"></div>
  </div>
</template>

<script>
import * as d3 from 'd3';

export default {
  methods: {

    temp() {


      d3.json("http://localhost:8080/Maslin,%20T.%20Paul/Occurrence%20of%20the%20Garter%20Snake,%20Thamnophis%20sirtalis,%20in%20the%20Great%20Plains%20and%20Rocky%20Mountains.json").then(function(entities) {

        var x = []
        var y = []

        for (var entity in entities) {
          if (entities.hasOwnProperty(entity)) {
              x.push(entity)
              y.push(entities[entity])
          }
        }

        d3.select(".chart")
          .selectAll("div")
            .data(y)
          .enter().append("div")
            .style("width", function(d) { return d * 10 + "px" })
            .text(function(d) { return d })


        // eslint-disable-next-line
        console.log(x)

        // eslint-disable-next-line
        console.log(y)
      })
    }
  },

  mounted: function () {

    this.temp()
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
  }
}
</script>

<style>
.chart div {
  font: 10px sans-serif;
  background-color: steelblue;
  text-align: right;
  padding: 3px;
  margin: 1px;
  color: white;
}

svg {
  margin: 25px;
  fill: none;
  stroke: #76BF8A;
  stroke-width: 3px;
}
</style>
