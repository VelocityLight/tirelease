import * as React from "react";
import Container from "@mui/material/Container";
import Layout from "../layout/Layout";

import { Accordion, AccordionDetails, AccordionSummary } from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Box from "@mui/material/Box";

import { useQuery, useQueryClient } from "react-query";
import { IssueGrid } from "../components/issues/IssueGrid";
import Columns from "../components/issues/GridColumns";
import { fetchVersion } from "../components/issues/fetcher/fetchVersion";
import { fetchIssue } from "../components/issues/fetcher/fetchIssue";

function Table() {
  const queryClient = useQueryClient();
  const [rowCount, setRowCount] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(100);
  const [currentPage, setCurrentPage] = React.useState(0);
  const issueQuery = useQuery(
    ["issue", rowsPerPage, currentPage],
    () => fetchIssue({ page: currentPage, perPage: rowsPerPage }),
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
  React.useEffect(() => {
    if (issueQuery.data?.response.total_page > currentPage) {
      queryClient.prefetchQuery(
        ["issue", rowsPerPage, currentPage + 1],
        fetchIssue
      );
    }
  });
  const versionQuery = useQuery(["version", "maintained"], fetchVersion);
  if (issueQuery.isLoading || versionQuery.isLoading) {
    return (
      <div>
        <p>Loading...</p>
      </div>
    );
  }
  if (issueQuery.error || versionQuery.error) {
    return (
      <div>
        <p>Error: {issueQuery.error || versionQuery.error}</p>
      </div>
    );
  }

  const columns = [
    Columns.repo,
    Columns.number,
    Columns.title,
    Columns.state,
    Columns.pr,
    Columns.type,
    Columns.severity,
    Columns.labels,
  ];
  for (const version of versionQuery.data) {
    columns.push(
      Columns.getAffectionOnVersion(version),
      Columns.getPROnVersion(version),
      Columns.getPickOnVersion(version)
    );
  }
  return (
    <IssueGrid
      data={issueQuery.data.data}
      columns={columns}
      paginationMode={"server"}
      rowCount={rowCount}
      onPageChange={(page, details) => {
        console.log(page, details);
        setCurrentPage(page);
      }}
      onPageSizeChange={(pageSize, details) => {
        setRowsPerPage(pageSize);
      }}
      pageSize={rowsPerPage}
      page={currentPage}
      filters={[]}
    ></IssueGrid>
  );
}

const AllIssues = () => {
  return (
    <Layout>
      <Container maxWidth="xxl" sx={{ mt: 4, mb: 4 }}>
        <Accordion defaultExpanded={true}>
          <AccordionSummary expandIcon={<ExpandMoreIcon />}>
            All Issues
          </AccordionSummary>
          <AccordionDetails>
            <Box sx={{ width: "100%" }}>
              <Table></Table>
            </Box>
          </AccordionDetails>
        </Accordion>
      </Container>
    </Layout>
  );
};

export default AllIssues;
