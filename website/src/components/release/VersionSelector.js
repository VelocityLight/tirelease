import * as React from "react";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import { useQuery } from "react-query";
import { url } from "../../utils";

const VersionSelector = ({ versionProp, onChange }) => {
  const [version, setVersion] = React.useState(versionProp || "none");
  const { isLoading, error, data } = useQuery("versions", () => {
    return fetch(url("version")).then((res) => {
      console.log(res);
      return res.json();
    });
  });

  if (isLoading) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p>Error: {error.message}</p>;
  }

  const versions = data.data.map((version) => version.name);

  const handleChange = (event) => {
    setVersion(event.target.value);
    onChange(event.target.value);
  };

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
          {versions.map((item) => (
            <MenuItem value={item}>{item}</MenuItem>
          ))}
        </Select>
      </FormControl>
    </>
  );
};

export default VersionSelector;
