import { DataGrid } from "@mui/x-data-grid";
import Columns from "./GridColumns";

const data = {
  Issue: {
    id: 1,
    issue_id: "I_kwDOGf5ce85E2TaX",
    number: 15,
    state: "OPEN",
    title: "Workflow & API requirements",
    owner: "VelocityLight",
    repo: "tirelease",
    html_url: "https://github.com/VelocityLight/tirelease/issues/15",
    created_at: "2022-03-01T09:22:04Z",
    updated_at: "2022-03-10T13:24:51Z",
    labels_string: '[{"name":"type/bug"}]',
    assignees_string: '[{"login":"VelocityLight"}]',
    labels: [
      {
        name: "type/bug",
      },
    ],
    assignees: [
      {
        login: "VelocityLight",
      },
    ],
  },
  IssueAffects: [
    {
      create_time: "0001-01-01T00:00:00Z",
      update_time: "0001-01-01T00:00:00Z",
      issue_id: "I_kwDOGf5ce85E2TaX",
      affect_version: "5.4",
      affect_result: "UnKnown",
    },
  ],
  IssuePrRelations: [
    {
      id: 210001,
      create_time: "2022-03-10T13:06:04Z",
      update_time: "2022-03-10T13:06:04Z",
      issue_id: "I_kwDOGf5ce85E2TaX",
      pull_request_id: "PR_kwDOGf5ce84z1Q4r",
    },
    {
      id: 210002,
      create_time: "2022-03-10T13:06:04Z",
      update_time: "2022-03-10T13:06:04Z",
      issue_id: "I_kwDOGf5ce85E2TaX",
      pull_request_id: "PR_kwDOGf5ce84z7gtv",
    },
  ],
  PullRequests: [
    {
      id: 2,
      pull_request_id: "PR_kwDOGf5ce84z1Q4r",
      number: 16,
      state: "closed",
      title: "issue_relation_info API complete",
      owner: "VelocityLight",
      repo: "tirelease",
      html_url: "https://github.com/VelocityLight/tirelease/pull/16",
      base_branch: "main",
      created_at: "2022-03-02T14:16:58Z",
      updated_at: "2022-03-10T13:47:44Z",
      closed_at: "2022-03-03T03:54:27Z",
      merged_at: "2022-03-03T03:54:27Z",
      merged: true,
      mergeable_state: "unknown",
      already_reviewed: true,
      labels_string: '[{"name":"status/LGT2"}]',
      assignees_string: "[]",
      requested_reviewers_string: "[]",
      labels: [
        {
          name: "status/LGT2",
        },
      ],
      assignees: [],
      requested_reviewers: [],
    },
    {
      id: 1,
      pull_request_id: "PR_kwDOGf5ce84z7gtv",
      number: 17,
      state: "closed",
      title: "Backend allapi",
      owner: "VelocityLight",
      repo: "tirelease",
      html_url: "https://github.com/VelocityLight/tirelease/pull/17",
      base_branch: "main",
      created_at: "2022-03-04T01:30:32Z",
      updated_at: "2022-03-11T02:06:38Z",
      closed_at: "2022-03-11T02:06:35Z",
      merged_at: "2022-03-11T02:06:35Z",
      merged: true,
      mergeable_state: "unknown",
      cherry_pick_approved: true,
      already_reviewed: true,
      labels_string: '[{"name":"cherry-pick-approved"},{"name":"status/LGT2"}]',
      assignees_string: "[]",
      requested_reviewers_string: "[]",
      labels: [
        {
          name: "cherry-pick-approved",
        },
        {
          name: "status/LGT2",
        },
      ],
      assignees: [],
      requested_reviewers: [],
    },
  ],
};

export function IssueGrid() {
  return (
    <DataGrid
      columns={[Columns.title, Columns.repo]}
      rows={[
        { id: 0, ...data },
        { id: 1, ...data },
      ]}
    ></DataGrid>
  );
}
