import axios from "axios";
import { roomsEndpoint } from "../constants";

export class RoomService {
  joinRoom(roomkey) {
    return axios.post(`${roomsEndpoint}/join`, { RoomKey: roomkey });
  }

  leaveRoom(userkey) {
    return axios.post(`${roomsEndpoint}/leave`, { UserKey: userkey });
  }

  newRoom() {
    return axios.get(`${roomsEndpoint}/new`);
  }
}
