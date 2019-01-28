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
            <div :class="{'qrcode-container': isInit}">
              <qrcode-stream @init="onInit" @decode="onDecode">
                <img v-if="isInit" class="overlay-qrcode" src="../../static/img/qr-code.png">
                <stretch v-else-if="isError !== true"/>
                <p v-else>Please allow access to your camera.</p>
              </qrcode-stream>
            </div>
          </card>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import { QrcodeStream } from "vue-qrcode-reader";
import { store } from "../store";
import { RoomService } from "../services/RoomService";
import Stretch from "../components/Stretch";

export default {
  data() {
    return {
      service: new RoomService(),
      isInit: false,
      isError: false
    };
  },
  components: {
    QrcodeStream,
    Stretch
  },
  methods: {
    onDecode(decodedString) {
      this.service.joinRoom(decodedString).then(result => {
        store.setRoomKey(result.data.Room.Key);
        store.setUserKey(result.data.User.Key);
        this.$router.push("client");
      });
    },
    onInit(promise) {
      promise
        .then(() => {
          this.isInit = true;
        })
        .catch(() => {
          this.$notify({
            type: "error",
            title: "Camera error",
            text: "Please allow camera access"
          });
          this.isError = true;
        });
    }
  }
};
</script>
<style lang="scss">
.qrcode-stream__overlay {
  display: flex;
  justify-content: center;
  align-items: center;
}

.qrcode-container {
  border: 1px SOLID black;
  border-radius: 20px;
  overflow: hidden;
}
</style>
