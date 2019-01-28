import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import Argon from "./plugins/argon-kit";
import VueYoutube from "vue-youtube";
import VueNativeSock from "vue-native-websocket";
import { wsEndpoint } from "./env";
import VueMq from "vue-mq";
import Notifications from "vue-notification";

Vue.use(VueNativeSock, `${wsEndpoint}/coordinator/ws`, {
  connectManually: true,
  format: "json"
});

Vue.use(VueMq, {
  breakpoints: {
    sm: 690,
    md: 1475,
    lg: Infinity
  }
});

Vue.use(Notifications);
Vue.config.productionTip = false;
Vue.use(Argon);
Vue.use(VueYoutube);

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
