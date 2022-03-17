import { DataGrid, GridToolbar } from "@mui/x-data-grid";
import Columns from "./GridColumns";
import TriageDialog from "./TriageDialog";
import { useState } from "react";

export function IssueGrid({
  data,
  filters = [],
  columns = [Columns.number, Columns.title],
}) {
  const [triage, setTriage] = useState(false);
  const onClose = () => {
    setTriage(false);
  };
  const openTriageDialog = () => {
    setTriage(true);
  };
  return (
    <>
      <div style={{ height: 600, width: "100%" }}>
        <DataGrid
          density="compact"
          columns={columns}
          rows={[
            ...data
              .map((item) => {
                return { ...item, id: item.Issue.issue_id };
              })
              .filter((item) => {
                for (const filter of filters) {
                  if (!filter(item)) {
                    return false;
                  }
                }
                return true;
              }),
          ]}
          onRowClick={(e) => {
            // openTriageDialog();
          }}
          components={{ Toolbar: GridToolbar }}
        ></DataGrid>
        <TriageDialog onClose={onClose} open={triage}></TriageDialog>
      </div>
    </>
  );
}
