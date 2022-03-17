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
  severity,
  state,
  type,
  openIn24h,
  openSince,
  NOT,
  AND,
} from "../components/issues/filter/index";

function RecentOpen() {
  const filters = [
    type("bug"),
    AND([NOT(severity("moderate")), NOT(severity("minor"))]),
    state("open"),
  ];
  const pickColumns = [];
  for (const version of currentVersions) {
    pickColumns.push(Columns.getAffectionOnVersion(version));
  }

  const [tab, setTab] = React.useState(0);
  const tabs = ["Created in 24h", "Created in 7d", "Created in 30d", "ALL"];
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
            Recent Open
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

export default RecentOpen;
