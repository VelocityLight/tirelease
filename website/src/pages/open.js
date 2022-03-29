import { AccordionDetails, AccordionSummary, Container } from "@mui/material";
import React from "react";
import { Layout } from "../layout/Layout";
import Accordion from "@mui/material/Accordion";
import ExpandMore from "@mui/icons-material/ExpandMore";
import Tabs from "@mui/material/Tabs";
import Tab from "@mui/material/Tab";
import { IssueGrid } from "../components/issues/IssueGrid";
import Columns from "../components/issues/GridColumns";
import { useQuery } from "react-query";
import { fetchVersion } from "../components/issues/fetcher/fetchVersion";
import { fetchIssue } from "../components/issues/fetcher/fetchIssue";
import {
  severity,
  state,
  type,
  openIn24h,
  openSince,
  NOT,
  AND,
} from "../components/issues/filter/index";

const Table = ({ tab }) => {
  const versionQuery = useQuery(["version", "maintained"], fetchVersion);
  const issueQuery = useQuery("issue", fetchIssue);
  if (issueQuery.isLoading || versionQuery.isLoading) {
    return (
      <div>
        <p>Loading...</p>
      </div>
    );
  }
  if (issueQuery.isError || versionQuery.isError) {
    return (
      <div>
        <p>{issueQuery.error || versionQuery.error}</p>
      </div>
    );
  }
  const filters = [
    type("bug"),
    AND([NOT(severity("moderate")), NOT(severity("minor"))]),
    state("open"),
  ];
  const pickColumns = [];
  for (const version of versionQuery.data) {
    pickColumns.push(Columns.getAffectionOnVersion(version));
  }
  switch (tab) {
    case 0:
      filters.push(openIn24h());
      break;
    case 1:
      filters.push(openSince(new Date().getTime() - 60 * 60 * 1000 * 24 * 7));
      break;
    case 2:
      filters.push(openSince(new Date().getTime() - 60 * 60 * 1000 * 24 * 30));
      break;
    case 3:
      break;
    default:
      break;
  }
  return (
    <IssueGrid
      columns={[
        Columns.repo,
        Columns.number,
        Columns.title,
        Columns.type,
        Columns.severity,
        Columns.state,
        ...pickColumns,
      ]}
      data={issueQuery.data.data}
      filters={filters}
    ></IssueGrid>
  );
};

function RecentOpen() {
  const [tab, setTab] = React.useState(0);
  const tabs = ["Created in 24h", "Created in 7d", "Created in 30d", "ALL"];
  const handleChange = (event, newValue) => {
    setTab(newValue);
  };
  return (
    <Layout>
      <Container maxWidth="xxl" sx={{ mt: 4, mb: 4 }}>
        <Accordion defaultExpanded={true}>
          <AccordionSummary expandIcon={<ExpandMore />}>
            Recent Open
          </AccordionSummary>
          <AccordionDetails>
            <Tabs value={tab} onChange={handleChange}>
              {tabs.map((v) => (
                <Tab label={v}></Tab>
              ))}
            </Tabs>
            <Table></Table>
          </AccordionDetails>
        </Accordion>
      </Container>
    </Layout>
  );
}

export default RecentOpen;
