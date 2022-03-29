import { url } from "../../../utils";

export function fetchIssue() {
  return fetch(url("issue"))
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
    .catch((e) => {
      console.log(e);
    });
}
