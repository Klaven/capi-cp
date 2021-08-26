<template>
  <div class="navbar">
    <ul>
      <li><a href=""><b>Management clusters</b></a></li>
      <li><a href=""><b>Workload clusters</b></a></li>
      <li><a href=""><b>Providers</b></a></li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "NavBar"
}
</script>

<style scoped>
  * {
    margin: 0px;
    padding: 0px;
  }

  .navbar {
    width: 250px;
    height: 100%;
  }

  .navbar ul {
    padding-left: 25px;
    text-align: left;
    list-style: none;
    line-height: 50px;
    list-style: '+ ';
  }

  .navbar a {
    color: black;
    text-decoration: none;
  }
</style>