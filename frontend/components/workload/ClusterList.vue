<template>
  <div class="wrapper">
    <p><b>Workload clusters</b></p>

    <div class="cluster-list">
      <div class="header">
        <CreateButton></CreateButton>
      </div>
      <table>
        <tr>
          <th>Name</th>
          <th>Available</th>
          <th>Pending</th>
          <th>Fail</th>
          <th>Management</th>
          <th>Provider</th>
          <th>Created</th>
          <th>Updated</th>
        </tr>
        <tr>
          <td><a href="">capi-workload-cluster-00</a></td>
          <td><b class="available">5</b>/<b class="total">5</b></td>
          <td><b class="pending">0</b></td>
          <td><b class="failed">0</b></td>
          <td><a href="">capi-mgmt-cluster-01</a></td>
          <td><a href="">Cluster API</a></td>
          <td>26-08-2021:12:00:00</td>
          <td>26-08-2021:12:00:00</td>
        </tr>
        <tr>
          <td><a href="">capi-workload-cluster-01</a></td>
          <td><b class="pending">4</b>/<b class="total">5</b></td>
          <td><b class="pending">1</b></td>
          <td><b class="failed">0</b></td>
          <td><a href="">capi-mgmt-cluster-01</a></td>
          <td><a href="">Cluster API</a></td>
          <td>26-08-2021:12:00:00</td>
          <td>26-08-2021:12:00:00</td>
        </tr>
        <tr>
          <td><a href="">capi-workload-cluster-02</a></td>
          <td><b class="failed">3</b>/<b class="total">5</b></td>
          <td><b class="pending">1</b></td>
          <td><b class="failed">1</b></td>
          <td><a href="">capi-mgmt-cluster-01</a></td>
          <td><a href="">Cluster API</a></td>
          <td>26-08-2021:12:00:00</td>
          <td>26-08-2021:12:00:00</td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import CreateButton from './CreateButton'

export default {
  name: "ClusterList",
  components: {
    CreateButton,
  }
}
</script>

<style scoped>
p {
  font-size: 1.5em;
  text-align: left;
  color: dimgray;
}

table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
}

th {
  color: gray;
}

th, td {
  padding: 8px;
}

tr:nth-child(even) {
  background-color: whitesmoke;
}

.available {
  color: green;
}

.total {
  color: green;
}

.pending {
  color: orange;
}

.failed {
  color: red;
}

.wrapper {
  float: left;
  width: 50%
}

.cluster-list {
  box-shadow: 0 3px 1px -2px rgba(0, 0, 0, .2), 0 2px 2px 0 rgba(0, 0, 0, .14), 0 1px 5px 0 rgba(0, 0, 0, .12);
  background: white;
  border-radius: 5px;
}

.header {
  display: inline-block;
  width: 100%;
}

.cluster-cards {
  padding: 10px;
}

.create-wrk-clstr-btn {
  float: right;
  padding-right: 10px;
  padding-top: 10px;
}
</style>