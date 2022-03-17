import * as React from "react";
import { Link } from "react-router-dom";

import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import ListSubheader from "@mui/material/ListSubheader";
import AccountTreeIcon from "@mui/icons-material/AccountTree";
import BugReportIcon from "@mui/icons-material/BugReport";
import ColorizeIcon from "@mui/icons-material/Colorize";
import AdUnitsIcon from "@mui/icons-material/AdUnits";

export const mainListItems = (
  <div>
    <ListSubheader inset>Data Center</ListSubheader>
    <ListItem button component={Link} to="/home/open">
      <ListItemIcon>
        <AdUnitsIcon />
      </ListItemIcon>
      <ListItemText primary="All Issues" />
    </ListItem>
    <ListItem button component={Link} to="/home/open">
      <ListItemIcon>
        <AdUnitsIcon />
      </ListItemIcon>
      <ListItemText primary="Open Issues" />
    </ListItem>
    <ListItem button component={Link} to="/home/affection">
      <ListItemIcon>
        <AdUnitsIcon />
      </ListItemIcon>
      <ListItemText primary="Affection Triage" />
    </ListItem>
    <ListItem button component={Link} to="/home/cherrypick">
      <ListItemIcon>
        <AdUnitsIcon />
      </ListItemIcon>
      <ListItemText primary="Closed Issues" />
    </ListItem>
  </div>
);

export const secondaryListItems = (
  <div>
    <ListSubheader inset>Release Management</ListSubheader>

    <ListItem button component={Link} to="/home/version">
      <ListItemIcon>
        <AccountTreeIcon />
      </ListItemIcon>
      <ListItemText primary="Version" />
    </ListItem>
    <ListItem button component={Link} to="/home/triage">
      <ListItemIcon>
        <ColorizeIcon />
      </ListItemIcon>
      <ListItemText primary="Version Triage" />
    </ListItem>
  </div>
);

export const thirdListItems = (
  <div>
    <ListSubheader inset>CI/CD Tools</ListSubheader>
    <ListItem button component={Link} to="/home/tifc">
      <ListItemIcon>
        <BugReportIcon />
      </ListItemIcon>
      <ListItemText primary="TiFailureChaser" />
    </ListItem>
  </div>
);

// Icons From：https://mui.com/components/material-icons/?query=project
