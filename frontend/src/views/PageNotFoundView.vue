<template>
  <h1>404 - NOT FOUND</h1>
  <p>Sorry, the page you are looking for does not exist.</p>
  <img src="../assets/dog-not-found.png" alt="dog" />
  <div>
    <a v-bind:href="'/login'">Go back to the Login page</a>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
img {
  width: 300px;
}
div a {
  text-decoration: none;
}
</style>
