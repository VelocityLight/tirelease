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
import Columns from "../issues/GridColumns";
import { useParams } from "react-router-dom";
import { DataGrid } from "@mui/x-data-grid";

function ReleaseCandidates({ version }) {
  const { isLoading, error, data } = useQuery(`release-${version}`, () => {
    return fetch(url(`issue/cherrypick/${version}?page=0&per_page=1000`))
      .then(async (res) => {
        const data = await res.json();
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

  console.log("version data", data);
  if (data?.data === undefined) {
    return <p>data is wrong, maybe your version is incorrect</p>;
  }
  const rows = data.data.version_triage_infos.map((item) => {
    return {
      id: item.issue_relation_info.Issue.issue_id,
      ...item.issue_relation_info,
      version_triage: item.version_triage,
      version_triage_merge_status: item.version_triage_merge_status,
    };
  });
  console.log("version rows", rows);
  const minorVersion = version.split(".").slice(0, 2).join(".");
  return (
    <div style={{ height: 600, width: "100%" }}>
      <DataGrid
        rows={rows}
        columns={[
          ...Columns.issueBasicInfo,
          Columns.triageStatus,
          Columns.getAffectionOnVersion(minorVersion),
          Columns.getPROnVersion(minorVersion),
          Columns.getPickOnVersion(minorVersion),
        ]}
      ></DataGrid>
    </div>
  );
}

const ReleaseTable = () => {
  const params = useParams();
  console.log(params.version);
  const [version, setVersion] = useState(
    params.version === undefined ? "none" : params.version
  );
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
