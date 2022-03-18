import * as React from "react";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
import TextField from "@mui/material/TextField";
import { FormHelperText, Input, Stack } from "@mui/material";
import { useMutation, useQueryClient } from "react-query";
import { url } from "../../utils";
import axios from "axios";

export const VersionAdd = ({ open, onClose }) => {
  const queryClient = useQueryClient();
  const [version, setVersion] = React.useState("");
  const create = useMutation(
    (data) => {
      return axios.post(url("version"), data);
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("versions");
        onClose();
      },
      onError: (e) => {
        console.log("error", e);
      },
    }
  );
  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>Create Version</DialogTitle>
      <DialogContent sx={{ display: "flex", flexDirection: "column" }}>
        <FormControl>
          <InputLabel htmlFor="my-input">Email address</InputLabel>
          <Input id="my-input" aria-describedby="my-helper-text" />
          <FormHelperText id="my-helper-text">
            We'll never share your email.
          </FormHelperText>
        </FormControl>
        <DialogContentText margin={2}>Version</DialogContentText>
        <Stack direction="row" spacing={2}>
          <FormControl>
            <InputLabel id="create-version">Version</InputLabel>
            <Select
              labelId="create-version"
              id="create-version-select"
              value={version}
              label="Version"
              onChange={(e) => {
                setVersion(e.target.value);
              }}
            >
              <MenuItem value={10}>Ten</MenuItem>
              <MenuItem value={20}>Twenty</MenuItem>
              <MenuItem value={30}>Thirty</MenuItem>
            </Select>
          </FormControl>
          <FormControl>
            <TextField label="name"></TextField>
          </FormControl>
        </Stack>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Close</Button>
        <Button
          onClick={() => {
            create.mutate();
          }}
          variant="contained"
        >
          Save
        </Button>
      </DialogActions>
    </Dialog>
  );
};
