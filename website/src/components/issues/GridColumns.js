import { renderIssueState } from "./renderer/IssueState";
import { renderAssignee } from "./renderer/Assignee";
import { getAffection, renderAffection } from "./renderer/Affection";
import { getPullRequest, renderPullRequest } from "./renderer/PullRequest";
import { getLabelValue, renderLabel } from "./renderer/Label";
import { getPickTriageValue, renderPickTriage } from "./renderer/PickTriage";

const id = {
  field: "id",
  headerName: "Id",
  hide: true,
  valueGetter: (params) => params.row.Issue.issue_id,
};

const repo = {
  field: "repo",
  headerName: "Repo",
  valueGetter: (params) => params.row.Issue.repo,
};

const number = {
  field: "number",
  headerName: "Number",
  valueGetter: (params) => params.row.Issue.number,
  renderCell: (params) => (
    <a href={params.row.Issue.html_url}>{params.row.Issue.number}</a>
  ),
};

const title = {
  field: "title",
  headerName: "Title",
  width: 480,
  valueGetter: (params) => params.row.Issue.title,
};

const type = {
  field: "type",
  headerName: "Type",
  width: 120,
  valueGetter: getLabelValue(
    (label) => label.name.startsWith("type/"),
    (label) => label.replace("type/", "")
  ),
  renderCell: renderLabel(
    (label) => label.name.startsWith("type/"),
    (label) => label.replace("type/", "")
  ),
};

const severity = {
  field: "severity",
  headerName: "Severity",
  width: 120,
  valueGetter: getLabelValue(
    (label) => label.name.startsWith("severity/"),
    (label) => label.replace("severity/", "")
  ),
  renderCell: renderLabel(
    (label) => label.name.startsWith("severity/"),
    (label) => label.replace("severity/", "")
  ),
};

const state = {
  field: "state",
  headerName: "State",
  valueGetter: (params) => params.row.Issue.state,
  renderCell: renderIssueState,
};

const assignee = {
  field: "assignee",
  headerName: "Assignee",
  valueGetter: (params) =>
    params.row.Issue.assignee.map((assignee) => assignee.login).join(","),
  renderCell: renderAssignee,
};

const labelFilter = (label) =>
  !label.name.startsWith("type/") &&
  !label.name.startsWith("severity/") &&
  !label.name.startsWith("affects-") &&
  !label.name.startsWith("may-affect-");

const labels = {
  field: "labels",
  headerName: "Labels",
  valueGetter: getLabelValue(labelFilter, (label) => label),
  renderCell: renderLabel(labelFilter, (label) => label),
};

const pr = {
  field: "pr",
  headerName: "PR",
  valueGetter: getPullRequest("master"),
  renderCell: renderPullRequest("master"),
};

function getAffectionOnVersion(version) {
  return {
    field: "affect_" + version,
    headerName: "Affect " + version,
    valueGetter: getAffection(version),
    renderCell: renderAffection(version),
  };
}

function getPROnVersion(version) {
  const branch = "release-" + version;
  return {
    field: "cherrypick_" + version,
    headerName: "PR for " + version,
    valueGetter: getPullRequest(branch),
    renderCell: renderPullRequest(branch),
  };
}

function getPickOnVersion(version) {
  return {
    field: "pick_" + version,
    headerName: "Pick to " + version,
    width: 240,
    valueGetter: getPickTriageValue(version),
    renderCell: renderPickTriage(version),
  };
}

const Columns = {
  id,
  repo,
  number,
  title,
  state,
  type,
  labels,
  assignee,
  severity,
  pr,
  getAffectionOnVersion,
  getPROnVersion,
  getPickOnVersion,
  issueBasicInfo: [id, repo, number, title],
};

export default Columns;
