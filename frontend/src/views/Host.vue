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
    <div class="wrapper">
      <div :class="{ 'side': !sideOpen, 'sideOpened': sideOpen }">
        <div v-if="$mq === 'lg' || sideOpen" class="side-flex-container">
          <Side class="side-container" @signout="signout()" @click="modalOpen = true"/>
          <div v-if="$mq !== 'lg'" v-on:click="sideOpen = !sideOpen">
            <div class="button-container">
              <i
                :class="{ 'ni ni-bold-right': !sideOpen, 'ni ni-bold-left': sideOpen }"
                :id="{'side-icon': !sideOpen, 'side-icon-open': sideOpen}"
              ></i>
            </div>
          </div>
        </div>
        <div v-else v-on:click="sideOpen = !sideOpen">
          <card class="border-0" id="container">
            <i class="ni ni-bold-right" id="side-icon"></i>
          </card>
        </div>
      </div>
      <div
        class="screen"
        :class="{ 'side-closed': !sideOpen, 'side-opened': sideOpen }"
        v-if="state.currentVideoId !== undefined"
      >
        <Player :videoId="state.currentVideoId" :skipSong="nextSong"/>
      </div>
      <div v-else id="logo-player">KaraokeJam</div>
      <div class="queue" :class="{ 'side-closed': !sideOpen, 'side-opened': sideOpen }">
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
            <qrcode-vue :value="state.roomKey"></qrcode-vue>
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
      state: store.state,
      modalOpen: false,
      value: store.state.roomKey,
      sideOpen: false
    };
  },
  methods: {
    nextSong() {
      store.nextSong();
    },
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
@font-face {
  font-family: HarlowSolid;
  src: url("../../public/fonts/HarlowSolidItalic.ttf");
}

#logo-player {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 72;
  color: rgba(255, 255, 255, 0.3);

  font-family: HarlowSolid;
}

.wrapper {
  display: grid;
  grid-template-columns: auto 3fr;
  grid-template-rows: 3fr 1fr;
  grid-template-areas:
    "side screen"
    "side queue";
}

.side {
  grid-area: side;
  visibility: visible;
  overflow: auto;
  height: 100vh;
  width: 25vw;
  padding: 10px;
}

.screen {
  grid-area: screen;
  padding: 10px;
  height: 100%;
  width: 100%;
}

.queue {
  grid-area: queue;
  height: 25vh;
  width: 75vw;
  padding: 10px;
}

#side-icon {
  text-align: center;
  font-size: 20px;
  position: absolute;
  left: 0;
  right: 0;
  top: 50%;
}

.side-flex-container {
  display: flex;
}

.sideOpened {
  height: 100%;
}

@media (max-width: 1475px) {
  .wrapper {
    display: grid;
    grid-template-columns: 1fr auto;
    grid-template-rows: 3fr 1fr;
    grid-template-areas:
      "side screen"
      "side  queue";
  }
  .side {
    width: 50px;
    overflow: hidden;
  }
  .queue {
    width: calc(100vw);
  }

  .sideOpened {
    grid-area: side;
    visibility: visible;
    overflow: auto;
    height: 100vh;
    width: 100%;
    padding: 10px;
    float: left;
  }

  #side-icon-open {
    text-align: center;
    font-size: 20px;
    position: static;
    left: 0;
    right: 0;
    top: 50%;
  }

  .button-container {
    width: 20px;
    height: 100%;
    background-color: rgba(244, 245, 247, 0.7);
    // background-color: red;
    z-index: 999;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  #logo-player {
    font-size: calc(40px + 2vw);
  }

  .side-flex-container {
    display: flex;
  }

  .side-container {
    border-radius: 6px;
  }

  .side-opened {
    width: calc(100vw - 420px) !important;
  }

  .side-closed {
    width: calc(100vw - 50px) !important;
  }
}
</style>
