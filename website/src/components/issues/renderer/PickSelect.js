import * as React from "react";
import FormControl from "@mui/material/FormControl";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import { useMutation } from "react-query";
import axios from "axios";
import { url } from "../../../utils";

export default function PickSelect({
  id,
  version = "master",
  patch = "master",
  pick = "unknown",
  onChange = () => {},
}) {
  const mutation = useMutation((newAffect) => {
    return axios.patch(url(`issue/${id}/cherrypick/${version}`), newAffect);
  });
  const [affects, setAffects] = React.useState(pick);

  const handleChange = (event) => {
    mutation.mutate({
      issue_id: id,
      version_name: version,
      triage_result: {
        unknown: "UnKnown",
        approve: "Accept",
        later: "Later",
        "won't fix": "Won't Fix",
      }[event.target.value],
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
            <div>An error occurred: {mutation.error.message}</div>
          ) : null}
          <FormControl variant="standard" sx={{ m: 1, minWidth: 120 }}>
            <Select
              id="demo-simple-select-standard"
              value={affects}
              onChange={handleChange}
              label="Affection"
              disabled={pick.startsWith("fixed")}
            >
              <MenuItem value="unknown">
                <em>unknown</em>
              </MenuItem>
              <MenuItem value={"approve"}>approve</MenuItem>
              <MenuItem value={"later"}>later</MenuItem>
              <MenuItem value={"won't fix"}>won't fix</MenuItem>
              <MenuItem value={"fixed"} disabled={true}>
                fixed in {patch}
              </MenuItem>
            </Select>
          </FormControl>
        </>
      )}
    </>
  );
}
