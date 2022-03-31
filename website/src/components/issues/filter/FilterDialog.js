import { Dialog } from "@mui/material";
import TiDialogTitle from "../../common/TiDialogTitle";

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

export function FilterDialog({ open, onClose, filters }) {
  return (
    <Dialog onClose={onClose} open={open} maxWidth="md" fullWidth>
      <TiDialogTitle onClose={onClose}>Filter Panel</TiDialogTitle>
    </Dialog>
  );
}
