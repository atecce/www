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

    <svg class="chart"></svg>
  </div>
</template>

<script>
import * as d3 from 'd3';

export default {
  methods: {

    temp() {

      d3.json("http://canon.atec.pub/Maslin,%20T.%20Paul/Occurrence%20of%20the%20Garter%20Snake,%20Thamnophis%20sirtalis,%20in%20the%20Great%20Plains%20and%20Rocky%20Mountains.json").then(function(entities) {

        var namedEntities = []
        var counts = []

        for (var entity in entities) {
          if (entities.hasOwnProperty(entity)) {
              namedEntities.push(entity)
              counts.push(entities[entity])
          }
        }

        var width = 420, barHeight = 20

        var x = d3.scaleLinear()
          .domain([0, d3.max(counts)])
          .range([0, width])

        var chart = d3.select(".chart")
          .attr("width", width)
          .attr("height", barHeight * namedEntities.length)

        var bar = chart.selectAll("g")
          .data(counts)
          .enter().append("g")
          .attr("transform", function(d, i) { return "translate(0," + i * barHeight + ")" })

        bar.append("rect")
          .attr("width", x) 
          .attr("height", barHeight - 1)

        bar.append("text")
          .attr("x", function(d) { return x(d) - 3 })
          .attr("y", barHeight / 2)
          .attr("dy", ".35em")
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
/* .chart div {
  font: 10px sans-serif;
  background-color: steelblue;
  text-align: right;
  padding: 3px;
  margin: 1px;
  color: white;
} */

.chart rect {
  fill: steelblue;
}

.chart text {
  fill: white;
  font: 10px sans-serif;
  text-anchor: end;
}
</style>
