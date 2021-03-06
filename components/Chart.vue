<template>
  <div>

    <br>
    <div class="field control has-icons-left">
      <input list="authors" class="input" placeholder="search for authors" 
             v-model="searchText" v-on:keyup="searchAuthors"/>
      <span class="icon is-large is-left">
        <i class="fa fa-search"></i>
      </span>
    </div>

    <!-- <datalist v-if="isDesktopFirefox()" id="authors">
      <option v-for="author in authors" :key="author" @change="getWorks">
        {{ author }}
      </option>
    </datalist>
    <div v-else class="select"> -->
    <div class="select">
      <select v-model="selectedAuthor" @change="getWorks">
        <option v-for="author in authors" :key="author">
          {{ author }}
        </option>
      </select>
    </div>

    <div class="select">
      <select v-model="currentWork" @change="getEntities">
        <option v-for="work in works" :key="work">
          {{ work }}
        </option>
      </select>
    </div>
    <br><br>
    
    <div class="scroller">
      <svg class="chart"></svg>
    </div>
  </div>
</template>

<script>
import * as d3 from 'd3'

export default {
  data () {
    return {

      root: "https://canon.atec.pub/",

      searchText: "",
      selectedAuthor: "",

      currentWork: "",

      authors: [],
      works: [],
    }
  },
  computed: {
    currentAuthor() {
      return this.selectedAuthor == "" ? this.searchText : this.selectedAuthor
    }
  },
  methods: {

    isDesktopFirefox: function() {
      return !navigator.userAgent.match(/AppleWebKit/) && navigator.userAgent.match(/Gecko/)
    },

    // TODO use fetch
    get(url) {
      var req = new XMLHttpRequest();
      req.open('GET', url.replace(/\(/g, "%28").replace(/\)/g, "%29"), false);

      req.send(null);
      if (req.status != 200) {
        return
      }

      try {
        return JSON.parse(req.responseText);
      } catch {
        return
      }
    },

    searchAuthors() {
      fetch(this.root + "authors?search=" + this.searchText).then(res => {
        res.body.getReader().read().then(({ value }) => {
          this.authors = new TextDecoder("utf-8").decode(value).split("\n")
          return 
        })
      })
    },

    getWorks() {
      this.works = this.get(this.root + "authors/" + this.currentAuthor)
    },

    getEntities() {

      const url = this.root + "authors/" + this.currentAuthor + "/works/" + this.currentWork

      d3.json(url.replace(/\(/g, "%28").replace(/\)/g, "%29")).then(function(entities) {

        // convert map to array of objects and sort
        var data = []
        for (var entity in entities) {
          if (entities.hasOwnProperty(entity)) {
              data.push({ "text": entity, "count": entities[entity] })
          }
        }

        data.sort(function(a, b) {
          return b.count - a.count
        })

        // create chart
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
          .selectAll(".tick text")
            .call(function(text, width) {
              text.each(function() {
                var text = d3.select(this),
                    words = text.text().split(/\s+/).reverse(),
                    word,
                    line = [],
                    lineNumber = 0,
                    lineHeight = 1.1, // ems
                    y = text.attr("y"),
                    dy = parseFloat(text.attr("dy")),
                    tspan = text.text(null).append("tspan").attr("x", 0).attr("y", y).attr("dy", dy + "em");

                // eslint-disable-next-line
                while (word = words.pop()) {
                  line.push(word);
                  tspan.text(line.join(" "));
                  if (tspan.node().getComputedTextLength() > width) {
                    line.pop();
                    tspan.text(line.join(" "));
                    line = [word];
                    tspan = text.append("tspan").attr("x", 0).attr("y", y).attr("dy", ++lineNumber * lineHeight + dy + "em").text(word);
                  }
                }
              })
            }, x.bandwidth())
        
        svg.append("g")
          .call(yAxis)
      })
    },
    
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

.input {
  width: 300px;
  margin-left: 10px;
}

select {
  margin-left: 10px;
}

span {
  margin-left: 10px;
}
</style>
