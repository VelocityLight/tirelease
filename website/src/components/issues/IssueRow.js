import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import Link from "@mui/material/Link";
import { TableRow, TableCell } from "@mui/material";

import SeveritySelector from "./SeveritySelector";
import Affects from "./Affects";
import ColorHash from "color-hash";
import AllColumns from "./ColumnDefs";

dayjs.extend(relativeTime);

export const IssueRow = ({ row, onlyVersion, columns }) => {
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
                  {dayjs(row.CreatedAt).fromNow() || "Unknown"}
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
                  {row.Assignee && (
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
                  )}
                  {row.Assignee || "Unassigned"}
                </TableCell>
              );
            case "Linked PR":
              return (
                <TableCell>
                  {row.PR && <a href={row.PR.Url}>{row.PR.Number}</a>}
                  {row.PR || "None"}
                </TableCell>
              );
            case "Affects":
              return (
                <TableCell>
                  <Affects
                    id={row.IssueID}
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
