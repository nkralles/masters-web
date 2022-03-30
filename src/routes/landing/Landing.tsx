import {Button, Grid, ThemeProvider} from "@mui/material";
import cheese from "../../cheese.png";
import SportsGolfIcon from "@mui/icons-material/SportsGolf";
import GolfCourseIcon from '@mui/icons-material/GolfCourse';
import React from "react";
import {createTheme} from "@mui/material/styles";
import {Link, Outlet} from 'react-router-dom';


const theme = createTheme({
    palette: {
        golden: {
            main: '#ffdf00',
            contrastText: '#446f42',
        },
    },
});

declare module '@mui/material/styles' {
    interface Palette {
        golden: Palette['primary'];
    }

    // allow configuration using `createTheme`
    interface PaletteOptions {
        golden?: PaletteOptions['primary'];
    }
}

// Update the Button's color prop options
declare module '@mui/material/Button' {
    interface ButtonPropsColorOverrides {
        golden: true;
    }
}
function Landing() {
    return (
        <ThemeProvider theme={theme}>
            <Grid className="masters" container justifyContent="center" alignItems="center">
                <Grid container justifyContent="center" spacing={2}>
                    <Grid item>
                        <img src={cheese} className="App-logo" alt="logo"/>
                    </Grid>
                </Grid>
                <Grid container justifyContent="center" spacing={2} mt={-20}>
                    <Grid item  >
                        <Button variant="contained" size="large" color="golden">
                            <Link to="/" style={{ textDecoration: 'none' }}>
                                <GolfCourseIcon/>
                                Leaderboard
                            </Link>
                        </Button>
                    </Grid>
                    <Grid item >
                        <Button variant="contained" size="large" color="golden">
                            <Link to="/entries" style={{ textDecoration: 'none' }}>
                                <SportsGolfIcon/>
                                Entries
                            </Link>
                        </Button>
                    </Grid>
                </Grid>
            </Grid>
        </ThemeProvider>
    );
}

export default Landing;