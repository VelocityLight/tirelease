import Table from "@mui/material/Table";
import TableContainer from "@mui/material/TableContainer";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { IssueRow } from "./IssueRow";

import AllColumns from "./ColumnDefs";
import { sampleData } from "./SampleData";

export const IssueTable = ({
  data = sampleData,
  onlyVersion,
  columns = [
    AllColumns.Repo,
    AllColumns.Issue,
    AllColumns.Title,
    AllColumns.Created,
    AllColumns.Severity,
    AllColumns.Assignee,
    AllColumns.LinkedPR,
    {
      ...AllColumns.Affects,
      columns: [AllColumns.Affects.columns[0], AllColumns.Affects.columns[1]],
    },
  ],
}) => {
  console.log(data, onlyVersion, columns);
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
            {data.map((row) => (
              <IssueRow
                key={row.Number}
                row={row}
                onlyVersion={onlyVersion}
                columns={columns}
              />
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
};
