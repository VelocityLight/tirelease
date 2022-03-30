import { url } from "../../../utils";

export function fetchIssue({ page = 0, perPage = 100, filters = [] }) {
  return fetch(
    url(
      `issue?page=${page}&per_page=${perPage}${filters
        .map((filter) => "&" + filter)
        .join("")}`
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
