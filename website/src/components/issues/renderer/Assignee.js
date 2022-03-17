import { Chip } from "@mui/material";

export function renderAssignee(params) {
  return params.row.Issue.assignee.map((assignee) => (
    <Chip label={assignee.login}></Chip>
  ));
}
