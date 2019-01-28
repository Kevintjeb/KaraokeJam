export const endpoint =
  process.env.NODE_ENV === "development"
    ? "http://www.karaokejam.localhost"
    : "https://www.karaokejam.online";
export const wsEndpoint =
  process.env.NODE_ENV === "development"
    ? "ws://www.karaokejam.localhost"
    : "wss://www.karaokejam.online";

export const youtube_api_key = "";
