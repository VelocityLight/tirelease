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
    <ListSubheader inset>Data Center</ListSubheader>
    <ListItem button component={Link} to="/home/open">
      <ListItemIcon>
        <AdjustIcon />
      </ListItemIcon>
      <ListItemText primary="Issues" />
    </ListItem>
    {/* <ListItem button component={Link} to="/home/close">
      <ListItemIcon>
        <CheckCircleOutlineIcon />
      </ListItemIcon>
      <ListItemText primary="Recent Close" />
    </ListItem> */}
    <ListItem button component={Link} to="/home/assignments">
      <ListItemIcon>
        <AssignmentIcon />
      </ListItemIcon>
      <ListItemText primary="Assignments" />
    </ListItem>
  </div>
);

export const secondaryListItems = (
  <div>
    <ListSubheader inset>Version Management</ListSubheader>

    <ListItem button component={Link} to="/home/release">
      <ListItemIcon>
        <LowPriorityIcon />
      </ListItemIcon>
      <ListItemText primary="Release" />
    </ListItem>
  </div>
);

export const thirdListItems = (
  <div>
    <ListSubheader inset>CI/CD Tools</ListSubheader>
    <ListItem button component={Link} to="/home/aboutci">
      <ListItemIcon>
        <BugReportIcon />
      </ListItemIcon>
      <ListItemText primary="TiFailureChaser" />
    </ListItem>
    {/* <ListItem button component={Link} to="/home/databoard">
      <ListItemIcon>
        <AccountTreeIcon />
      </ListItemIcon>
      <ListItemText primary="Version Reports" />
    </ListItem> */}
  </div>
);

// Icons From：https://mui.com/components/material-icons/?query=project
