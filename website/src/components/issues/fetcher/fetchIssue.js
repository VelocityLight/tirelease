import { url } from "../../../utils";

export function fetchIssue({ page = 0, perPage = 100, state = undefined }) {
  return fetch(
    url(
      `issue?page=${page}&per_page=${perPage}${state ? `&state=${state}` : ""}`
    )
  )
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
    .catch((e) => {
      console.log(e);
    });
}
