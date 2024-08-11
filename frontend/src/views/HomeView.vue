<template>
  <Header :active-page="activePage" />
  <Chart />
</template>

<script>
import Header from "@/components/Header.vue";
import { ref } from "vue";
import Chart from "@/components/Charts.vue";

export default {
  name: "HomeView",
  components: {
    Header: Header,
    Chart: Chart,
  },
  setup() {
    const activePage = ref("finance");
    return { activePage };
  },
};
</script>

<style>
* {
  margin: 0;
}
</style>