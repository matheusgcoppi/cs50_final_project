<template>
  <div class="header">
    <a class="logo" :href="'/'">Barber Finance</a>
    <div class="links">
      <a :href="'/'" :class="{ active: activePage === 'finance' }">Finance</a>
      <a :href="'/income'" :class="{ active: activePage === 'income' }"
        >Income</a
      >
      <a :href="'/expense'" :class="{ active: activePage === 'expense' }"
        >Expense</a
      >
      <a :href="'/about'" :class="{ active: activePage === 'about' }">About</a>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    activePage: String,
  },
};
</script>

<style scoped>
* {
  margin: 0;
}

.header {
  background-color: #3498db;
  width: 100vw;
  height: 80px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  margin-left: 300px;
  font-weight: bold;
  color: white;
  font-size: 22px;
  text-decoration: none;
}

.links {
  display: flex;
  gap: 40px;
  margin-right: 180px;
}

.links a {
  text-decoration: none;
  font-weight: bold;
  color: white;
  font-size: 22px;
  position: relative;
}

.links a.active::after {
  content: "";
  position: absolute;
  bottom: -5px;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: white;
}
</style>
