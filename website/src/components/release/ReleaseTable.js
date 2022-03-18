import { Stack } from "@mui/material";
import VersionSelector from "./VersionSelector";
import { useState } from "react";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import { useQuery } from "react-query";
import { url } from "../../utils";
import { IssueGrid } from "../issues/IssueGrid";
import Columns from "../issues/GridColumns";
import { OR } from "../issues/filter";
import { affectUnknown, affectYes, pick } from "../issues/filter/index";

function ReleaseCandidates({ version }) {
  const { isLoading, error, data } = useQuery(`release-${version}`, () => {
    return fetch(url(`issue/cherrypick/${version}`))
      .then((res) => {
        const data = res.json();
        return data;
      })
      .catch((e) => {
        console.log(e);
      });
  });
  console.log(isLoading, error, data);
  if (isLoading) {
    return <p>Loading...</p>;
  }
  if (error) {
    return <p>Error: {error.message}</p>;
  }
  const filters = [
    OR([affectUnknown(version), affectYes(version)]),
    pick(version, "unknown"),
  ];
  console.log("version", data);
  const rows = data.data.version_triage_infos.map(
    (item) => item.issue_relation_info
  );
  console.log("version", rows);
  return (
    <IssueGrid
      data={rows}
      columns={[
        ...Columns.issueBasicInfo,
        Columns.getAffectionOnVersion(version),
        Columns.getPROnVersion(version),
        Columns.getPickOnVersion(version),
      ]}
      filters={filters}
    ></IssueGrid>
  );
}

const ReleaseTable = () => {
  const [version, setVersion] = useState("none");
  const [open, setOpen] = useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <>
      <Stack spacing={1}>
        <Stack direction={"row"} justifyContent={"space-between"}>
          <VersionSelector
            versionProp={version}
            onChange={(v) => {
              setVersion(v);
            }}
          />
          <Button
            variant="outlined"
            onClick={handleClickOpen}
            disabled={version === "none"}
          >
            Release
          </Button>
          <Dialog
            open={open}
            onClose={handleClose}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
          >
            <DialogTitle id="alert-dialog-title">
              {"Are You Sure to Release?"}
            </DialogTitle>
            <DialogContent>
              <DialogContentText id="alert-dialog-description">
                Release {version} will create a new patch version, and this is
                inreversible, please make sure all issues that are affected by
                this release are triaged and settled.
              </DialogContentText>
            </DialogContent>
            <DialogActions>
              <Button onClick={handleClose}>Cancel</Button>
              <Button onClick={handleClose} autoFocus>
                Release
              </Button>
            </DialogActions>
          </Dialog>
        </Stack>
        {version !== "none" && (
          <ReleaseCandidates version={version}></ReleaseCandidates>
        )}
      </Stack>
    </>
  );
};

export default ReleaseTable;
