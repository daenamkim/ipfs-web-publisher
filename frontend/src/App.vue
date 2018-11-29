<template>
  <div class="app">
    <v-app>
      <v-container fluid>
        <Login v-if="$store.state.view === 'login'" />
        <div v-else>
          <Navbar />
          <Publish />
        </div>
      </v-container>
    </v-app>
  </div>
</template>

<script>
import Navbar from "./components/Navbar";
import Login from "./components/Login";
import Publish from "./components/Publish";

export default {
  name: "app",
  components: {
    Navbar,
    Login,
    Publish
  }
};
</script>

<style>
.app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
}

h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
