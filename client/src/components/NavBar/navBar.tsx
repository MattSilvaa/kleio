import React from 'react'
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';

function NavBar() {
    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar color="inherit" position="static">
                <Toolbar>
                    <Typography variant="h6" component="div" sx={{ flexGrow: 1 }} align="center">
                        Kleio
                    </Typography>
                </Toolbar>
            </AppBar>
        </Box>
    );
}

export default NavBar;
