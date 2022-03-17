import * as React from "react";
import Container from "@mui/material/Container";
import Layout from "../layout/Layout";

import Tab from "@mui/material/Tab";
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Link,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Tabs from "@mui/material/Tabs";
import Box from "@mui/material/Box";

import { useQuery } from "react-query";
import { url } from "../utils";
import { IssueGrid } from "../components/issues/IssueGrid";
import Columns from "../components/issues/GridColumns";
import {
  affectState,
  repo,
  severity,
  state,
  hasPR,
  noPR,
} from "../components/issues/filter/index";

function OpenedToday() {
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
  return (
    <IssueGrid
      data={data.data}
      columns={[
        Columns.repo,
        Columns.number,
        Columns.title,
        Columns.state,
        Columns.pr,
        Columns.type,
        Columns.severity,
        Columns.labels,
        Columns.getAffectionOnVersion("5.4"),
        Columns.getPROnVersion("5.4"),
        Columns.getPickOnVersion("5.4"),
        Columns.getAffectionOnVersion("5.3"),
        Columns.getPROnVersion("5.3"),
        Columns.getPickOnVersion("5.3"),
        Columns.getAffectionOnVersion("5.2"),
        Columns.getPROnVersion("5.2"),
        Columns.getPickOnVersion("5.2"),
        Columns.getAffectionOnVersion("5.1"),
        Columns.getPROnVersion("5.1"),
        Columns.getPickOnVersion("5.1"),
        Columns.getAffectionOnVersion("5.0"),
        Columns.getPROnVersion("5.0"),
        Columns.getPickOnVersion("5.0"),
        Columns.getAffectionOnVersion("4.0"),
        Columns.getPROnVersion("4.0"),
        Columns.getPickOnVersion("4.0"),
      ]}
      filters={[
        repo("tidb"),
        state("closed"),
        // severity("critical"),
        affectState("5.3", "yes"),
        hasPR("master"),
        hasPR("release-5.3"),
      ]}
    ></IssueGrid>
  );
}

const RecentOpen = () => {
  const [value, setValue] = React.useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  return (
    <Layout>
      <Container maxWidth="xxl" sx={{ mt: 4, mb: 4 }}>
        <Accordion defaultExpanded={true}>
          <AccordionSummary expandIcon={<ExpandMoreIcon />}>
            Recent Open
          </AccordionSummary>
          <AccordionDetails>
            <Box sx={{ width: "100%" }}>
              <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
                <Tabs
                  value={value}
                  onChange={handleChange}
                  aria-label="basic tabs example"
                >
                  <Tab label="Opened Today" />
                  <Tab label="Opened This Week" />
                  <Tab label="Opened This Month" />
                  <Tab label="All Open Issues" />
                </Tabs>
              </Box>
              <OpenedToday />
            </Box>
          </AccordionDetails>
        </Accordion>
      </Container>
    </Layout>
  );
};

export default RecentOpen;
