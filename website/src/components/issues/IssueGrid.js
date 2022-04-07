import { DataGrid, GridToolbar } from "@mui/x-data-grid";
import Columns from "./GridColumns";
import TriageDialog from "./TriageDialog";
import { useState, useEffect } from "react";
import { useQuery, useQueryClient } from "react-query";
import { fetchIssue } from "./fetcher/fetchIssue";
import { Button, Stack } from "@mui/material";
import { FilterDialog } from "./filter/FilterDialog";

export function IssueGrid({
  filters = [],
  columns = [Columns.number, Columns.title],
}) {
  const queryClient = useQueryClient();
  const [filterDialog, setFilterDialog] = useState(false);
  const [rowCount, setRowCount] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(100);
  const [currentPage, setCurrentPage] = useState(1);
  // stale while revalidate
  // js spread operator
  // const a = [1, 2, 3];
  // [0, ...a, 4] == [0, 1, 2, 3, 4];
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
  const [triageData, setTriageData] = useState(undefined);
  const onClose = () => {
    setTriageData(undefined);
  };
  const openTriageDialog = (data) => {
    setTriageData(data);
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
    ...issueQuery.data?.data.map((item) => {
      return { ...item, id: item.issue.issue_id };
    }),
  ];
  return (
    <Stack spacing={1}>
      <Stack direction={"row"} justifyContent={"flex-end"} spacing={2}>
        <Button
          variant="contained"
          onClick={() => {
            setFilterDialog(true);
          }}
        >
          Filter
        </Button>
      </Stack>
      <div style={{ height: 600, width: "100%" }}>
        <DataGrid
          density="compact"
          columns={columns}
          rows={rows}
          onRowClick={(e) => {
            console.log(e);
            openTriageDialog(e);
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
        <TriageDialog
          onClose={onClose}
          open={triageData !== undefined}
          row={triageData?.row}
          columns={triageData?.columns}
        ></TriageDialog>
        <FilterDialog
          open={filterDialog}
          onClose={() => {
            setFilterDialog(false);
          }}
        ></FilterDialog>
      </div>
    </Stack>
  );
}
