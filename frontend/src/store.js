import { QueueService } from "./services/QueueService";

const queueService = new QueueService();

export var store = {
  state: {
    roomKey: "",
    userKey: "",
    queue: [],
    currentVideoId: undefined,
    currentSong: undefined,
    isClient: false
  },
  setIsClient(bool) {
    this.state.isClient = bool;
  },
  setRoomKey(newRoomKey) {
    this.state.roomKey = newRoomKey;
  },
  setUserKey(newUserKey) {
    this.state.userKey = newUserKey;
  },
  hasRoomKey() {
    return this.state.roomKey.includes("ROOM_KEY:");
  },
  addSongToQueue(song) {
    if (!this.state.isClient) {
      if (this.state.currentVideoId === undefined) {
        this.state.currentVideoId = song.youtubeId;
        this.state.currentSong = song;
      } else this.state.queue.push(song);
    } else {
      if (this.state.currentVideoId === undefined)
        this.state.currentVideoId = "client";
      else this.state.queue.push(song);
    }
  },
  nextSong() {
    let song = this.state.queue[0];
    if (song !== undefined) {
      this.state.currentVideoId = song.youtubeId;
      queueService.deleteSong(song, this.state.roomKey);
    } else this.state.currentVideoId = undefined;
    this.state.currentSong = song;
  },
  getSongById(id) {
    for (let i = 0; i < this.state.queue.length; i++) {
      if (this.state.queue[i].id === id) {
        return this.state.queue[i];
      }
    }
  },
  removeFromQueue(id) {
    for (let i = 0; i < this.state.queue.length; i++) {
      if (this.state.queue[i].id === id) {
        this.state.queue.splice(i, 1);
      }
    }
  },

  reset() {
    this.state.roomKey = "";
    this.state.userKey = "";
    this.state.queue = [];
    this.state.currentVideoId = undefined;
    this.state.currentSong = undefined;
  }
};
