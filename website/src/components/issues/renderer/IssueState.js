import { Chip } from "@mui/material";

export function renderIssueState(params) {
  const state = params.row.Issue.state;
  return <Chip label={state}></Chip>;
}
