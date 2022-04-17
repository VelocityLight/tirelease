import PickSelect from "./PickSelect";
import { getAffection } from "./Affection";

export function getPickTriageValue(version) {
  return (params) => {
    const affection = getAffection(version)(params);
    if (affection === "-" || affection === "no") {
      return <>not affect</>;
    }
    const version_triage = params.row.version_triages?.filter((t) =>
      t.version_name.startsWith(version)
    )[0];
    if (version_triage === undefined) {
      return "-"
    }
    return version_triage.triage_result;
  };
}

export function renderPickTriage(version) {
  return (params) => {
    
    const affection = getAffection(version)(params);
    if (affection === "-" || affection === "no") {
      return <>not affect</>;
    }
    let version_triage = params.row.version_triages?.filter((t) =>
      t.version_name.startsWith(version)
    )[0];
    const pick = version_triage === undefined ? "-" : version_triage.triage_result;
    const patch = version_triage === undefined ? "-" : version_triage.version_name;

    return (
      <>
        <PickSelect
          id={params.row.issue.issue_id}
          version={version}
          patch={patch}
          pick={pick}
        ></PickSelect>
      </>
    );
  };
}
