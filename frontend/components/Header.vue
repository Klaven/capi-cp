<template>
  <div class="header">
    <p>{{ title }}</p>
  </div>
</template>

<script>
export default {
  name: "Header",
  props: ["title"],
}
</script>

<style scoped>
  .header {
    display: flex;
    background: royalblue;
    align-items: center;
    width: 100%;
    height: 70px;
  }
  .header p {
    font-size: 2em;
    float: left;
    color: white;
    padding-left: 20px;
  }
</style>