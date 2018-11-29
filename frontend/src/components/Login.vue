<template>
  <div class="login">
    <div>
      <img class="logo" src="../assets/logo.png" />
    </div>
    <div class="title">
      <span>Web Publisher</span>
    </div>
    <v-btn color="primary" v-on:click="login">
      Login
    </v-btn>
  </div>
</template>

<script>
export default {
  name: "Login",
  methods: {
    login() {
      this.$store.dispatch('switchView', 'publish')
    }
  }
}
</script>

<style socped>
.login {
  text-align: center;
}

.title {
  margin-top: 20px;
  margin-bottom: 20px;
}

.title span {
  font-weight: bold;
  font-size: 40px;
  color: #5f5f5f;
}

.logo {
  width: 400px;
  height: 400px;
}
</style>
