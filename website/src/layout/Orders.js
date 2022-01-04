import * as React from 'react';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import AccountTreeIcon from '@mui/icons-material/AccountTree';
import BugReportIcon from '@mui/icons-material/BugReport';
import StorageIcon from '@mui/icons-material/Storage';
import FeaturedPlayListIcon from '@mui/icons-material/FeaturedPlayList';
import { Link } from 'react-router-dom';

export const mainListItems = (
  <div>
    <ListSubheader inset>Data Center</ListSubheader>
    <ListItem button component={Link} to="/home/example">
      <ListItemIcon>
        <StorageIcon />
      </ListItemIcon>
      <ListItemText primary="Repositories" />
    </ListItem>
    <ListItem button component={Link} to="/home/triage">
      <ListItemIcon>
        <FeaturedPlayListIcon />
      </ListItemIcon>
      <ListItemText primary="Issues" />
    </ListItem>
  </div>
);

export const secondaryListItems = (
  <div>
    
    <ListSubheader inset>Process Management</ListSubheader>
    <ListItem button component={Link} to="/home/triage">
      <ListItemIcon>
        <AccountTreeIcon />
      </ListItemIcon>
      <ListItemText primary="Projects" />
    </ListItem>
  </div>
);

export const thirdListItems = (
  <div>
    <ListSubheader inset>Visual Reports</ListSubheader>
    <ListItem button component={Link} to="/home/aboutci">
      <ListItemIcon>
        <BugReportIcon />
      </ListItemIcon>
      <ListItemText primary="About CI" />
    </ListItem>
    <ListItem button component={Link} to="/home/databoard">
      <ListItemIcon>
        <AccountTreeIcon />
      </ListItemIcon>
      <ListItemText primary="About Project" />
    </ListItem>
  </div>
);

// Icons From：https://mui.com/components/material-icons/?query=project