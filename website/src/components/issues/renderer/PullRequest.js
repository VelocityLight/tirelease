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
    console.log(pr);
    const iconType =
      pr.state === "closed" ? (pr.merged ? "merged" : "pr_closed") : pr.state;
    const iconColor = {
      pr_closed: "#cf222e",
      merged: "#8250df",
      open: "#2da44e",
    }[iconType];

    return (
      <Chip
        icon={<GithubIcon type={iconType} fill="white" />}
        label={"#" + pr.number}
        onClick={() => {
          window.open(pr.html_url);
        }}
        size="small"
        style={{
          color: "white",
          backgroundColor: iconColor,
        }}
      ></Chip>
    );
  };
}
