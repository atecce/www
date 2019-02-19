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

    <select v-model="currentAuthor">
      <option v-for="author in authors">
        {{ author }}
      </option>
    </select>
    <br>

    <select v-model="currentWork">
      <option v-for="work in works">
        {{ work }}
      </option>
    </select>
    <br>

    <button @click="getEntities">Get Entities</button>

    <div class="scroller">
      <svg class="chart"></svg>
    </div>
  </div>
</template>

<script>
import * as d3 from 'd3';

export default {
  data () {
    return {
      root: "http://localhost:8081/",
      currentAuthor:  "",
      currentWork: ""
    }
  },
  computed: {
    authors() {
      return this.get(this.root)
    },
    works() {
      return this.get(this.root + this.currentAuthor)
    }
  },
  methods: {
    get(url) {
      console.log(url)
      var req = new XMLHttpRequest();
      req.open('GET', url, false);
      req.send(null);
      return JSON.parse(req.responseText);
    },
    getEntities() {

      const width = 45000, height = 500

      d3.select(".chart")
        .attr("width", width)
        .attr("height", height)

      const margin = ({top: 20, right: 0, bottom: 30, left: 40})

      d3.json(this.root+this.currentAuthor+"/"+this.currentWork).then(function(entities) {
        
        // eslint-disable-next-line
        console.log("in callback")

        var data = []
        for (var entity in entities) {
          if (entities.hasOwnProperty(entity)) {
              data.push({ "text": entity, "count": entities[entity] })
          }
        }

        data.sort(function(a, b) {
          return a.count < b.count
        })

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
.chart rect {
  fill: steelblue;
}

.chart text {
  fill: black;
  font: 10px sans-serif;
  text-anchor: end;
}

.scroller {
  width: 960;
  height: 500;
  overflow-x: scroll;
}
</style>
