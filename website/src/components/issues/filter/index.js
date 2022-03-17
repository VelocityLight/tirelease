export function repo(name) {
  return (item) => item.Issue.repo === name;
}

export function state(status) {
  return (item) => item.Issue.state === status;
}

export function stateOpen() {
  return state("open");
}

export function severity(level) {
  return (item) =>
    item.Issue.labels.filter((label) => label.name === `severity/${level}`)
      .length !== 0;
}

export function type(issueType) {
  return (item) =>
    item.Issue.labels.filter((label) => label.name === `type/${issueType}`)
      .length !== 0;
}

export function affectState(version, state) {
  return (item) => {
    const affect = item.IssueAffects.filter(
      (affects) => affects.affect_version === version
    )[0];
    if (affect === undefined) {
      return false;
    }
    return (
      affect.affect_result.toLocaleLowerCase() === state.toLocaleLowerCase()
    );
  };
}

export function affectYes(version) {
  return affectState(version, "yes");
}

export function affectUnknown(version) {
  return affectState(version, "unknown");
}

export function PR(branch, has) {
  return (item) => {
    const pr = item.PullRequests.filter((pr) => pr.base_branch === branch)[0];
    return (pr === undefined) !== has;
  };
}

export function hasPR(branch) {
  return PR(branch, true);
}

export function noPR(branch) {
  return PR(branch, false);
}

export function pick(version, state) {
  return (item) => {
    const pick = item.VersionTriages.filter((t) =>
      t.version_name.startsWith(version)
    )[0];
    if (pick === undefined && state === "unknown") {
      return true;
    }
    return pick.triage_result.toLocaleLowerCase() === state.toLocaleLowerCase();
  };
}

export function OR(filters) {
  return (item) => {
    for (const filter of filters) {
      if (filter(item)) {
        return true;
      }
    }
    return false;
  };
}

export function AND(filters) {
  return (item) => {
    for (const filter of filters) {
      if (!filter(item)) {
        return false;
      }
    }
    return true;
  };
}

export function NOT(filter) {
  return (item) => !filter(item);
}

export function closedByPRDuring(from, to) {
  const f = new Date(from);
  const t = new Date(to);
  return (item) => {
    const pr = item.PullRequests.filter((pr) => pr.base_branch === "master")[0];
    if (pr === undefined || pr.merged_at === undefined) {
      return false;
    }
    const mergedAt = new Date(pr.merged_at);
    return mergedAt - f >= 0 && t - mergedAt > 0;
  };
}

export function closedByPRSince(from) {
  return closedByPRDuring(from, new Date());
}

export function closedByPRIn24h() {
  return closedByPRSince(new Date(new Date().getTime() - 24 * 60 * 60 * 1000));
}

export function openDuring(from, to) {
  const f = new Date(from);
  const t = new Date(to);
  return (item) => {
    const createdAt = new Date(item.Issue.created_at);
    return createdAt - f >= 0 && t - createdAt > 0;
  };
}

export function openSince(from) {
  return openDuring(from, new Date());
}

export function openIn24h() {
  return openSince(new Date(new Date().getTime() - 24 * 60 * 60 * 1000));
}
