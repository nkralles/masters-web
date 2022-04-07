import {Button, Grid, LinearProgress, LinearProgressProps, ThemeProvider} from "@mui/material";
import cheese from "../../cheese.png";
import SportsGolfIcon from "@mui/icons-material/SportsGolf";
import GolfCourseIcon from '@mui/icons-material/GolfCourse';
import React from "react";
import {createTheme} from "@mui/material/styles";
import {Link} from 'react-router-dom';
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import GithubCorner from 'react-github-corner';


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
    const [progress, setProgress] = React.useState(10);
    React.useEffect(() => {
        const timer = setInterval(() => {
            setProgress((prevProgress) => (prevProgress >= 100 ? 10 : prevProgress + 10));
        }, 800);
        return () => {
            clearInterval(timer);
        };
    }, []);


    return (
        <ThemeProvider theme={theme}>
            <GithubCorner href="https://github.com/nkralles/masters-web" size={100}  octoColor='#ffdf00' bannerColor='#446f42'/>
            <Grid className="masters" container justifyContent="center" alignItems="center">
                <Grid container justifyContent="center" spacing={2}>
                    <Grid item>
                        <img src={cheese} className="App-logo" alt="logo"/>
                    </Grid>
                </Grid>
                {/*Placeholder till i etl data*/}
                {process.env.NODE_ENV !== "production" ? <Grid container justifyContent="center" spacing={2} mt={-20}>
                    <Grid item>
                        <Button variant="contained" size="large" color="golden">
                            <Link to="/leaderboard" style={{textDecoration: 'none'}}>
                                <GolfCourseIcon/>
                                Leaderboard
                            </Link>
                        </Button>
                    </Grid>
                    <Grid item>
                        <Button variant="contained" size="large" color="golden">
                            <Link to="/entries" style={{textDecoration: 'none'}}>
                                <SportsGolfIcon/>
                                Entries
                            </Link>
                        </Button>
                    </Grid>
                </Grid> :
                    <Box sx={{ width: '50%' }}>
                        <LinearProgressWithLabel value={progress} />
                    </Box>
                }
            </Grid>
        </ThemeProvider>
    );
}

export default Landing;


function LinearProgressWithLabel(props: LinearProgressProps & { value: number }) {
    return (
        <Box sx={{ display: 'flex', alignItems: 'center' }}>
            <Box sx={{ width: '100%', mr: 1 }}>
                <LinearProgress variant="determinate" {...props} />
            </Box>
            <Box sx={{ minWidth: 50 }}>
                <Typography variant="body2" color="text.secondary">COMING SOON...</Typography>
            </Box>
        </Box>
    );
}