import { DataGrid } from "@mui/x-data-grid";
import Columns from "./GridColumns";
import TriageDialog from "./TriageDialog";
import { useState } from "react";

export function IssueGrid(props) {
  const [triage, setTriage] = useState(false);
  const onClose = () => {
    setTriage(false);
  };
  const openTriageDialog = () => {
    setTriage(true);
  };
  return (
    <div style={{ height: 600, width: "100%" }}>
      <DataGrid
        columns={[
          Columns.repo,
          Columns.number,
          Columns.title,
          Columns.state,
          Columns.pr,
          Columns.type,
          Columns.labels,
          Columns.getAffectionOnVersion("5.4"),
          Columns.getPROnVersion("5.4"),
          Columns.getAffectionOnVersion("5.3"),
          Columns.getPROnVersion("5.3"),
          Columns.getAffectionOnVersion("5.2"),
          Columns.getPROnVersion("5.2"),
          Columns.getAffectionOnVersion("5.1"),
          Columns.getPROnVersion("5.1"),
          Columns.getAffectionOnVersion("5.0"),
          Columns.getPROnVersion("5.0"),
          Columns.getAffectionOnVersion("4.0"),
          Columns.getPROnVersion("4.0"),
        ]}
        rows={[
          ...props.data.map((item) => {
            return { ...item, id: item.Issue.issue_id };
          }),
        ]}
        onRowClick={(e) => {
          // openTriageDialog();
        }}
      ></DataGrid>
      <TriageDialog onClose={onClose} open={triage}></TriageDialog>
    </div>
  );
}
