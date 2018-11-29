<template>
  <div class="publish">
    <v-container fluid>
      <v-layout row justify-center class="new-publish">
        <v-dialog v-model="dialog" persistent max-width="600px">
          <v-btn slot="activator" color="primary" dark>New Web Publish</v-btn>
          <v-card>
            <v-card-title>
              <span class="headline">User Profile</span>
            </v-card-title>
            <v-card-text>
              <v-container grid-list-md>
                <v-layout wrap>
                  <v-flex xs12 sm6 md4>
                    <v-text-field label="Legal first name*" required></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md4>
                    <v-text-field label="Legal middle name" hint="example of helper text only on focus"></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md4>
                    <v-text-field
                      label="Legal last name*"
                      hint="example of persistent helper text"
                      persistent-hint
                      required
                    ></v-text-field>
                  </v-flex>
                  <v-flex xs12>
                    <v-text-field label="App Name*" required></v-text-field>
                  </v-flex>
                </v-layout>
              </v-container>
              <small>*indicates required field</small>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" flat @click="dialog = false">Close</v-btn>
              <v-btn color="blue darken-1" flat @click="dialog = false">Save</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-layout>















        <v-layout column>
          <v-flex xs12>
      <!-- <v-dialog v-model="dialog" max-width="500px">
        <v-spacer />
        <v-btn slot="activator" color="primary" dark class="mb-2">New Item</v-btn>
        <v-btn slot="activator" fab dark color="primary">
          <v-icon dark>add</v-icon>
        </v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">{{ formTitle }}</span>
          </v-card-title>
          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 sm6 md4>
                  <v-text-field v-model="editedItem.name" label="App Name"></v-text-field>
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-text-field v-model="editedItem.calories" label="App File"></v-text-field>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" flat @click="close">Cancel</v-btn>
            <v-btn color="blue darken-1" flat @click="save">Save</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog> -->
      <v-data-table
        :headers="headers"
        :items="publishList"
        class="elevation-1"
        hide-actions
      >
        <template slot="items" slot-scope="props">
          <td class="text-xs-center">{{ props.item.Id }}</td>
          <td class="text-xs-center">{{ props.item.Name }}</td>
          <td class="text-xs-center">{{ props.item.GatewayUrl }}</td>
          <td class="text-xs-center">{{ props.item.UpdatedAt }}</td>
          <td class="justify-center layout px-0">
            <v-icon
              small
              class="mr-2"
              @click="editItem(props.item)"
            >
              edit
            </v-icon>
            <v-icon
              small
              @click="deleteItem(props.item)"
            >
              delete
            </v-icon>
          </td>


        </template>
      </v-data-table>
          </v-flex>
      </v-layout>
      </v-container>
  </div>
</template>

<script>
import moment from 'moment';

export default {
  name: "Publish",
  computed: {
    formTitle () {
      return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
    }
  },
  data: () => {
    return {
      dialog: false,
      headers: [
        { text: 'ID', align: 'center', value: 'Id' },
        { text: 'App Name', align: 'center', value: 'Name' },
        { text: 'Gateway URL', align: 'center', value: 'GatewayUrl' },
        { text: 'Updated At', align: 'center', value: 'UpdatedAt' },
        { text: 'Actions', align: 'center', sortable: false }
      ],
      publishList: [
        {
          value: false,
          Id: 2,
          Name: 'Frozen Yogurt3',
          GatewayUrl: 'test2',
          UpdatedAt: moment(new Date()).fromNow()
        },
        {
          value: false,
          Id: 1,
          Name: 'Frozen Yogurt1',
          GatewayUrl: 'test3',
          UpdatedAt: moment(new Date()).fromNow()
        },
        {
          value: false,
          Id: 3,
          Name: 'Frozen Yogurt2',
          GatewayUrl: 'test1',
          UpdatedAt: moment(new Date()).fromNow()
        },
        {
          value: false,
          Id: 5,
          Name: 'Frozen Yogurt4',
          GatewayUrl: 'test7',
          UpdatedAt: moment(new Date()).fromNow()
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
      // this.editedIndex = this.desserts.indexOf(item)
      // this.editedItem = Object.assign({}, item)
      this.dialog = true
    },
    deleteItem (item) {
      // const index = this.desserts.indexOf(item)
      alert("delete")
      // confirm('Are you sure you want to delete this item?') && this.desserts.splice(index, 1)
    },
    close () {
      this.dialog = false
      // setTimeout(() => {
      //   this.editedItem = Object.assign({}, this.defaultItem)
      //   this.editedIndex = -1
      // }, 300)
    },
    save () {
      // if (this.editedIndex > -1) {
      //   Object.assign(this.desserts[this.editedIndex], this.editedItem)
      // } else {
      //   this.desserts.push(this.editedItem)
      // }
      this.close()
    }
  },
  mounted() {
    this.$store.dispatch('addPublish', 23);
  }
}

</script>

<style scoped>
.new-publish {
  margin-top: 20px;
  margin-bottom: 20px;
}
</style>
