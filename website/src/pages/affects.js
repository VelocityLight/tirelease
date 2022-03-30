import { Layout } from "../layout/Layout";
import Container from "@mui/material/Container";
import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import AccordionDetails from "@mui/material/AccordionDetails";
import Box from "@mui/material/Box";
import Tabs from "@mui/material/Tabs";
import Tab from "@mui/material/Tab";
import React from "react";
import { IssueGrid } from "../components/issues/IssueGrid";
import { useQuery } from "react-query";
import Columns from "../components/issues/GridColumns";
import { fetchVersion } from "../components/issues/fetcher/fetchVersion";

const VersionTabs = () => {
  const [tab, setTab] = React.useState(0);

  const handleChange = (event, newValue) => {
    setTab(newValue);
  };

  const versionQuery = useQuery(["version", "maintained"], fetchVersion);
  if (versionQuery.isLoading) {
    return (
      <div>
        <p>Loading...</p>
      </div>
    );
  }
  if (versionQuery.isError) {
    return (
      <div>
        <p>{versionQuery.error}</p>
      </div>
    );
  }
  const affectColumns = [];
  const currentVersions = versionQuery.data;
  if (tab === 0) {
    affectColumns.push(...currentVersions.map(Columns.getAffectionOnVersion));
  } else {
    affectColumns.push(Columns.getAffectionOnVersion(currentVersions[tab - 1]));
  }

  const filters = [
    // "type_label=type/bug",
    // "state=open",
    "severity_labels=severity/major",
    "severity_labels=severity/critical",
    // OR([severity("critical"), severity("major")]),
  ];
  // if (tab === 0) {
  //   filters.push(
  //     OR(currentVersions.map((version) => affectState(version, "unknown")))
  //   );
  // } else {
  //   filters.push(affectState(currentVersions[tab - 1], "unknown"));
  // }

  return (
    <>
      <Tabs value={tab} onChange={handleChange} aria-label="basic tabs example">
        <Tab label="All" />
        {currentVersions.map((v) => (
          <Tab label={v}></Tab>
        ))}
      </Tabs>
      <IssueGrid
        columns={[
          Columns.repo,
          Columns.number,
          Columns.title,
          Columns.state,
          Columns.type,
          Columns.severity,
          ...affectColumns,
        ]}
        filters={filters}
      ></IssueGrid>
    </>
  );
};

const AffectTriage = () => {
  return (
    <Layout>
      <Container maxWidth="xxl" sx={{ mt: 4, mb: 4 }}>
        <Accordion defaultExpanded={true}>
          <AccordionSummary expandIcon={<ExpandMoreIcon />}>
            Affection Triage
          </AccordionSummary>
          <AccordionDetails>
            <Box sx={{ width: "100%" }}>
              <Box sx={{ borderBottom: 1, borderColor: "divider" }}></Box>
              <VersionTabs></VersionTabs>
            </Box>
          </AccordionDetails>
        </Accordion>
      </Container>
    </Layout>
  );
};

export default AffectTriage;
