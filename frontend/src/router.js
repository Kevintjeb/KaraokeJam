import Vue from "vue";
import Router from "vue-router";
import AppHeader from "./layout/AppHeader";
import AppFooter from "./layout/AppFooter";
//import Components from "./views/Components.vue";
//import Landing from "./views/Landing.vue";
import Host from "./views/Host.vue";
import Client from "./views/Client.vue";
import LandingPage from "./views/LandingPage.vue";
import { store } from "./store";
import JoinRoom from "./views/JoinRoom";
Vue.use(Router);

export default new Router({
  linkExactActiveClass: "active",
  routes: [
    {
      path: "/",
      name: "landingpage",
      components: {
        header: AppHeader,
        default: LandingPage,
        footer: AppFooter
      }
    },
    {
      path: "/joinroom",
      name: "joinroom",
      components: {
        header: AppHeader,
        default: JoinRoom,
        footer: AppFooter
      }
    },
    {
      path: "/host",
      name: "host",
      components: {
        header: AppHeader,
        default: Host,
        footer: AppFooter
      },
      beforeEnter: (to, from, next) => {

        if (store.hasRoomKey()) {
          store.setIsClient(false);
          return next(true);
        }
        else return next("/");
      }
    },
    {
      path: "/client",
      name: "client",
      components: {
        header: AppHeader,
        default: Client,
        footer: AppFooter
      },
      beforeEnter: (to, from, next) => {
        if (store.hasRoomKey()) {
          store.setIsClient(true);
          return next(true);
        }
        else {
          return next("/");
        }
      }
    }
  ],
  scrollBehavior: to => {
    if (to.hash) {
      return { selector: to.hash };
    } else {
      return { x: 0, y: 0 };
    }
  }
});
