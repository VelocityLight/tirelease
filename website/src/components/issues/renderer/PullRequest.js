import { Chip } from "@mui/material";

function filter(params, branch) {
  return params.row.PullRequests.filter((pr) => pr.base_branch === branch);
}

export function getPullRequest(branch) {
  return (params) => {
    const pr = filter(params, branch);
    if (pr.length === 0) {
      return "Not Found";
    }
    return "#" + pr[0].number;
  };
}

export function renderPullRequest(branch) {
  return (params) => {
    return <Chip label={getPullRequest(branch)(params)}></Chip>;
  };
}
