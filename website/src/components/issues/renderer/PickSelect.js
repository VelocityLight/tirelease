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
  pick = "UnKnown",
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
      triage_result: event.target.value,
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
              disabled={pick.startsWith("Released")}
            >
              <MenuItem value={"-"} disabled={true}>-</MenuItem>
              <MenuItem value={"UnKnown"}>UnKnown</MenuItem>
              <MenuItem value={"Accept"}>
                <div style={{ color: "green", fontWeight: "bold" }}>
                  Accept
                </div>
              </MenuItem>
              <MenuItem value={"Later"}>Later</MenuItem>
              <MenuItem value={"Won't Fix"}>Won't Fix</MenuItem>
              <MenuItem value={"Accept(Frozen)"} disabled={true}>Accept(Frozen)</MenuItem>
              <MenuItem value={"Released"} disabled={true}>
                Released in {patch}
              </MenuItem>
            </Select>
          </FormControl>
        </>
      )}
    </>
  );
}
