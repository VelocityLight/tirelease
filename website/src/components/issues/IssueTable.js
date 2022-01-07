import Table from "@mui/material/Table";
import TableContainer from "@mui/material/TableContainer";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { IssueRow } from "./IssueRow";

import { sampleData } from "./SampleData";

const columns = {
  Repo: {
    display: true,
    filter: null,
    selection: null,
  },
  Issue: {
    display: true,
    filter: null,
    selection: null,
  },
  Title: {
    display: true,
    filter: null,
    selection: null,
  },
  Created: {
    display: true,
    filter: null,
    selection: null,
  },
  Updated: {
    display: true,
    filter: null,
    selection: null,
  },
};

export const IssueTable = () => {
  return (
    <>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 950 }} size="small">
          <TableHead>
            <TableRow>
              <TableCell width={10}>Repo</TableCell>
              <TableCell width={10}>Issue</TableCell>
              <TableCell width={30}>Title</TableCell>
              <TableCell width={60}>Created</TableCell>
              <TableCell width={10}>Severity</TableCell>
              <TableCell width={10}>Status</TableCell>
              <TableCell width={10}>Assignee</TableCell>
              <TableCell>Linked PR</TableCell>
              <TableCell>Affects</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {sampleData.map((row) => (
              <IssueRow key={row.Number} row={row} />
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
};
