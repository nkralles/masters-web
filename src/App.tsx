import React from 'react';
import './App.css';
import {useRoutes} from 'react-router-dom';
import Landing from "./routes/landing/Landing";
import Entries from "./routes/entries/Entries";

import { LicenseInfo } from '@mui/x-data-grid-pro';
import Leaderboard from "./routes/Leaderboard/Leaderboard";

LicenseInfo.setLicenseKey(
    '0c120603ea1187358fdee5a806215fcdT1JERVI6NDAzMzQsRVhQSVJZPTE2Nzk3MTMyMTkwMDAsS0VZVkVSU0lPTj0x',
);

const App: React.FC = (): JSX.Element => {
    const landingRoute = {
        path: '/',
        element: <Landing/>,
        children: [
            {path: '/', element: <Landing/>},
        ],
    };
    const entriesRoute = {
        path: '/entries',
        element: <Entries/>,
    };
    const leaderboardRoute = {
        path: '/leaderboard',
        element: <Leaderboard/>,
    };

    const routing = useRoutes([landingRoute, entriesRoute, leaderboardRoute]);
    return <>{routing}</>;
};

export default App;
