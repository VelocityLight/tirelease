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
import { url, currentVersions } from "../utils";
import Columns from "../components/issues/GridColumns";
import {
  repo,
  state,
  affectState,
  severity,
  hasPR,
  OR,
} from "../components/issues/filter/index";

const AffectTriage = () => {
  const [tab, setTab] = React.useState(0);

  const handleChange = (event, newValue) => {
    setTab(newValue);
  };

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

  const affectColumns = [];
  if (tab === 0) {
    affectColumns.push(...currentVersions.map(Columns.getAffectionOnVersion));
  } else {
    affectColumns.push(Columns.getAffectionOnVersion(currentVersions[tab - 1]));
  }

  const filters = [OR([severity("critical"), severity("major")])];
  if (tab === 0) {
    filters.push(
      OR(currentVersions.map((version) => affectState(version, "unknown")))
    );
  } else {
    filters.push(affectState(currentVersions[tab - 1], "unknown"));
  }

  return (
    <Layout>
      <Container maxWidth="xxl" sx={{ mt: 4, mb: 4 }}>
        <Accordion defaultExpanded={true}>
          <AccordionSummary expandIcon={<ExpandMoreIcon />}>
            Affection Triage
          </AccordionSummary>
          <AccordionDetails>
            <Box sx={{ width: "100%" }}>
              <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
                <Tabs
                  value={tab}
                  onChange={handleChange}
                  aria-label="basic tabs example"
                >
                  <Tab label="All" />
                  {currentVersions.map((v) => (
                    <Tab label={v}></Tab>
                  ))}
                </Tabs>
              </Box>
              {isLoading && <p>Loading...</p>}
              {error && <p>Error: {error.message}</p>}
              {data && (
                <IssueGrid
                  data={data.data}
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
              )}
            </Box>
          </AccordionDetails>
        </Accordion>
      </Container>
    </Layout>
  );
};

export default AffectTriage;
