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
  Link,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import AffectsSelector from "./AffectsSelector";
import { useState } from "react";
import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import AccordionDetails from "@mui/material/AccordionDetails";
import ReleaseSelector from "./ReleaseSelector";

export default function Affects({ affectsProp, expand } = { expand: false }) {
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
      <Accordion defaultExpanded={expand}>
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
              <Table sx={{ minWidth: 950 }} size="small">
                <TableHead>
                  <TableRow>
                    <TableCell>Version</TableCell>
                    <TableCell>Affects</TableCell>
                    <TableCell>PR</TableCell>
                    <TableCell>State</TableCell>
                    <TableCell>Release</TableCell>
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
                                  ...affects.map(
                                    ({ version, affect, ...rest }) => {
                                      if (version === item.version) {
                                        return {
                                          version,
                                          affect: targetValue,
                                          ...rest,
                                        };
                                      }
                                      return { version, affect, ...rest };
                                    }
                                  ),
                                ]);
                              }}
                            ></AffectsSelector>
                          </TableCell>
                          <TableCell>
                            {item.pr && (
                              <Link href={item.pr.Url}>{item.pr.Number}</Link>
                            )}
                          </TableCell>
                          <TableCell>{item.pr && item.pr.State}</TableCell>
                          <TableCell>
                            {item.Release && (
                              <ReleaseSelector
                                releaseProp={item.Release}
                                onChange={(triageStatus) => {
                                  setAffects([
                                    ...affects.map(
                                      ({ version, Release, ...rest }) => {
                                        if (version === item.version) {
                                          return {
                                            version,
                                            Release: {
                                              ...Release,
                                              TriageStatus: triageStatus,
                                            },
                                            ...rest,
                                          };
                                        }
                                        return { version, Release, ...rest };
                                      }
                                    ),
                                  ]);
                                }}
                              />
                            )}
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
