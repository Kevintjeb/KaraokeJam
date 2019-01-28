<template>
  <section class="section-hero section-shaped my-0">
    <div class="shape shape-style-1 shape-primary">
      <span class="span-150"></span>
      <span class="span-50"></span>
      <span class="span-100"></span>
      <span class="span-50"></span>
      <span class="span-75"></span>
      <span class="span-75"></span>
      <span class="span-50"></span>
      <span class="span-100"></span>
      <span class="span-50"></span>
      <span class="span-100"></span>
    </div>
    <div class="wrapper-client">
      <div class="side-client">
        <Side @signout="signout()" @click="modalOpen = true"/>
      </div>
      <div class="queue-client">
        <Queue :addSongToQueue="addSongToQueue"/>
      </div>
      <div>
        <modal
          :show.sync="modalOpen"
          gradient="danger"
          modal-classes="modal-danger modal-dialog-centered"
        >
          <template slot="header">
            <h4 class="modal-title" id="exampleModalLabel">Room code</h4>
          </template>
          <div class="py-3 text-center">
            <qrcode-vue :value="value"></qrcode-vue>
          </div>
        </modal>
      </div>
    </div>
  </section>
</template>


<script>
import Side from "../layout/Side";
import Queue from "../layout/Queue";
import Player from "../layout/Player";
import QrcodeVue from "qrcode.vue";
import { store } from "../store";
export default {
  name: "App",
  components: {
    Side,
    Queue,
    Player,
    QrcodeVue
  },
  data() {
    return {
      modalOpen: false,
      value: store.state.roomKey
    };
  },
  methods: {
    signout() {
      store.reset();
      this.$router.replace("/");
    },
    addSongToQueue(song) {
      store.addSongToQueue(song);
    }
  }
};
</script>

<style lang="scss">
@import "../assets/scss/definitions.scss";

.side-client {
  grid-area: side;
  visibility: visible;
  overflow: auto;
  height: 100vh;
  width: 25vw;
  padding: 10px;
}

.queue-client {
  grid-area: queue;
  height: 25vh;
  width: 75vw;
  padding: 10px;
}
.wrapper-client {
  grid-template-columns: 1fr;
  grid-template-rows: 1fr;
  grid-template-areas:
    "side"
    "queue";
}

.side-client {
  width: 100vw;
  height: 75vh;
  display: block;
}

.queue-client {
  width: 100vw;
  display: block;
}

.screen-client {
  display: none;
}
</style>
