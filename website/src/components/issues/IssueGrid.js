import { DataGrid, GridToolbar } from "@mui/x-data-grid";
import Columns from "./GridColumns";
import TriageDialog from "./TriageDialog";
import { useState, useEffect } from "react";
import { useQuery, useQueryClient } from "react-query";
import { fetchIssue } from "./fetcher/fetchIssue";

export function IssueGrid({
  filters = [],
  columns = [Columns.number, Columns.title],
}) {
  const queryClient = useQueryClient();
  const [rowCount, setRowCount] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(100);
  const [currentPage, setCurrentPage] = useState(0);
  const issueQuery = useQuery(
    ["issue", ...filters, rowsPerPage, currentPage],
    () => fetchIssue({ filters, page: currentPage, perPage: rowsPerPage }),
    {
      onSuccess: (data) => {
        console.log("setRowCount", rowCount, data.response.total_count);
        setRowCount(data.response.total_count);
      },
      keepPreviousData: true,
      staleTime: 5000,
    }
  );
  // prefetch next page
  useEffect(() => {
    if (issueQuery.data?.response.total_page > currentPage) {
      queryClient.prefetchQuery(
        ["issue", ...filters, rowsPerPage, currentPage + 1],
        () =>
          fetchIssue({ filters, page: currentPage + 1, perPage: rowsPerPage })
      );
    }
  });
  const [triage, setTriage] = useState(false);
  const onClose = () => {
    setTriage(false);
  };
  const openTriageDialog = () => {
    setTriage(true);
  };

  if (issueQuery.isLoading) {
    return (
      <div>
        <p>Loading...</p>
      </div>
    );
  }
  if (issueQuery.isError) {
    return (
      <div>
        <p>error: {issueQuery.error}</p>
      </div>
    );
  }

  const rows = [
    ...issueQuery.data?.data
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
          paginationMode={"server"}
          rowCount={rowCount}
          page={currentPage}
          pageSize={rowsPerPage}
          onPageChange={(page, details) => {
            console.log(page, details);
            setCurrentPage(page);
          }}
          onPageSizeChange={(pageSize, details) => {
            setRowsPerPage(pageSize);
          }}
        ></DataGrid>
        <TriageDialog onClose={onClose} open={triage}></TriageDialog>
      </div>
    </>
  );
}
