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
        {version !== "none" && <IssueTable onlyVersion={version}></IssueTable>}
      </Stack>
    </>
  );
};

export default ReleaseTable;
