import * as React from "react";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";

const VersionSelector = ({ versionProp, onChange }) => {
  const [version, setVersion] = React.useState(versionProp || "none");

  const handleChange = (event) => {
    setVersion(event.target.value);
    onChange(event.target.value);
  };

  const items = ["5.4", "5.3", "5.2", "5.1", "5.0", "4.0"];

  return (
    <>
      <FormControl variant="standard" sx={{ m: 0, minWidth: 120 }}>
        <Select
          value={version}
          onChange={handleChange}
          displayEsmpty
          inputProps={{ "aria-label": "Without label" }}
        >
          <MenuItem value="none">
            <em>none</em>
          </MenuItem>
          {items.map((item) => (
            <MenuItem value={item}>{item}</MenuItem>
          ))}
        </Select>
      </FormControl>
    </>
  );
};

export default VersionSelector;
