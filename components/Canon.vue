<template>
  <div>

    <section class="hero">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            canon 
          </h1>
          <h2 class="subtitle">
            an attempt to get meaning out of text
          </h2>
        </div>
      </div>
    </section>

    <img src="../crawl.gif">

    <Chart></Chart> 

  </div>
</template>

<script>
import Header from './Header.vue'
import Chart from './Chart.vue'

export default {
  components: {
    Header,
    Chart
  }
}
</script>
