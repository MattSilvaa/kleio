import {createTheme} from '@mui/material/styles'

export const navBarTheme = createTheme({
        palette: {
            primary: {
                main: '#757075',
                // light: will be calculated from palette.primary.main,
                // dark: will be calculated from palette.primary.main,
                // contrastText: will be calculated to contrast with palette.primary.main
            }
        }
    }
)