import Link from "@mui/material/Link";
import { Accordion, AccordionDetails, AccordionSummary } from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Paper from "@mui/material/Paper";
import {
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  TableContainer,
} from "@mui/material";

import SeveritySelector from "./SeveritySelector";
import AffectsSelector from "./AffectsSelector";

export const IssueRow = ({ row }) => {
  return (
    <TableRow
      key={row.id}
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
        <Accordion>
          <AccordionSummary expandIcon={<ExpandMoreIcon />}>
            Affects 5.1, 5.2, 5.3, unkown for 4.0
          </AccordionSummary>
          <AccordionDetails>
            <TableContainer component={Paper}>
              <Table sx={{ minWidth: 650 }} size="small">
                <TableHead>
                  <TableRow>
                    <TableCell>Version</TableCell>
                    <TableCell>Affects</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {["master", "5.4", "5.3", "5.2", "5.1", "5.0", "4.0"].map(
                    (version) => (
                      <TableRow
                        key={version}
                        sx={{
                          "&:last-child td, &:last-child th": {
                            border: 0,
                          },
                        }}
                      >
                        <TableCell>{version}</TableCell>
                        <TableCell>
                          <AffectsSelector version={version}></AffectsSelector>
                        </TableCell>
                      </TableRow>
                    )
                  )}
                </TableBody>
              </Table>
            </TableContainer>
          </AccordionDetails>
        </Accordion>
      </TableCell>
    </TableRow>
  );
};
