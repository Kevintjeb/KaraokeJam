<template>
  <card class="border-0" id="container">
    <div class="card-body" style="padding: 0 !important">
      <div class="title-container">
        <div class="sign-out-button-container-client">
          <base-button
            rounded
            iconOnly
            @click="$emit('signout')"
            class="sign-out-button-client"
            icon="fa fa-sign-out"
          ></base-button>
        </div>
        <div class="text-center-client text-muted mb-4" id="logo-client">KaraokeJam</div>

        <div class="qr-code-button-container-client">
          <base-button
            rounded
            iconOnly
            @click="$emit('click')"
            class="qr-code-button-client"
            icon="fa fa-qrcode"
          ></base-button>
        </div>
      </div>
      <div class="input-group mb-4" style="padding: 0 !important">
        <input
          class="form-control"
          id="searchbox"
          placeholder="Search by track or artist"
          type="text"
          v-on:keyup.enter="fetchSongs($event)"
          v-model="searchInput"
        >
        <div v-on:click="fetchSongs(searchInput)" class="input-group-append">
          <span class="input-group-text">
            <i class="ni ni-zoom-split-in"></i>
          </span>
        </div>
      </div>
      <div v-if="results !== null && results !== undefined" class="parent-card">
        <card class="song-item" v-for="result in results" :key="result.id">
          <img
            :src="getCover(result)"
            alt="Rounded image"
            class="img-fluid img-side rounded shadow"
            style="width: 100px; float: left"
          >
          <div id="track-text">
            <h3 class="song-title-client">
              <strong>{{ result.name }}</strong>
            </h3>
            <h5>{{ result.artist[0].name }}</h5>
            <h6>{{ result.year }}</h6>
          </div>
          <div class="card-button">
            <base-button
              v-on:click="requestSong(result)"
              rounded
              iconOnly
              class="song-button"
              icon="fa fa-plus"
            ></base-button>
          </div>
        </card>
      </div>
      <div v-else-if="results === undefined">
        <div class="song-message">
          <div id="track-text">
            <h3 class="song-title-client">
              <strong>Start Jamming!</strong>
            </h3>
            <h5>Please search on a track name.</h5>
          </div>
        </div>
      </div>
      <div v-else-if="results === null">
        <div class="song-message">
          <div id="track-text">
            <h3 class="song-title-client">
              <strong>There were no results!</strong>
            </h3>
            <h5>Please search on a track name.</h5>
          </div>
        </div>
      </div>
    </div>
  </card>
</template>

<script>
import { store } from "../store";
import { SearchService } from "../services/SearchService";
import { QueueService } from "../services/QueueService";
import { VideoService } from "../services/VideoService";
export default {
  data() {
    return {
      results: undefined,
      state: store.state,
      searchService: new SearchService(),
      queueService: new QueueService(),
      videoService: new VideoService()
    };
  },
  props: ["videoId", "skipSong"],
  methods: {
    getCover(song) {
      if (song != null) {
        return song.album.cover["640"];
      } else return "";
    },
    fetchSongs(event) {
      this.searchService.search(event.target.value).then(response => {
        this.results = response.data;
      });
    },
    requestSong(song) {
      this.videoService
        .getKaraokeSong(song.name, song.artist[0].name)
        .then(id => {
          song["youtubeId"] = id;
          this.queueService.addSong(song, this.state.roomKey);
        })
        .catch(() => {
          this.$notify({
            title: "Karaoke song error",
            text: "We encountered an error with this song.",
            type: "warn"
          });
        });
    }
  }
};
</script>

<style lang="scss">
@import "../assets/scss/definitions.scss";

.title-container {
  display: flex;
  justify-content: space-between;
}

.song-message {
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.parent-card {
  overflow-y: scroll;
  height: calc(100% - 220px);
}
::-webkit-scrollbar {
  height: 12px;
  width: 12px;
  -webkit-border-radius: 1ex;
}

::-webkit-scrollbar-thumb {
  background: #fff;
  -webkit-border-radius: 1ex;
}

::-webkit-scrollbar-corner {
  background: #fff;
}
.song-item {
  margin: 10px 5px 0px 0px;
  height: 175px;
  overflow: hidden;
}

#logo-client {
  font-family: HarlowSolid;
  color: #fff;
  text-align: center;
  font-size: 32pt;
}

#container {
  background-color: rgba(244, 245, 247, 0.7) !important;
  height: 100%;
  width: 100%;
}

#searchbox {
  background-color: rgba(255, 255, 255, 0.3);
  color: #000;
  font-size: 16px;
  padding: 8 !important;
}
#searchbox ::placeholder {
  color: cadetblue;
}
#song {
  margin-bottom: 10px;
}
#track-text {
  padding-left: 10px;
  display: flex;
  flex-direction: column;
}

.card-button {
  align-self: flex-end;
  position: absolute;
  right: 15;
  bottom: 0;
  top: 120;
}

.song-button {
  background-color: $button-primary !important;
  border: none !important;
}

.qr-code-button-container-client {
  z-index: 999;

  .qr-code-button-client {
    background-color: $button-primary !important;
    border: none !important;
    font-size: 32;
    width: 52;
    height: 52;
  }
}

canvas {
  width: calc(150px + 10vw) !important;
  height: auto !important;
}

.sign-out-button-container-client {
  z-index: 999;

  .sign-out-button-client {
    background-color: $button-primary !important;
    border: none !important;
    font-size: 32;
    width: 52;
    height: 52;

    i {
      margin-right: 5px;
      -webkit-transform: rotate(180deg);
      -moz-transform: rotate(180deg);
      -ms-transform: rotate(180deg);
      -o-transform: rotate(180deg);
      transform: rotate(180deg);
    }
  }
}

@media (max-width: 420px) {
  .sign-out-button-container-client {
    .sign-out-button-client {
      width: 36;
      height: 36;
      font-size: 22;
    }
  }

  .qr-code-button-container-client {
    .qr-code-button-client {
      width: 36;
      height: 36;
      font-size: 22;
    }
  }

  #logo-client {
    font-size: 28;
  }

  .song-title-client {
    font-size: 22;
  }

  .img-side {
    width: 72px !important;
  }
}
</style>
