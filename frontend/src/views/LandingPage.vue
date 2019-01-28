<template>
  <section class="section section-shaped section-lg my-0">
    <div class="shape shape-style-1 shape-primary">
      <span class="span-150"></span>
      <span class="span-50"></span>
      <span class="span-50"></span>
      <span class="span-75"></span>
      <span class="span-100"></span>
      <span class="span-75"></span>
      <span class="span-50"></span>
      <span class="span-100"></span>
      <span class="span-50"></span>
      <span class="span-100"></span>
    </div>
    <div class="container pt-lg-md">
      <div class="row justify-content-center">
        <div class="col-lg-5">
          <card
            type="secondary"
            shadow
            header-classes="bg-white pb-5"
            body-classes="px-lg-5 py-lg-5"
            class="border-0"
          >
            <div class="text-center text-muted mb-4" id="logo">KaraokeJam</div>
            <div class="text-center" id="all">Welcome to the karaoke platform.
              <div class="buttons">
                <base-button class="button my-4" id="buttons" v-on:click="handleHostClick()">HOST</base-button>
                <base-button class="button my-4" id="buttons" v-on:click="handleJoinClick()">JOIN</base-button>
              </div>
              <span>Let's jam and get started!</span>
            </div>
          </card>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import { store } from "../store";
import { RoomService } from "../services/RoomService";

export default {
  data() {
    return {
      service: new RoomService()
    };
  },
  methods: {
    handleHostClick() {
      this.service.newRoom().then(response => {
        store.setRoomKey(response.data.Room.Key);
        store.setUserKey(response.data.User.Key);

        this.$router.replace("host");
      });
    },
    handleJoinClick() {
      this.$router.replace("joinroom");
    }
  }
};
</script>
<style lang="scss">
@import "../assets/scss/definitions.scss";

#logo {
  font-family: HarlowSolid;
  color: #fff;
  text-align: center;
  font-size: 42pt;
}

#all {
  font-size: 18px;
}

#buttons {
  font-size: 32px;
  margin: 0;
}
.buttons {
  display: flex;
  justify-content: space-around;
  align-items: center;
  .button {
    background-color: $button-primary !important;
    border: none !important;
  }
}

.section-lg {
  padding-top: 0 !important;
  padding-bottom: 0 !important;
}

@media (max-width: 450px) {
  #logo {
    font-size: 42px;
  }

  #buttons {
    font-size: 22px !important;
  }
}
</style>
