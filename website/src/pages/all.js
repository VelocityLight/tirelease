import * as React from "react";
import Container from "@mui/material/Container";
import Layout from "../layout/Layout";

import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Chip,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Box from "@mui/material/Box";

import { useQuery } from "react-query";
import { IssueGrid } from "../components/issues/IssueGrid";
import Columns from "../components/issues/GridColumns";
import { fetchVersion } from "../components/issues/fetcher/fetchVersion";
import { fetchIssue } from "../components/issues/fetcher/fetchIssue";

function Table() {
  const issueQuery = useQuery("issue", fetchIssue);
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
