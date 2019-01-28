<template>
  <card type="secondary" shadow class="border-0" id="queue-container">
    <div class="scroll-container">
      <card
        type="secondary"
        v-for="song in songs"
        :key="song.id + Math.random()*2000"
        shadow
        class="border-0"
        id="queue-song"
      >
        <img
          :src="getCover(song)"
          alt="Rounded image"
          class="img-fluid img-queue rounded shadow"
          style="width: 100px; float: left"
        >
        <div class="track-text-queue">
          <h3>
            <strong>{{song.name.substring(0, 20)}}</strong>
          </h3>
          <h5>{{ song.artist[0].name}}</h5>
          <h2 id="time-text">
            <strong>{{ parseToMinutes(song.duration) }}</strong>
          </h2>
        </div>

        <i class="ni ni-fat-remove" v-on:click="handleSongRemove(song.id)" id="remove-from-queue"></i>
      </card>
    </div>
  </card>
</template>

<script>
import { QueueService } from "../services/QueueService";
import { store } from "../store";
import { VideoService } from "../services/VideoService";

const songDurationBase = 60000;
export default {
  data() {
    return {
      queueService: new QueueService(),
      videoService: new VideoService(),
      songs: store.state.queue
    };
  },
  props: ["addSongToQueue"],
  methods: {
    parseToMinutes(duration) {
      return parseFloat(Math.round(duration) / songDurationBase).toFixed(2);
    },
    getCover(song) {
      if (song != null) {
        return song.album.cover["640"];
      } else return "";
    },
    handleSongRemove(id) {
      let currentSong = store.getSongById(id);
      this.queueService.deleteSong(currentSong, store.state.roomKey);
    }
  },
  mounted: () => {
    //Horizontal scrolling with mousewheel
    const el = document.querySelector("#queue-container");

    function scrollHorizontally(e) {
      e = window.event || e;
      e.preventDefault();
      el.scrollLeft -= e.wheelDelta || -e.detail;
    }

    if (!el) {
      return;
    }

    if (el.addEventListener) {
      el.addEventListener("mousewheel", scrollHorizontally, false);
      el.addEventListener("DOMMouseScroll", scrollHorizontally, false);
    } else {
      el.attachEvent("onmousewheel", scrollHorizontally);
    }
  },
  created() {
    this.queueService.getSongs(store.state.roomKey).then(({ data }) => {
      data.forEach(song => {
        this.addSongToQueue(song);
      });
    });

    this.$connect();
    this.$socket.onopen = () => {
      this.$socket.sendObj({
        RoomKey: store.state.roomKey,
        UserKey: store.state.UserKey
      });
    };

    this.$options.sockets.onmessage = ({ data }) => {
      let parsed = JSON.parse(data);

      let song = parsed.Track;
      let type = parsed.Type;

      switch (type) {
        case "ADD_SONG": {
          this.videoService
            .getKaraokeSong(song.name, song.artist[0].name)
            .then(id => {
              song["youtubeId"] = id;
              this.addSongToQueue(song);
            })
            .catch(() => {
              this.$notify({
                title: "Karaoke song error",
                text: "We encountered an error with this song.",
                type: "warn"
              });
            });
          break;
        }
        case "REMOVE_SONG": {
          store.removeFromQueue(song.id);
          break;
        }
      }
    };
  },
  beforeDestroy() {
    this.$disconnect();
    delete this.$options.sockets.onmessage;
  }
};
</script>


<style>
.scroll-container {
  width: 999999999px; /* side scroller */
}
#queue-container {
  background-color: rgba(244, 245, 247, 0.7) !important;
  width: 100%;
  height: 100%;
  overflow-x: scroll;
  overflow-y: hidden;
  /* min-height: 124px; */
}
#queue-song {
  margin-right: 10px;
  height: 100%;
  float: left;
  width: 400px;
  overflow: hidden;
}

#queue-container .card-body {
  height: 100%;
}
#queue-song .card-body {
  height: 100%;
}

#queue-container .card-body {
  padding: 12px;
}

.track-text-queue {
  display: flex;
  flex-direction: column;
}

#remove-from-queue {
  position: absolute;
  bottom: 0;
  right: 0;
  font-size: 40px;
  color: rgba(82, 95, 127, 0.2);
}
#remove-from-queue:hover {
  color: rgb(82, 95, 127);
}

.img-queue {
  width: 100px !important;
}

.track-text-queue {
  padding-left: 8px;
}

.track-text-queue h3 {
  font-size: 24px !important;
  overflow: hidden;
  text-overflow: ellipsis;
  height: 55px;
}
.track-text-queue h5 {
  font-size: 20px !important;
}

#queue-song {
  display: flex;
  width: 300px;
}

#time-text {
  font-size: 18px;
}
@media (max-height: 775px) and (max-width: 1250px) {
  #queue-container .card-body {
    padding: 4px;
  }

  .img-queue {
    width: 100px !important;
  }

  .track-text-queue {
    padding-left: 8px;
  }

  .track-text-queue h3 {
    font-size: 20px !important;
    height: 50px;
  }
  .track-text-queue h5 {
    font-size: 18px !important;
  }

  #queue-song {
    display: flex;
    width: 300px;
  }

  #time-text {
    font-size: 14px;
  }
}

@media (max-width: 420px) and (max-height: 850px) {
  .img-queue {
    width: 75px !important;
  }

  .track-text-queue {
    padding-left: 4px;
  }

  .track-text-queue h3 {
    font-size: 18px !important;
    height: 44px;
  }
  .track-text-queue h5 {
    font-size: 16px !important;
  }

  #queue-song {
    display: flex;
    width: 250px;
  }

  #time-text {
    font-size: 14px;
  }
}

@media (max-width: 420px) and (max-height: 575px) {
  .img-queue {
    width: 75px !important;
  }

  .track-text-queue {
    padding-left: 4px;
  }

  .track-text-queue h3 {
    font-size: 18px !important;
  }
  .track-text-queue h5 {
    font-size: 16px !important;
  }

  #queue-song {
    display: flex;
    width: 250px;
  }

  #time-text {
    font-size: 14px;
  }
}

@media (max-width: 500px) and (max-height: 650px) {
  #queue-container .card-body {
    padding: 4px;
  }

  #queue-container {
    height: 100%;
  }
}

@media (min-width: 750px) and (max-height: 1250px) and (max-width: 775px) {
  #queue-container .card-body {
    padding: 8px;
  }

  #queue-container {
    height: 100%;
  }

  .img-queue {
    width: 100px !important;
  }

  .track-text-queue {
    padding-left: 8px;
  }

  .track-text-queue h3 {
    font-size: 20px !important;
    height: 125px;
  }

  .track-text-queue h5 {
    font-size: 18px !important;
  }

  #time-text {
    font-size: 16px;
  }

  #queue-song {
    margin-right: 10px;
    height: 100%;
    float: left;
    display: flex;
    width: 300px;
    overflow: hidden;
  }
}

@media (min-width: 750px) and (max-height: 775px) and (max-width: 775px) {
  #queue-container .card-body {
    padding: 8px;
  }

  #queue-container {
    height: 100%;
  }

  .img-queue {
    width: 75px !important;
  }

  .track-text-queue {
    padding-left: 8px;
  }

  .track-text-queue h3 {
    font-size: 16px !important;
    height: 40px;
  }

  .track-text-queue h5 {
    font-size: 16px !important;
  }

  #time-text {
    font-size: 14px;
  }

  #queue-song {
    margin-right: 10px;
    height: 100%;
    float: left;
    display: flex;
    width: 300px;
    overflow: hidden;
  }
}

@media (min-width: 1000px) and (max-height: 1250px) and (max-width: 1050px) {
  #queue-container .card-body {
    padding: 12px;
  }

  #queue-container {
    height: 100%;
  }

  .img-queue {
    width: 100px !important;
  }

  .track-text-queue {
    padding-left: 8px;
  }

  .track-text-queue h3 {
    font-size: 20px !important;
    height: 56px;
  }

  .track-text-queue h5 {
    font-size: 18px !important;
  }

  #time-text {
    font-size: 16px;
  }

  #queue-song {
    margin-right: 10px;
    height: 100%;
    float: left;
    display: flex;
    width: 300px;
    overflow: hidden;
  }
}
</style>


