import Button from "@mui/material/Button";
import {
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Stack,
  Chip,
  Link,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import AffectsSelector from "./AffectsSelector";
import { useState } from "react";
import ReleaseSelector from "./ReleaseSelector";

import * as React from "react";
import VisibilityIcon from "@mui/icons-material/Visibility";
import ToggleButton from "@mui/material/ToggleButton";
import ToggleButtonGroup from "@mui/material/ToggleButtonGroup";

function ToggleButtons({ onExpand, onShow }) {
  const [buttonState, setButtonState] = React.useState([]);

  const handleChange = (event, change) => {
    setButtonState(change);
    console.log(event, change);
    onExpand(change.includes("expand"));
    onShow(change.includes("show"));
  };

  return (
    <ToggleButtonGroup
      size="small"
      value={buttonState}
      onChange={handleChange}
      aria-label="list state"
    >
      <ToggleButton value="expand" aria-label="expand">
        <ExpandMoreIcon />
      </ToggleButton>
      <ToggleButton value="show" aria-label="show">
        <VisibilityIcon />
      </ToggleButton>
    </ToggleButtonGroup>
  );
}

export default function Affects(
  { affectsProp, expandProp, showProp, onlyVersion } = {
    expandProp: false,
    showProp: false,
  }
) {
  const [showNotAffect, setShowNotAffect] = useState(showProp);
  const [expand, setExpand] = useState(expandProp);
  const [affects, setAffects] = useState(
    onlyVersion
      ? affectsProp.filter((item) => item.version === onlyVersion)
      : affectsProp
  );

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
      <Stack direction={"row"} spacing={1} alignItems={"center"}>
        {!onlyVersion && (
          <ToggleButtons
            onExpand={(expand) => {
              setExpand(expand);
            }}
            onShow={(show) => {
              setShowNotAffect(show);
            }}
          ></ToggleButtons>
        )}
        {!expand && (
          <>
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
            {showNotAffect &&
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
          </>
        )}
      </Stack>

      {expand && (
        <Stack alignItems={"flex-start"} spacing={1}>
          {/* <TableContainer component={Paper}> */}
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
                  if (showNotAffect && item.affect === "no") {
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
                              ...affects.map(({ version, affect, ...rest }) => {
                                if (version === item.version) {
                                  return {
                                    version,
                                    affect: targetValue,
                                    ...rest,
                                  };
                                }
                                return { version, affect, ...rest };
                              }),
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
          {/* </TableContainer> */}
        </Stack>
      )}
    </>
  );
}
