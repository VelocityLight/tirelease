import { AccordionDetails, AccordionSummary, Container } from "@mui/material";
import React from "react";
import { Layout } from "../layout/Layout";
import Accordion from "@mui/material/Accordion";
import ExpandMore from "@mui/icons-material/ExpandMore";
import Tabs from "@mui/material/Tabs";
import Tab from "@mui/material/Tab";
import { IssueGrid } from "../components/issues/IssueGrid";
import Columns from "../components/issues/GridColumns";
import { currentVersions, url } from "../utils";
import { useQuery } from "react-query";
import {
  OR,
  severity,
  state,
  type,
  closedByPRIn24h,
  closedByPRSince,
} from "../components/issues/filter/index";

function PickTriage() {
  const filters = [
    type("bug"),
    OR([severity("critical"), severity("major")]),
    state("closed"),
  ];
  const pickColumns = [];
  for (const version of currentVersions) {
    pickColumns.push(
      Columns.getAffectionOnVersion(version),
      Columns.getPROnVersion(version),
      Columns.getPickOnVersion(version)
    );
  }

  const [tab, setTab] = React.useState(0);
  const tabs = ["Closed in 24h", "Closed in 7d", "Closed in 30d", "ALL"];
  switch (tab) {
    case 0:
      filters.push(closedByPRIn24h());
      break;
    case 1:
      filters.push(
        closedByPRSince(new Date().getTime() - 60 * 60 * 1000 * 24 * 7)
      );
      break;
    case 2:
      filters.push(
        closedByPRSince(new Date().getTime() - 60 * 60 * 1000 * 24 * 30)
      );
      break;
    case 3:
      break;
    default:
      break;
  }

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
  return (
    <Layout>
      <Container maxWidth="xxl" sx={{ mt: 4, mb: 4 }}>
        <Accordion defaultExpanded={true}>
          <AccordionSummary expandIcon={<ExpandMore />}>
            Cherrypick Triage
          </AccordionSummary>
          <AccordionDetails>
            <Tabs value={tab} onChange={handleChange}>
              {tabs.map((v) => (
                <Tab label={v}></Tab>
              ))}
            </Tabs>
            {isLoading && <p>Loading...</p>}
            {error && <p>Error: {error.message}</p>}
            {data && (
              <IssueGrid
                columns={[
                  Columns.repo,
                  Columns.number,
                  Columns.title,
                  Columns.type,
                  Columns.severity,
                  Columns.state,
                  Columns.pr,
                  ...pickColumns,
                ]}
                data={data.data}
                filters={filters}
              ></IssueGrid>
            )}
          </AccordionDetails>
        </Accordion>
      </Container>
    </Layout>
  );
}

export default PickTriage;
