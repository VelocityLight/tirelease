import { Chip } from "@mui/material";
import AffectionSelect from "./AffectionSelect";

export function getAffection(version) {
  return (params) => {
    const affects = params.row.issue_affects?.filter(
      (affects) => affects.affect_version === version
    )[0];
    if (affects === undefined) {
      return "-";
    }
    return affects.affect_result;
  };
}

export function renderAffection(version) {
  return (params) => {
    const affection = getAffection(version)(params);
    return (
      <AffectionSelect
        id={params.row.issue.issue_id}
        version={version}
        affection={affection}
      ></AffectionSelect>
    );
  };
}
