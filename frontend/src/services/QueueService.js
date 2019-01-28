import axios from "axios";
import { queueEndpoint } from "../constants";

export class QueueService {
  getSongs(roomkey) {
    return axios.get(`${queueEndpoint}/${roomkey}`);
  }

  addSong(song, roomkey) {
    return axios.post(`${queueEndpoint}`, {
      RoomKey: roomkey,
      Song: song,
      UserKey: ""
    });
  }

  deleteSong(song, roomkey) {
    return axios.post(`${queueEndpoint}/delete`, {
      RoomKey: roomkey,
      Song: song
    });
  }
}
