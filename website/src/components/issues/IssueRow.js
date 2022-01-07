import Link from "@mui/material/Link";
import { TableRow, TableCell } from "@mui/material";

import SeveritySelector from "./SeveritySelector";
import Affects from "./Affects";

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
      <TableCell>{row.Assignee}</TableCell>
      <TableCell> None </TableCell>
      <TableCell>
        <Affects affectsProp={row.Affects}></Affects>
      </TableCell>
    </TableRow>
  );
};
