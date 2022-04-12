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
import { useParams, useNavigate } from "react-router-dom";
import { DataGrid, GridToolbar } from "@mui/x-data-grid";
import TriageDialog from "../issues/TriageDialog";

function ReleaseCandidates({ version }) {
  const [versionTriageData, setVersionTriageData] = useState(undefined);
  const onClose = () => {
    setVersionTriageData(undefined);
  };
  const openVersionTriageDialog = (data) => {
    setVersionTriageData(data);
  };

  const { isLoading, error, data } = useQuery(`release-${version}`, () => {
    return fetch(url(`issue/cherrypick/${version}?page=1&per_page=1000`))
      .then(async (res) => {
        const data = await res.json();
        return data;
      })
      .catch((e) => {
        console.log(e);
      });
  });
  if (isLoading) {
    return <p>Loading...</p>;
  }
  if (error) {
    return <p>Error: {error.message}</p>;
  }

  if (data?.data === undefined) {
    return <p>data is wrong, maybe your version is incorrect</p>;
  }
  const rows = data.data.version_triage_infos.map((item) => {
    return {
      id: item.issue_relation_info.issue.issue_id,
      ...item.issue_relation_info,
      version_triage: item.version_triage,
      version_triage_merge_status: item.version_triage_merge_status,
    };
  });
  const minorVersion = version.split(".").slice(0, 2).join(".");

  return (
    <div style={{ height: 650, width: "100%" }}>
      <DataGrid
        rows={rows}
        columns={[
          ...Columns.issueBasicInfo,
          Columns.triageStatus,
          Columns.releaseBlock,
          Columns.getAffectionOnVersion(minorVersion),
          Columns.getPROnVersion(minorVersion),
          Columns.getPickOnVersion(minorVersion),
          Columns.comment,
        ]}
        onRowClick={(e) => {
          console.log(e);
          openVersionTriageDialog(e);
        }}
        components={{
          Toolbar: GridToolbar,
        }}
      ></DataGrid>
      <TriageDialog
        onClose={onClose}
        open={versionTriageData !== undefined}
        row={versionTriageData?.row}
        columns={versionTriageData?.columns}
      ></TriageDialog>
    </div>
  );
}

const ReleaseTable = () => {
  const navigate = useNavigate();
  const params = useParams();
  const version = params.version === undefined ? "none" : params.version;
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
              navigate(`/home/triage/${v}`, { replace: true });
            }}
          />
          {/* <Button
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
          </Dialog> */}
        </Stack>
        {version !== "none" && (
          <ReleaseCandidates version={version}></ReleaseCandidates>
        )}
      </Stack>
    </>
  );
};

export default ReleaseTable;
