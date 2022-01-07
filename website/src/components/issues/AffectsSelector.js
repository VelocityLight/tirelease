import * as React from "react";
import Radio from "@mui/material/Radio";
import RadioGroup from "@mui/material/RadioGroup";
import FormControlLabel from "@mui/material/FormControlLabel";
import FormControl from "@mui/material/FormControl";

export default function AffectsSelector(
  { version, affectsProp, onChange } = {
    version: "master",
    affectsProp: "unknown",
    onChange: () => {},
  }
) {
  const [affects, setAffects] = React.useState(affectsProp || "unknown");

  const handleChange = (event) => {
    onChange(event.target.value);
    setAffects(event.target.value);
  };

  return (
    <FormControl component="fieldset">
      <RadioGroup
        row
        aria-label="affects"
        name="row-radio-buttons-group"
        value={affects}
        onChange={handleChange}
      >
        <FormControlLabel value="yes" control={<Radio />} label="yes" />
        <FormControlLabel value="no" control={<Radio />} label="no" />
        <FormControlLabel value="unknown" control={<Radio />} label="unknown" />
      </RadioGroup>
    </FormControl>
  );
}
