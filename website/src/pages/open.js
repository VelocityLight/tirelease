import * as React from "react";
import Container from "@mui/material/Container";
import Layout from "../layout/Layout";
import { IssueTable } from "../components/issues/IssueTable";

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

import { sampleData } from "../components/issues/SampleData";
import { useQuery } from "react-query";
import AllColumns from "../components/issues/ColumnDefs";
import TabPanel from "../components/issues/TablePanel";
import { url } from "../utils";
import { IssueGrid } from "../components/issues/IssueGrid";

function a11yProps(index) {
  return {
    id: `simple-tab-${index}`,
    "aria-controls": `simple-tabpanel-${index}`,
  };
}

function OpenedToday() {
  /*
  const { isLoading, error, data } = useQuery("openedToday", () => {
    return fetch(url("issue?state=open"))
      .then((res) => {
        const data = res.json();
        console.log(data);
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
  */
  return (
    <IssueGrid data={sampleData.data}></IssueGrid>
    // <IssueTable
    //   data={data.data}
    //   columns={[
    //     AllColumns.Repo,
    //     AllColumns.Issue,
    //     AllColumns.Title,
    //     AllColumns.Created,
    //     AllColumns.Severity,
    //     AllColumns.Assignee,
    //     AllColumns.LinkedPR,
    //     // {
    //     //   ...AllColumns.Affects,
    //     //   columns: [...AllColumns.Affects.columns],
    //     // },
    //   ]}
    // ></IssueTable>
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
                  <Tab label="Opened Today" {...a11yProps(0)} />
                  <Tab label="Opened This Week" {...a11yProps(1)} />
                  <Tab label="Opened This Month" {...a11yProps(2)} />
                  <Tab label="All Open Issues" {...a11yProps(2)} />
                </Tabs>
              </Box>
              <TabPanel value={value} index={0}>
                <OpenedToday />
              </TabPanel>
              <TabPanel value={value} index={1}>
                Item Two
              </TabPanel>
              <TabPanel value={value} index={2}>
                Item Three
              </TabPanel>
              <TabPanel value={value} index={3}>
                Item Four
              </TabPanel>
            </Box>
          </AccordionDetails>
        </Accordion>
      </Container>
    </Layout>
  );
};

export default RecentOpen;
