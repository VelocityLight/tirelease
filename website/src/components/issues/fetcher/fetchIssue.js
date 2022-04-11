import { url } from "../../../utils";

export function fetchIssue({ page = 1, perPage = 100, filters = [] }) {
  const queryString = `${filters
    .map((filter) => {
      const transform = filter.stringify || ((filter) => filter);
      const param = transform(filter);
      if (param.length > 0) {
        return "&" + param;
      }
      return "";
    })
    .join("")}`;
  console.log(
    "fetchIssue",
    url(`issue?page=${page}&per_page=${perPage}${queryString}`)
  );
  return fetch(url(`issue?page=${page}&per_page=${perPage}${queryString}`))
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
    .catch((e) => {
      console.log(e);
    });
}
