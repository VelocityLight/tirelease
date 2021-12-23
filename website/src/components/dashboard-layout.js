import * as React from 'react';
import { styled, createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import MuiDrawer from '@mui/material/Drawer';
import Box from '@mui/material/Box';
import MuiAppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';
import Badge from '@mui/material/Badge';
import Container from '@mui/material/Container';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import Link from '@mui/material/Link';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import NotificationsIcon from '@mui/icons-material/Notifications';
import { DashboardNavbar } from './dashboard-navbar';
import { DashboardSidebar } from './dashboard-sidebar';

const mdTheme = createTheme();

// export const DashboardLayout = (props) => {
//     const { children } = props;
//     const [ open ] = React.useState(true);

//     return (
//         <ThemeProvider theme={mdTheme}>
//             <DashboardLayoutRoot>
//                 {children}
//             </DashboardLayoutRoot>
//             <DashboardNavbar 
//                 open={open}
//             />
//             <DashboardSidebar
//                 open={open}
//             />
//         </ThemeProvider>
//     );
// };

function DashboardContent() {
    const { children } = React.useState(true);

    const [open, setOpen] = React.useState(true);
    const toggleDrawer = () => {
        setOpen(!open);
    };

    return (
        <ThemeProvider theme={mdTheme}>
            <Box sx={{ display: 'flex' }}>
                <CssBaseline />
                <DashboardNavbar 
                    open = {open}
                    toggleDrawer = {toggleDrawer}
                />
                <DashboardSidebar 
                    open = {open}
                    toggleDrawer = {toggleDrawer}
                />
                <Box
                    component="main"
                    sx={{
                        backgroundColor: (theme) =>
                        theme.palette.mode === 'light'
                            ? theme.palette.grey[100]
                            : theme.palette.grey[900],
                        flexGrow: 1,
                        height: '100vh',
                        overflow: 'auto',
                    }}
                >
                    {children}
                </Box>
            </Box>
        </ThemeProvider>
    );
}
export default function Dashboard() {
    return <DashboardContent />;
  }