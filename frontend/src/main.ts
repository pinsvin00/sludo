import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
// import {
//   MdButton,
//   MdContent,
//   MdTabs,
//   MdField,
//   MdSwitch,
// } from "vue-material/dist/components";
import VueMaterial from "vue-material";
import "vue-material/dist/vue-material.min.css";
import "vue-material/dist/theme/default.css";

// Vue.use(MdButton);
// Vue.use(MdContent);
// Vue.use(MdTabs);
// Vue.use(MdField);
// Vue.use(MdSwitch);

Vue.use(VueMaterial);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
