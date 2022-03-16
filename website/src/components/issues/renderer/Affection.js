import { Chip } from "@mui/material";
import AffectionSelect from "./AffectionSelect";

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
    const affection = getAffection(version)(params);
    if (affection === "N/A") {
      return <>N/A</>;
    }
    return (
      <AffectionSelect
        id={params.row.Issue.issue_id}
        version={version}
        affection={affection}
      ></AffectionSelect>
    );
  };
}
