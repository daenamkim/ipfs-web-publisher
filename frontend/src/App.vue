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
import HelloWorld from "./components/HelloWorld.vue";
import Login from "./components/Login";
import Publish from "./components/Publish";

export default {
  name: "app",
  components: {
    HelloWorld,
    Login,
    Publish
  },
  data () {
    return {
      headers: [
        { text: 'ID', value: 'Id' },
        { text: 'App Name', align: 'left', value: 'Name' },
        { text: 'Gateway URL', align: 'left', value: 'GatewayUrl' },
        { text: 'Updated At', align: 'left', value: 'UpdatedAt' }
      ],
      publishList: [
        {
          value: false,
          Id: 2,
          Name: 'Frozen Yogurt3',
          GatewayUrl: 'test2',
          UpdatedAt: new Date().toString()
        },
        {
          value: false,
          Id: 1,
          Name: 'Frozen Yogurt1',
          GatewayUrl: 'test3',
          UpdatedAt: new Date().toString()
        },
        {
          value: false,
          Id: 3,
          Name: 'Frozen Yogurt2',
          GatewayUrl: 'test1',
          UpdatedAt: new Date().toString()
        },
        {
          value: false,
          Id: 5,
          Name: 'Frozen Yogurt4',
          GatewayUrl: 'test7',
          UpdatedAt: new Date().toString()
        }
      ],
      editedIndex: -1,
      editedItem: {
        name: '',
        calories: 0,
        fat: 0,
        carbs: 0,
        protein: 0
      },
      defaultItem: {
        name: '',
        calories: 0,
        fat: 0,
        carbs: 0,
        protein: 0
      }
    }
  },
  methods: {
    editItem (item) {
      this.editedIndex = this.desserts.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },
    deleteItem (item) {
      const index = this.desserts.indexOf(item)
      confirm('Are you sure you want to delete this item?') && this.desserts.splice(index, 1)
    },
    close () {
      this.dialog = false
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      }, 300)
    },
    save () {
      if (this.editedIndex > -1) {
        Object.assign(this.desserts[this.editedIndex], this.editedItem)
      } else {
        this.desserts.push(this.editedItem)
      }
      this.close()
    }
  },
  mounted() {
    this.$store.dispatch('addPublish', 23);
  }
};
</script>

<style>
.app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
