<template>
  <div class="navbar">
    <v-container>
      <v-layout column>
        <v-flex xs12>
          <v-toolbar color="primary" dark fixed>
            <v-toolbar-title>IPFS Web Publisher</v-toolbar-title>
            <v-spacer></v-spacer>
            <!-- TODO: mobile friendly -->
            <v-toolbar-items class="hidden-sm-and-down">
              <v-btn flat v-on:click="logout">Logout</v-btn>
            </v-toolbar-items>
          </v-toolbar>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "Navbar",
  methods: {
    logout() {
      this.$store.dispatch("switchView", "login");
    }
  }
}
</script>
