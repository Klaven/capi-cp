<template>
  <Header title="Overview"></Header>
  <NavBar></NavBar>
  <div class="space">
    <MgmtClusterList></MgmtClusterList>
    <WorkloadClusterList></WorkloadClusterList>
  </div>
</template>

<script>
import Header from './components/Header'
import NavBar from './components/NavBar'
import MgmtClusterList from './components/management/ClusterList'
import WorkloadClusterList from './components/workload/ClusterList'

export default {
  name: 'App',
  components: {
    Header,
    NavBar,
    MgmtClusterList,
    WorkloadClusterList,
  }
}
</script>

<style>
html, body {
  height: 100%
}

body {
  background: whitesmoke;
}

#app {
  height: 100%;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.navbar {
  padding-top: 20px;
  float: left;
}

.space {
  width: 100%;
}
</style>
