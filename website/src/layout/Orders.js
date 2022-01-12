import * as React from "react";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import ListSubheader from "@mui/material/ListSubheader";
import AccountTreeIcon from "@mui/icons-material/AccountTree";
import LowPriorityIcon from "@mui/icons-material/LowPriority";
import BugReportIcon from "@mui/icons-material/BugReport";
import { Link } from "react-router-dom";
import AdjustIcon from "@mui/icons-material/Adjust";
import CheckCircleOutlineIcon from "@mui/icons-material/CheckCircleOutline";
import AssignmentIcon from "@mui/icons-material/Assignment";

export const mainListItems = (
  <div>
    <ListSubheader inset>Routinely</ListSubheader>
    <ListItem button component={Link} to="/open">
      <ListItemIcon>
        <AdjustIcon />
      </ListItemIcon>
      <ListItemText primary="Recent Open" />
    </ListItem>
    <ListItem button component={Link} to="/close">
      <ListItemIcon>
        <CheckCircleOutlineIcon />
      </ListItemIcon>
      <ListItemText primary="Recent Close" />
    </ListItem>
    <ListItem button component={Link} to="/assignments">
      <ListItemIcon>
        <AssignmentIcon />
      </ListItemIcon>
      <ListItemText primary="Assignments" />
    </ListItem>
  </div>
);

export const secondaryListItems = (
  <div>
    <ListSubheader inset>Engineering</ListSubheader>

    <ListItem button component={Link} to="/release">
      <ListItemIcon>
        <LowPriorityIcon />
      </ListItemIcon>
      <ListItemText primary="Release" />
    </ListItem>
  </div>
);

export const thirdListItems = (
  <div>
    <ListSubheader inset>Tools</ListSubheader>
    <ListItem button component={Link} to="/home/aboutci">
      <ListItemIcon>
        <BugReportIcon />
      </ListItemIcon>
      <ListItemText primary="TiFailureChaser" />
    </ListItem>
    <ListItem button component={Link} to="/home/databoard">
      <ListItemIcon>
        <AccountTreeIcon />
      </ListItemIcon>
      <ListItemText primary="Version Reports" />
    </ListItem>
  </div>
);

// Icons From：https://mui.com/components/material-icons/?query=project
