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

      const width = 960, height = 500

      d3.select(".chart")
        .attr("width", width)
        .attr("height", height)

      const margin = ({top: 20, right: 0, bottom: 30, left: 40})

      d3.json("http://canon.atec.pub/Maslin,%20T.%20Paul/Occurrence%20of%20the%20Garter%20Snake,%20Thamnophis%20sirtalis,%20in%20the%20Great%20Plains%20and%20Rocky%20Mountains.json").then(function(entities) {
        
        // eslint-disable-next-line
        console.log("in callback")

        var data = []

        for (var entity in entities) {
          if (entities.hasOwnProperty(entity)) {
              data.push({ "text": entity, "count": entities[entity] })
          }
        }

        // eslint-disable-next-line
        console.log(data)

        const x = d3.scaleBand()
          .domain(data.map(d => d.text))
          .range([margin.left, width - margin.right])
          .padding(0.1)

        const y = d3.scaleLinear()
          .domain([0, d3.max(data, d => d.count)]).nice()
          .range([height - margin.bottom, margin.top])

        const xAxis = g => g
          .attr("transform", `translate(0,${height - margin.bottom})`)
          .call(d3.axisBottom(x).tickSizeOuter(0))

        const yAxis = g => g
          .attr("transform", `translate(${margin.left},0)`)
          .call(d3.axisLeft(y))
          .call(g => g.select(".domain").remove())

        const svg = d3.selectAll("svg")
      
        // eslint-disable-next-line
        console.log(svg)

        svg.append("g")
            .attr("fill", "steelblue")
          .selectAll("rect")
          .data(data)
          .enter().append("rect")
          // .join("rect")
            .attr("x", d => x(d.text))
            .attr("width", x.bandwidth())
            .attr("y", d => y(d.count))
            .attr("height", d => y(0) - y(d.count))
      
        // eslint-disable-next-line
        console.log(svg)
        
        svg.append("g")
          .call(xAxis)
      
        // eslint-disable-next-line
        console.log(svg)
        
        svg.append("g")
          .call(yAxis)
      
        // eslint-disable-next-line
        console.log(svg)
      })
    }
  },

  mounted: function () {

    // eslint-disable-next-line
    console.log("mounted")

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
