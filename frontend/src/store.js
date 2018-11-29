import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    publishSelected: {},
    publishList: [],
    view: "123"
  },
  mutations: {
    updatePublishList () {
      // TODO: 
    },
    switchView (state, view) {
      state.view = view;
    }
  },
  actions: {
    addPublish ({ commit }, id) {
      const param = id !== undefined ? id : "";
      // TODO: 
      axios.get(`http://localhost:8081/publish/${param}`)
        .then(response => {
          console.log(response);
          commit("updatePublishList", response);
        })
        .catch()
    },
    updatePublish () {
      // TODO: 
    },
    getPublish () {
      // TODO: 
    },
    switchView ({ commit }, view) {
      commit("switchView", view)
    }
  }
});
