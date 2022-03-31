class Filter {
  getQuery() {}
  render() {}
}

class SeverityOption {
  major = "severity/major";
  critical = "severity/critical";
  moderate = "severity/moderate";
  minor = "severity/minor";
}

class SeverityFilter extends Filter {
  selectedSeverities = [SeverityOption.major, SeverityOption.critical];
  constructor() {}
}

export function FilterDialog({ open, onClose, filters }) {}
