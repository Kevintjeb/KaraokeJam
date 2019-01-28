import axios from "axios";
import { videoEndpoint } from "../constants";

export class VideoService {
  getKaraokeSong(title, artist) {
    var q = artist + "+" + title + "+karaoke";
    q.split(" ").join("+");

    return new Promise((resolve, reject) => {
      axios
        .get(
          `${videoEndpoint}&part=snippet&maxResults=1&type=video&videoEmbeddable=true&q=${q}`
        )
        .then(function(response) {
          if (response.data.items === undefined)
            reject(new Error("No reponse body received"));

          resolve(response.data.items[0].id.videoId);
        })
        .catch(function(error) {
          reject(error);
        });
    });
  }
}
