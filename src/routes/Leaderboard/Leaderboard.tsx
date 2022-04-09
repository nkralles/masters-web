import * as React from 'react';
import Nav from "../../components/Nav";
import {useEffect, useState} from "react";
import {api} from "../../lib/api";
import './Leaderboard.css'
import {Grid, Link, List, ListItem, ListItemAvatar, ListItemText} from "@mui/material";
import {FlagMap} from "../../lib/flags";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import DeleteIcon from '@mui/icons-material/Delete';

interface Score {
    first_name: string;
    last_name: string
    lastUpdated: string;
    countryCode: string;
    playerId: number;
    rank: number;
    total: number;
    standing: number;
    rounds: {
        round: number;
        toPar: number;
        lastUpdated: string;
    }[];
}

type Props = {
    className?: string
}

function Leaderboard(props?: Props) {
    const [scores, setScores] = useState([]);
    useEffect(() => {
        ;(async () => {
            try {
                const scores: Promise<Score[]> | any = await api.getJSON('/scores')
                setScores(scores.data)
            } catch (err) {
                console.error(err)
            }
        })()
    }, [])
    return (
        <div>
            <Nav/>
            <div className={props && props.className ? props.className : 'leaderboard'}>
                <Grid
                    container
                    direction="column"
                    justifyContent="center"
                    alignItems="center"
                >
                    <Grid item xl>
                        <List sx={{width: '100%', maxWidth: 360, bgcolor: '#DFDFD9FF'}}>
                            <ListItem key='====' alignItems="center">
                                <ListItemText
                                              primary={
                                                  <React.Fragment>
                                                      <Typography variant="h4" gutterBottom textAlign={"center"}>
                                                         LEADERBOARD
                                                      </Typography>
                                                  </React.Fragment>
                                              }
                                              secondary={
                                                  <React.Fragment>
                                                      <Grid
                                                          container
                                                          direction="row"
                                                          justifyContent="space-evenly"
                                                          alignItems="flex-start"
                                                      >
                                                          <Grid item>
                                                              <Typography variant="caption" display="block" gutterBottom justifyContent={'space-evenly'}>
                                                                  <Link href='/api/scores' target='_blank' alignItems={"start"}>internal source</Link>
                                                              </Typography>
                                                          </Grid>
                                                          <Grid item>
                                                              <Typography variant="caption" display="block" gutterBottom justifyContent={'space-evenly'}>
                                                                  <Link rel="noreferrer" href='https://www.masters.com/en_US/scores/feeds/2022/scores.json' target='_blank' alignItems={"end"}>external feed</Link>
                                                              </Typography>
                                                          </Grid>
                                                      </Grid>
                                                  </React.Fragment>
                                              }
                                >

                                </ListItemText>
                            </ListItem>
                            {scores.map((s: Score) => (
                                <ListItem key={s.playerId} alignItems="flex-start"
                                          secondaryAction={
                                              <Link fontSize={"x-large"} style={{textDecoration: 'none'}}>
                                                  <span>{s.total > 100 ? 'CUT' : s.total}</span>
                                              </Link>
                                          }
                                >
                                    <ListItemAvatar>
                                            <Link fontSize={"xx-large"} style={{textDecoration: 'none'}}>
                                                <span>{FlagMap[s.countryCode] ? FlagMap[s.countryCode] : '🏴‍☠️'}</span>
                                            </Link>
                                    </ListItemAvatar>
                                    <ListItemText primary={`${s.standing} ${s.first_name} ${s.last_name}`}
                                                  secondary={
                                                      <React.Fragment>
                                                          {s.rounds.map((r, index) => (
                                                              <Typography
                                                                  sx={{display: 'inline'}}
                                                                  component="span"
                                                                  variant="body2"
                                                                  color="text.primary"
                                                              >R{r.round} {r.toPar} {index !== s.rounds.length - 1 ? '|' : ''} </Typography>
                                                          ))}
                                                      </React.Fragment>
                                                  }
                                    >

                                    </ListItemText>
                                </ListItem>
                            ))}
                        </List>
                    </Grid>
                </Grid>
            </div>
        </div>
    );
}

export default Leaderboard;