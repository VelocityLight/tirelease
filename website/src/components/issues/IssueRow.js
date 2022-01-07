import Link from "@mui/material/Link";
import { TableRow, TableCell } from "@mui/material";

import SeveritySelector from "./SeveritySelector";
import Affects from "./Affects";
import ColorHash from "color-hash";

export const IssueRow = ({ row }) => {
  return (
    <TableRow
      key={row.Number}
      sx={{
        "&:last-child td, &:last-child th": { border: 0 },
      }}
    >
      <TableCell>tidb</TableCell>

      <TableCell>
        <Link href={row.Url}>{row.Number}</Link>
      </TableCell>
      <TableCell>{row.Title}</TableCell>

      <TableCell>{row.CreatedAt || new Date().toDateString()}</TableCell>
      <TableCell>
        <SeveritySelector severityProp={row.Severity || ""} />
      </TableCell>
      <TableCell>{row.State}</TableCell>
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
      <TableCell> None </TableCell>
      <TableCell>
        <Affects affectsProp={row.Affects}></Affects>
      </TableCell>
    </TableRow>
  );
};
