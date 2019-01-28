import { endpoint } from "./env";
import { youtube_api_key } from "./env";

export const roomsEndpoint = endpoint + "/rooms";
export const queueEndpoint = endpoint + "/songs";
export const searchEndpoint = endpoint + "/search";

export const videoEndpoint = "https://www.googleapis.com/youtube/v3/search?key=" + youtube_api_key;
