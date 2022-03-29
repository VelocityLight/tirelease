import { Chip } from "@mui/material";
import { GithubIcon } from "../../icons/github.js";

function filter(params, branch) {
  return params.row.PullRequests?.filter((pr) => pr.base_branch === branch)[0];
}

export function getPullRequest(branch) {
  return (params) => {
    const pr = filter(params, branch);
    if (pr === undefined) {
      return "Not Found";
    }
    return "#" + pr.number;
  };
}

export function renderPullRequest(branch) {
  return (params) => {
    const pr = filter(params, branch);
    if (pr === undefined) {
      return <>Not Found</>;
    }
    const merged = pr.state === "merged";
    return (
      <Chip
        icon={<GithubIcon type={merged ? "merged" : "pull"} fill="white" />}
        label={"#" + pr.number}
        onClick={() => {
          window.open(pr.html_url);
        }}
        size="small"
        style={{
          color: "white",
          backgroundColor: merged ? "#8250df" : "#2da44e",
        }}
      ></Chip>
    );
  };
}
