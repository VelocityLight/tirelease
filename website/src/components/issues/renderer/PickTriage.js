import PickSelect from "./PickSelect";
import Button from "@mui/material/Button";
import { getAffection } from "./Affection";

export function getPickTriageValue(version) {
  return (params) => {
    const affection = getAffection(version)(params);
    if (affection === "N/A" || affection === "no") {
      return "N/A";
    }
    const pick = params.row.version_triages?.filter((t) =>
      t.version_name.startsWith(version)
    )[0];
    if (pick === undefined) {
      return "unknown";
    }
    return pick.triage_result?.toLocaleLowerCase();
  };
}

export function renderPickTriage(version) {
  return (params) => {
    const affection = getAffection(version)(params);
    if (affection === "N/A" || affection === "no") {
      return <>not affect</>;
    }
    let pick = params.row.version_triages?.filter((t) =>
      t.version_name.startsWith(version)
    )[0];
    if (pick === undefined && params.row.version_triage !== undefined) {
      pick = params.row.version_triage;
    }
    console.log(pick);
    const value =
      pick === undefined ? "unknown" : pick.triage_result?.toLocaleLowerCase();
    const patch = pick === undefined ? "unknown" : pick.version_name;
    return (
      <>
        <PickSelect
          id={params.row.issue.issue_id}
          version={version}
          patch={patch}
          pick={value}
        ></PickSelect>
        <Button>Add Note</Button>
      </>
    );
  };
}
