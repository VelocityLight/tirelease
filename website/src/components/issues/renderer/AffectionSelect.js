import * as React from "react";
import FormControl from "@mui/material/FormControl";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import { useMutation } from "react-query";
import axios from "axios";
import { url } from "../../../utils";

export default function AffectionSelect({
  id,
  version = "master",
  affection = "UnKnown",
  onChange = () => {},
}) {
  const mutation = useMutation((newAffect) => {
    return axios.patch(url(`issue/${id}/affect/${version}`), newAffect);
  });
  const [affects, setAffects] = React.useState(affection || "UnKnown");

  const handleChange = (event) => {
    mutation.mutate({
      issue_id: id,
      affect_version: version,
      affect_result: event.target.value,
    });
    onChange(event.target.value);
    setAffects(event.target.value);
  };

  return (
    <>
      {mutation.isLoading ? (
        <p>Updating...</p>
      ) : (
        <>
          {mutation.isError ? (
            <div>An error occurred: {JSON.stringify(mutation.error)}</div>
          ) : null}
          <FormControl variant="standard" sx={{ m: 1, minWidth: 120 }}>
            <Select
              id="demo-simple-select-standard"
              value={affects}
              onChange={handleChange}
              label="Affection"
            >
              <MenuItem value={"-"} disabled={true}>-</MenuItem>
              <MenuItem value={"UnKnown"}>UnKnown</MenuItem>
              <MenuItem value={"No"}>No</MenuItem>
              <MenuItem value={"Yes"}>Yes</MenuItem>
            </Select>
          </FormControl>
        </>
      )}
    </>
  );
}
