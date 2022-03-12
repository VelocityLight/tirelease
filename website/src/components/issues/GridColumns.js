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
};

const title = {
  field: "title",
  headerName: "Title",
  valueGetter: (params) => params.row.Issue.title,
};

const type = {
  field: "type",
  headerName: "Type",
  valueGetter: (params) => params.row.Issue.type,
};

const Columns = { id, repo, number, title, type };

export default Columns;
