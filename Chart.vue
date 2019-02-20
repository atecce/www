<template>
  <div>
    <select v-model="currentAuthor">
      <option v-for="author in authors" :key="author">
        {{ author }}
      </option>
    </select>
    <br>

    <select v-model="currentWork">
      <option v-for="work in works" :key="work">
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

      root: "https://canon.atec.pub/",

      currentAuthor: "",
      currentWork: "",

      authors: [],
    }
  },
  computed: {
    works() {
      return this.get(this.root + this.currentAuthor)
    }
  },
  methods: {
    get(url) {
      var req = new XMLHttpRequest();
      req.open('GET', url, false);
      req.send(null);
      try {
        return JSON.parse(req.responseText);
      } catch {
        return
      }
    },
    fetchAuthors() {
      fetch(this.root).then(res => {

        const r = res.body.getReader()

        r.read().then(({ done, value }) => {

          const newAuthors = new TextDecoder("utf-8").decode(value).split("\n")

          this.authors = this.authors.concat(newAuthors)

          if (done) {
            return
          }
        })
      })
    },

    getEntities() {

      d3.json(this.root+this.currentAuthor+"/"+this.currentWork).then(function(entities) {

        var data = []
        for (var entity in entities) {
          if (entities.hasOwnProperty(entity)) {
              data.push({ "text": entity, "count": entities[entity] })
          }
        }

        data.sort(function(a, b) {
          return a.count < b.count
        })

        const width = data.length * 75, height = 500

        d3.select(".chart")
         .selectAll("g").remove()

        d3.select(".chart")
          .attr("width", width)
          .attr("height", height)

        const margin = ({top: 20, right: 0, bottom: 30, left: 40})

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
            .attr("x", d => x(d.text))
            .attr("width", 50)
            .attr("y", d => y(d.count))
            .attr("height", d => y(0) - y(d.count))
        
        svg.append("g")
          .call(xAxis)
        
        svg.append("g")
          .call(yAxis)
      })
    }
  },
  mounted: function() {
      this.fetchAuthors()
  },
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
  margin-left: 10px;
  width: 960;
  height: 500;
  overflow-x: scroll;
}

select {
  margin-left: 10px;
}

button {
  margin-left: 10px;
}
</style>
