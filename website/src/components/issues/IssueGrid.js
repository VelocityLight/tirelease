import { DataGrid, GridToolbar } from "@mui/x-data-grid";
import Columns from "./GridColumns";
import TriageDialog from "./TriageDialog";
import { useState } from "react";

export function IssueGrid({
  data,
  filters = [],
  columns = [Columns.number, Columns.title],
  paginationMode = "client",
  rowCount = -1,
  page = 0,
  pageSize = 100,
  onPageChange = (page, details) => {},
  onPageSizeChange = (pageSize, details) => {},
}) {
  console.log("IssueGrid", paginationMode, rowCount);
  const [triage, setTriage] = useState(false);
  const onClose = () => {
    setTriage(false);
  };
  const openTriageDialog = () => {
    setTriage(true);
  };
  const rows = [
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
  ];
  if (rowCount === -1) {
    rowCount = rows.length;
  }
  return (
    <>
      <div style={{ height: 600, width: "100%" }}>
        <DataGrid
          density="compact"
          columns={columns}
          rows={rows}
          onRowClick={(e) => {
            // openTriageDialog();
          }}
          components={{ Toolbar: GridToolbar }}
          paginationMode={paginationMode}
          rowCount={rowCount}
          page={page}
          pageSize={pageSize}
          onPageChange={onPageChange}
          onPageSizeChange={onPageSizeChange}
        ></DataGrid>
        <TriageDialog onClose={onClose} open={triage}></TriageDialog>
      </div>
    </>
  );
}
