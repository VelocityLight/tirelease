import { Chip } from "@mui/material";

export function getAffection(version) {
  return (params) => {
    const affects = params.row.IssueAffects.filter(
      (affects) => affects.affect_version === version
    );
    if (affects.length === 0) {
      return "N/A";
    }
    return { Yes: "yes", No: "no", UnKnown: "unknown" }[
      affects[0].affect_result
    ];
  };
}

export function renderAffection(version) {
  return (params) => {
    return <Chip label={getAffection(version)(params)}></Chip>;
  };
}
