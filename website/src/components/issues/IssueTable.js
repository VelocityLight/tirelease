import Table from "@mui/material/Table";
import TableContainer from "@mui/material/TableContainer";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { IssueRow } from "./IssueRow";

import { sampleData } from "./SampleData";

const AllColumns = {
  Repo: {
    title: "Repo",
    display: true,
    filter: null,
    selection: null,
  },
  Issue: {
    title: "Issue",
    display: true,
    filter: null,
    selection: null,
  },
  Title: {
    title: "Title",
    display: true,
    filter: null,
    selection: null,
  },
  Created: {
    title: "Created",
    display: true,
    filter: null,
    selection: null,
  },
  Updated: {
    title: "Updated",
    display: true,
    filter: null,
    selection: null,
  },
  State: {
    title: "State",
    display: true,
    filter: null,
    selection: null,
  },
  LinkedPR: {
    title: "Linked PR",
    display: true,
    filter: null,
    selection: null,
  },
  Assignee: {
    title: "Assignee",
    display: true,
    filter: null,
    selection: null,
  },
  Affects: {
    title: "Affects",
    display: true,
    filter: null,
    selection: null,
  },
  Severity: {
    title: "Severity",
    display: true,
    filter: null,
    selection: null,
  },
};

export const IssueTable = ({
  onlyVersion,
  columns = [
    AllColumns.Repo,
    AllColumns.Issue,
    AllColumns.Title,
    AllColumns.Created,
    AllColumns.Severity,
    AllColumns.State,
    AllColumns.Assignee,
    AllColumns.LinkedPR,
    AllColumns.Affects,
  ],
}) => {
  console.log(onlyVersion, columns);
  return (
    <>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 950 }} size="small">
          <TableHead>
            <TableRow>
              {columns.map((column) => {
                if (column.display) {
                  return <TableCell>{column.title}</TableCell>;
                }
                return <></>;
              })}
            </TableRow>
          </TableHead>
          <TableBody>
            {sampleData.map((row) => (
              <IssueRow key={row.Number} row={row} onlyVersion={onlyVersion} />
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
};
