import { Stack } from "@mui/material";
import VersionSelector from "./VersionSelector";
import { useState } from "react";
import { IssueTable } from "../issues/IssueTable";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import AllColumns from "../issues/ColumnDefs";
import { useQuery } from "react-query";

function ReleaseCandidates({ version }) {
  const { isLoading, error, data } = useQuery("releaseCandidates", () => {
    return fetch("http://172.16.5.65:30750/issue")
      .then((res) => {
        const data = res.json();
        console.log(data);
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
  console.log(data);
  return (
    <IssueTable
      data={data}
      columns={[
        AllColumns.Repo,
        AllColumns.Issue,
        AllColumns.Title,
        AllColumns.ClosedAt,
        AllColumns.Assignee,
        AllColumns.Severity,
        AllColumns.ClosedBy,
        AllColumns.Affects,
      ]}
      onlyVersion={version}
    ></IssueTable>
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
