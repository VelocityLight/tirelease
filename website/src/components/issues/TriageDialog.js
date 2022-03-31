import * as React from "react";
import PropTypes from "prop-types";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogActions from "@mui/material/DialogActions";
import IconButton from "@mui/material/IconButton";
import CloseIcon from "@mui/icons-material/Close";
import { Paper, Table, TableBody, TableCell, TableRow } from "@mui/material";

const BootstrapDialogTitle = (props) => {
  const { children, onClose, ...other } = props;

  return (
    <DialogTitle sx={{ m: 0, p: 2 }} {...other}>
      {children}
      {onClose ? (
        <IconButton
          aria-label="close"
          onClick={onClose}
          sx={{
            position: "absolute",
            right: 8,
            top: 8,
            color: (theme) => theme.palette.grey[500],
          }}
        >
          <CloseIcon />
        </IconButton>
      ) : null}
    </DialogTitle>
  );
};

BootstrapDialogTitle.propTypes = {
  children: PropTypes.node,
  onClose: PropTypes.func.isRequired,
};

export default function TriageDialog({ row, columns, open, onClose }) {
  return (
    <div>
      <Dialog
        onClose={onClose}
        open={open}
        sx={{ overflow: "visible" }}
        maxWidth={"md"}
        fullWidth
        scroll="paper"
      >
        <BootstrapDialogTitle id="customized-dialog-title" onClose={onClose}>
          Issue info
        </BootstrapDialogTitle>
        <Table>
          <TableBody>
            {columns?.map((column) => {
              return (
                <TableRow>
                  <TableCell>{column.headerName}</TableCell>
                  <TableCell>
                    {(() => {
                      if (column.renderCell) {
                        return column.renderCell({ row });
                      }
                      return column.valueGetter({ row });
                    })()}
                  </TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
        <DialogContent dividers></DialogContent>
        <DialogActions>
          <Button autoFocus onClick={onClose}>
            Close
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}
