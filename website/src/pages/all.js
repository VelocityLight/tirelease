import * as React from "react";
import Container from "@mui/material/Container";
import Layout from "../layout/Layout";

import { Accordion, AccordionDetails, AccordionSummary } from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Box from "@mui/material/Box";

import { useQuery } from "react-query";
import { url, currentVersions } from "../utils";
import { IssueGrid } from "../components/issues/IssueGrid";
import Columns from "../components/issues/GridColumns";

function Table() {
  const { isLoading, error, data } = useQuery("issue", () => {
    return fetch(url("issue"))
      .then((res) => {
        const data = res.json();
        return data;
      })
      .catch((e) => {
        console.log(e);
      });
  });
  console.log(isLoading, error, data);
  if (isLoading) {
    return (
      <div>
        <p>Loading...</p>
      </div>
    );
  }
  if (error) {
    return (
      <div>
        <p>Error: {error.message}</p>
      </div>
    );
  }
  console.log("fetched data", data);
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
  for (const version of currentVersions) {
    columns.push(
      Columns.getAffectionOnVersion(version),
      Columns.getPROnVersion(version),
      Columns.getPickOnVersion(version)
    );
  }
  return (
    <IssueGrid data={data.data} columns={columns} filters={[]}></IssueGrid>
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
