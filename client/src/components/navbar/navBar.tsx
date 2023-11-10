import * as React from 'react'
import AppBar from '@mui/material/AppBar'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import Container from '@mui/material/Container'
import Button from '@mui/material/Button'
import {noop} from 'lodash'
import SearchIcon from '@mui/icons-material/Search'
import {navBarTheme} from '../theme'
import {ThemeProvider} from '@mui/material'
import useMediaQuery from '@mui/material/useMediaQuery'
import {fetchJobCount} from "../../api/jobServices";
import {JobLocations, JobTitles} from "../shared";

const pages = ['About', 'FAQ', 'Contact Us']

export const NavBar = () => {
    const isMobile = useMediaQuery('(max-width:600px)')

    return (
        <ThemeProvider theme={navBarTheme}>
            <AppBar position="static">
                <Toolbar>
                    <Container
                        sx={{
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: isMobile ? 'center' : 'flex-start',
                            flexDirection: isMobile ? 'column' : 'row',
                        }}
                    >
                        <SearchIcon sx={{display: 'flex', mr: 1}}/>
                        <Typography
                            variant="h6"
                            noWrap
                            component="a"
                            href="#app-bar-with-responsive-menu"
                            sx={{
                                display: 'flex',
                                mr: 2,
                                fontFamily: 'monospace',
                                fontWeight: 700,
                                letterSpacing: '.3rem',
                                color: 'inherit',
                                textDecoration: 'none',
                            }}
                        >
                            Kleio
                        </Typography>
                    </Container>
                    <Container
                        sx={{
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: isMobile ? 'center' : 'flex-end',
                            gap: 5,
                            flexDirection: isMobile ? 'column' : 'row',
                        }}
                    >
                        {pages.map((page) => (
                            <Button
                                key={page}
                                onClick={() => fetchJobCount({date: '01-01-2023', location: JobLocations.DENVER, jobType: JobTitles.DATA_SCIENTIST})}
                                sx={{
                                    my: 2,
                                    color: 'white',
                                    display: isMobile ? 'block' : 'inline-block',
                                }}
                            >
                                {page}
                            </Button>
                        ))}
                    </Container>
                </Toolbar>
            </AppBar>
        </ThemeProvider>
    )
}
