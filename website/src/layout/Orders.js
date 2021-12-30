import * as React from 'react';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import DashboardIcon from '@mui/icons-material/Dashboard';
import { Link } from 'react-router-dom';

export const mainListItems = (
  <div>
    <ListSubheader inset>Data Center</ListSubheader>
    <ListItem button component={Link} to="/home/example">
      <ListItemIcon>
        <DashboardIcon />
      </ListItemIcon>
      <ListItemText primary="Repositories" />
    </ListItem>
    <ListItem button component={Link} to="/home/triage">
      <ListItemIcon>
        <DashboardIcon />
      </ListItemIcon>
      <ListItemText primary="Issues" />
    </ListItem>
  </div>
);

export const secondaryListItems = (
  <div>
    
    <ListSubheader inset>Project Management</ListSubheader>
    <ListItem button component={Link} to="/home/triage">
      <ListItemIcon>
        <DashboardIcon />
      </ListItemIcon>
      <ListItemText primary="Projects" />
    </ListItem>
  </div>
);

export const thirdListItems = (
  <div>
    <ListSubheader inset>Reports</ListSubheader>
    <ListItem button component={Link} to="/home/databoard">
      <ListItemIcon>
        <DashboardIcon />
      </ListItemIcon>
      <ListItemText primary="About Project" />
    </ListItem>
  </div>
);
