import { useState } from "react";
import { Stack, Typography, Button, Dialog } from "@mui/material";
import TiDialogTitle from "../../common/TiDialogTitle";

function Comment({ row }) {
  const [open, setOpen] = useState(false);
  return (
    <>
      <Stack direction={"row"} justifyContent={"space-between"}>
        <Button
          variant="contained"
          onClick={(e) => {
            setOpen(true);
            e.stopPropagation();
          }}
        >
          ADD
        </Button>
        <Typography>{row.version_triage.comment}</Typography>
      </Stack>
      <Dialog
        open={open}
        onClose={() => {
          setOpen(false);
        }}
      >
        <TiDialogTitle>Add Comment</TiDialogTitle>
      </Dialog>
    </>
  );
}

export function renderComment({ row }) {
  return <Comment row={row} />;
}
