import * as React from "react";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import { FormHelperText } from "@mui/material";

const getStyle = (status) => {
  switch (status) {
    case "Accept":
      return { color: "green", fontWeight: "bold" };
    case "Won't Fix":
      return { color: "red", fontWeight: "bold" };
    case "Released":
      return { color: "green", fontWeight: "" };
    case "Later":
      return { color: "orange", fontWeight: "bold" };
    default:
      return {};
  }
};

const ReleaseSelector = ({ releaseProp, onChange }) => {
  const [release, setRelease] = React.useState(
    releaseProp.TriageStatus || "unknown"
  );

  const handleChange = (event) => {
    setRelease(event.target.value);
    onChange(event.target.value);
  };

  const items = ["Accept", "Won't Fix", "Later", "Released"];

  return (
    <>
      <FormControl variant="standard" sx={{ m: 0, minWidth: 120 }}>
        {releaseProp.Patch && (
          <FormHelperText>
            For {releaseProp.BaseVersion}.{releaseProp.Patch}
          </FormHelperText>
        )}
        <Select
          value={release}
          onChange={handleChange}
          displayEmpty
          inputProps={{ "aria-label": "Without label" }}
          sx={getStyle(release)}
        >
          {items.map((item) => (
            <MenuItem value={item}>{item}</MenuItem>
          ))}
        </Select>
      </FormControl>
    </>
  );
};

export default ReleaseSelector;
