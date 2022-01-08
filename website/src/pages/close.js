import * as React from "react";
import Container from "@mui/material/Container";
import Layout from "../layout/Layout";
import { IssueTable } from "../components/issues/IssueTable";
import { useQuery } from "react-query";

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

import PropTypes from "prop-types";
import Typography from "@mui/material/Typography";
import { sampleData } from "../components/issues/SampleData";
import AllColumns from "../components/issues/ColumnDefs";

function TabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          <Typography>{children}</Typography>
        </Box>
      )}
    </div>
  );
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.number.isRequired,
  value: PropTypes.number.isRequired,
};

function a11yProps(index) {
  return {
    id: `simple-tab-${index}`,
    "aria-controls": `simple-tabpanel-${index}`,
  };
}

function ClosedToday() {
  const { isLoading, error, data } = useQuery("closedToday", () => {
    return fetch("http://172.16.5.65:30750/issue")
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
    return <p>Loading...</p>;
  }
  if (error) {
    return <p>Error: {error.message}</p>;
  }
  console.log(data);
  return (
    <IssueTable
      data={data}
      columns={[
        AllColumns.Repo,
        AllColumns.Issue,
        AllColumns.Title,
        AllColumns.ClosedAt,
        AllColumns.Assignee,
        AllColumns.Severity,
        AllColumns.ClosedBy,
        AllColumns.Affects,
      ]}
    ></IssueTable>
  );
}

const RecentClose = () => {
  const [value, setValue] = React.useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };
  return (
    <Layout>
      <Container maxWidth="xxl" sx={{ mt: 4, mb: 4 }}>
        <p>Welcome to Tissue!</p>
        <Accordion defaultExpanded={true}>
          <AccordionSummary expandIcon={<ExpandMoreIcon />}>
            Recent Close{" "}
          </AccordionSummary>
          <AccordionDetails>
            <Box sx={{ width: "100%" }}>
              <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
                <Tabs
                  value={value}
                  onChange={handleChange}
                  aria-label="basic tabs example"
                >
                  <Tab label="Closed Today" {...a11yProps(0)} />
                  <Tab label="Closed This Week" {...a11yProps(1)} />
                  <Tab label="Closed This Month" {...a11yProps(2)} />
                  <Tab label="All Closed Issues" {...a11yProps(2)} />
                </Tabs>
              </Box>
              <TabPanel value={value} index={0}>
                <ClosedToday></ClosedToday>
              </TabPanel>
              <TabPanel value={value} index={1}>
                Item Two
              </TabPanel>
              <TabPanel value={value} index={2}>
                Item Three
              </TabPanel>
              <TabPanel value={value} index={3}>
                Item Three
              </TabPanel>
            </Box>
          </AccordionDetails>
        </Accordion>
      </Container>
    </Layout>
  );
};

export default RecentClose;
