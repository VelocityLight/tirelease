import Button from "@mui/material/Button";
import {
  Table,
  TableHead,
  TableContainer,
  TableRow,
  TableCell,
  TableBody,
  Paper,
  Stack,
  Chip,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import AffectsSelector from "./AffectsSelector";
import { useState } from "react";
import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import AccordionDetails from "@mui/material/AccordionDetails";

export default function Affects({ affectsProp }) {
  const [foldNotAffect, setFoldNotAffect] = useState(true);

  const onToggleFold = () => {
    setFoldNotAffect(!foldNotAffect);
  };

  const [affects, setAffects] = useState(affectsProp);
  const unknown = affects
    .filter(({ affect }) => affect === "unknown")
    .map(({ version }) => version);
  const affected = affects
    .filter(({ affect }) => affect === "yes")
    .map(({ version }) => version);
  const notAffected = affects
    .filter(({ affect }) => affect === "no")
    .map(({ version }) => version);
  return (
    <>
      <Accordion>
        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
          <Stack direction={"row"} spacing={1}>
            {affected.map((version) => {
              console.log(version);
              return <Chip label={"" + version} color="error" />;
            })}
            {unknown.map((version) => {
              console.log(version);
              return (
                <Chip label={"" + version} variant="outlined" color="error" />
              );
            })}
            {!foldNotAffect &&
              notAffected.map((version) => {
                console.log(version);
                return (
                  <Chip
                    label={"" + version}
                    variant="outlined"
                    color="success"
                  />
                );
              })}
          </Stack>
        </AccordionSummary>
        <AccordionDetails>
          <Stack alignItems={"flex-start"} spacing={1}>
            <TableContainer component={Paper}>
              <Table sx={{ minWidth: 650 }} size="small">
                <TableHead>
                  <TableRow>
                    <TableCell>Version</TableCell>
                    <TableCell>Affects</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {affects
                    .filter((item) => {
                      if (foldNotAffect && item.affect === "no") {
                        return false;
                      }
                      return true;
                    })
                    .map((item) => {
                      return (
                        <TableRow
                          key={item.version}
                          sx={{
                            "&:last-child td, &:last-child th": {
                              border: 0,
                            },
                          }}
                        >
                          <TableCell>
                            <Chip
                              label={item.version}
                              color={item.affect !== "no" ? "error" : "success"}
                              variant={
                                item.affect !== "yes" ? "outlined" : "filled"
                              }
                            />
                          </TableCell>
                          <TableCell>
                            <AffectsSelector
                              version={item.version}
                              affectsProp={item.affect}
                              onChange={(targetValue) => {
                                setAffects([
                                  ...affects.map(({ version, affect }) => {
                                    if (version === item.version) {
                                      return { version, affect: targetValue };
                                    }
                                    return { version, affect };
                                  }),
                                ]);
                              }}
                            ></AffectsSelector>
                          </TableCell>
                        </TableRow>
                      );
                    })}
                </TableBody>
              </Table>
            </TableContainer>
            <Button size="small" variant="outlined" onClick={onToggleFold}>
              {foldNotAffect ? "Show" : "Hide"} Not Affected
            </Button>
          </Stack>
        </AccordionDetails>
      </Accordion>
    </>
  );
}
