import config from "./config";

export function url(url) {
  return `${config.SERVER_HOST}${url}`;
}
