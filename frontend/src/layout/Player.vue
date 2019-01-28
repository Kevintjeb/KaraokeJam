<template>
  <card class="border-0" id="player-container">
    <youtube
      :video-id="this.videoId"
      ref="youtube"
      @playing="playing"
      @ready="ready"
      @paused="paused"
      @ended="ended"
      :player-vars="playerVars"
    ></youtube>

    <div class="floating-control-container">
      <div class="player-button">
        <base-button
          v-on:click="togglePlayer()"
          rounded
          iconOnly
          class="song-button"
          :icon="controlIcon"
        ></base-button>
      </div>
      <div class="player-button">
        <base-button
          size="lg"
          v-on:click="skipCurrentSong()"
          rounded
          iconOnly
          class="song-button"
          icon="fa fa-step-forward"
        ></base-button>
      </div>
    </div>
  </card>
</template>

<script>
import { VideoService } from "../services/VideoService";

export default {
  data() {
    return {
      playerVars: {
        autoplay: 1,
        modestbranding: 1,
        showinfo: 0,
        controls: 0,
        disablekb: 0,
        fs: 0,
        rel: 0
      },
      isPlaying: false,
      videoService: new VideoService()
    };
  },
  props: ["videoId", "skipSong"],
  methods: {
    skipCurrentSong() {
      this.skipSong();
    },
    togglePlayer() {
      if (!this.isPlaying) this.playVideo();
      else this.pauseVideo();
    },
    pauseVideo() {
      this.player.pauseVideo();
    },
    playVideo() {
      this.player.playVideo();
    },
    playing() {
      if (!this.isPlaying) this.isPlaying = true;
    },
    ready() {},
    ended() {
      this.skipSong();
    },
    paused() {
      this.isPlaying = false;
    }
  },
  computed: {
    player() {
      return this.$refs.youtube.player;
    },
    controlIcon() {
      return this.isPlaying ? "fa fa-pause" : "fa fa-play";
    }
  }
};
</script>

<style lang="scss">
.player-parent {
  width: 100%;
  height: 100%;
}

.floating-control-container {
  position: absolute;
  top: 80%;
  right: 50%;
  left: 50%;

  display: flex;
  justify-content: space-around;
  align-items: center;

  .player-button {
    margin: 10px;

    .btn-icon-only {
      width: 75px;
      height: 75px;
      font-size: 35px;
    }
  }
}

#player-container {
  padding: 0 !important;
  background-color: transparent !important;
  .card-body {
    padding: 0 !important;
    width: 100%;
    height: 100%;
    border-radius: 8px;
    overflow: hidden;
    border: 0px SOLID rgba(244, 245, 247, 0.7);
  }
}

iframe {
  width: 100%;
  height: 100%;
}
</style>
