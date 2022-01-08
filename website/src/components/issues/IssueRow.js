import Link from "@mui/material/Link";
import { TableRow, TableCell } from "@mui/material";

import SeveritySelector from "./SeveritySelector";
import Affects from "./Affects";
import ColorHash from "color-hash";
import AllColumns from "./ColumnDefs";

export const IssueRow = ({ row, onlyVersion, columns }) => {
  console.log(row, onlyVersion);
  return (
    <TableRow
      key={row.Number}
      sx={{
        "&:last-child td, &:last-child th": { border: 0 },
      }}
    >
      {columns.map((column) => {
        if (column.display) {
          switch (column.title) {
            case "Repo":
              return <TableCell>tidb</TableCell>;
            case "Issue":
              return (
                <TableCell>
                  <Link href={row.Url}>{row.Number}</Link>
                </TableCell>
              );
            case "Title":
              return <TableCell>{row.Title}</TableCell>;
            case "Created":
              return (
                <TableCell>
                  {row.CreatedAt || new Date().toDateString()}
                </TableCell>
              );
            case "Severity":
              return (
                <TableCell>
                  <SeveritySelector severityProp={row.Severity || ""} />
                </TableCell>
              );
            case "State":
              return <TableCell>{row.State}</TableCell>;
            case AllColumns.ClosedAt.title:
              return (
                <TableCell>
                  {row.ClosedAt || new Date().toDateString()}
                </TableCell>
              );
            case AllColumns.ClosedBy.title:
              return <TableCell>PR</TableCell>;
            case "Assignee":
              return (
                <TableCell>
                  <button
                    style={{
                      backgroundColor: new ColorHash().hex(row.Assignee),
                      border: 0,
                      borderRadius: "20px",
                      padding: "5px 10px",
                      maxWidth: "12em",
                      minWidth: "5em",
                    }}
                    href={"https://github.com/" + row.Assignee}
                  >
                    {row.Assignee}
                  </button>
                </TableCell>
              );
            case "Linked PR":
              return <TableCell> None </TableCell>;
            case "Affects":
              return (
                <TableCell>
                  <Affects
                    affectsProp={row.Affects}
                    onlyVersion={onlyVersion}
                    expandProp={onlyVersion !== undefined}
                    showProp={false}
                    columns={column.columns}
                  ></Affects>
                </TableCell>
              );

            default:
              return <></>;
          }
        }
        return <></>;
      })}
    </TableRow>
  );
};
