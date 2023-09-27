import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import SearchIcon from '@mui/icons-material/Search';
import {navBarTheme} from '../theme'
import {ThemeProvider} from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2'; // Grid version 2

const pages = ['About', 'FAQ', 'Blog'];

export const NavBar = () => {

    return (
        <ThemeProvider theme={navBarTheme}>
            <AppBar position="static">
                <Container maxWidth="xl">
                    <Toolbar disableGutters>
                        <SearchIcon sx={{display: {xs: 'none', md: 'flex'}, mr: 1}}/>
                        <Typography
                            variant="h6"
                            noWrap
                            component="a"
                            href="#app-bar-with-responsive-menu"
                            sx={{
                                mr: 2,
                                display: {xs: 'none', md: 'flex'},
                                fontFamily: 'monospace',
                                fontWeight: 700,
                                letterSpacing: '.3rem',
                                color: 'inherit',
                                textDecoration: 'none',
                            }}
                        >
                            Kleio
                        </Typography>
                            <Box sx={{flexGrow: 12, display: {xs: 'none', md: 'flex'}}}>
                                {pages.map((page) => (
                                    <Button
                                        key={page}
                                        onClick={() => {
                                        }}
                                        sx={{my: 2, color: 'white', display: 'block'}}
                                    >
                                        {page}
                                    </Button>
                                ))}
                            </Box>
                    </Toolbar>
                </Container>
            </AppBar>
        </ThemeProvider>
    );
}