import React, {useEffect, useState} from 'react';
import Nav from "../../components/Nav";
import Box from '@mui/material/Box';
import {DataGridPro, GridColDef} from "@mui/x-data-grid-pro";
import {api} from "../../lib/api";
import {GridRenderCellParams} from "@mui/x-data-grid/models/params/gridCellParams";
import {TextField} from "@mui/material";
import ClearIcon from '@mui/icons-material/Clear';
import SearchIcon from '@mui/icons-material/Search';
import IconButton from '@mui/material/IconButton';
import HtmlIcon from '@mui/icons-material/Html';
import SvgIcon from '@mui/material/SvgIcon';
import {ReactComponent as CsvIcon} from './csv.svg';


const renderPlayerGrid = function (v: GridRenderCellParams) {
    const style = {
        textAlign: 'center' as 'center'
    }
    const flag: string = flagMap[v.value.cc] !== undefined ? flagMap[v.value.cc] + " " : "";
    return (
        <div style={style}>
            <b>{flag}{v.value.first_name} {v.value.last_name}</b>
            <br></br>
            <span>{v.value.toPar}</span>
        </div>
    )
}

const flagMap: { [key: string]: string } = {
    "ARG": "🇦🇷",
    "AUS": "🇦🇺",
    "AUT": "🇦🇹",
    "BEL": "🇧🇪",
    "CAN": "🇨🇦",
    "CHI": "🇨🇱",
    "CHN": "🇨🇳",
    "COL": "🇨🇴",
    "DEN": "🇩🇰",
    "ENG": "🇬🇧",
    "ESP": "🇪🇸",
    "FRA": "🇫🇷",
    "GER": "🇩🇪",
    "IND": "🇮🇳",
    "IRL": "🇮🇪",
    "ITA": "🇮🇹",
    "JPN": "🇯🇵",
    "KOR": "🇰🇷",
    "MEX": "🇲🇽",
    "NIR": "🏴󠁧󠁢󠁮󠁩󠁲󠁿",
    "NOR": "🇳🇴",
    "NZL": "🇳🇿",
    "POL": "🇵🇱",
    "POR": "🇵🇹",
    "RSA": "🇿🇦",
    "SCO": "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
    "SVK": "🇸🇰",
    "SWE": "🇸🇪",
    "THA": "🇹🇭",
    "TPE": "🇹🇼",
    "USA": "🇺🇸",
    "VEN": "🇻🇪",
    "WAL": "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
    "ZIM": "🇿🇼"
}

const columns: GridColDef[] = [
    {
        field: 'entryName',
        headerName: 'Entry Name',
        width: 165,
    },
    {
        field: 'top12_1',
        headerName: 'Top 12',
        renderCell: renderPlayerGrid,
        sortable: false,
        width: 180,
    },
    {
        field: 'top12_2',
        headerName: 'Top 12',
        renderCell: renderPlayerGrid,
        sortable: false,
        width: 150,
    },
    {
        field: 'top12_3',
        headerName: 'Top 12',
        renderCell: renderPlayerGrid,
        sortable: false,
        width: 150,
    },
    {
        field: 'wildcard_1',
        headerName: 'Wildcard',
        renderCell: renderPlayerGrid,
        sortable: false,
        width: 150,
    },
    {
        field: 'wildcard_2',
        headerName: 'Wildcard',
        renderCell: renderPlayerGrid,
        sortable: false,
        width: 150,
    },
    {
        field: 'wildcard_3',
        headerName: 'Wildcard',
        renderCell: renderPlayerGrid,
        sortable: false,
        width: 150,
    },
    {
        field: 'wildcard_4',
        headerName: 'Wildcard',
        renderCell: renderPlayerGrid,
        sortable: false,
        width: 150,
    },
    {
        field: 'wildcard_5',
        headerName: 'Wildcard',
        renderCell: renderPlayerGrid,
        sortable: false,
        width: 150,
    },
    {
        field: 'winning_score',
        headerName: 'Winning Score',
        width: 100,
    },
    {
        field: 'total',
        headerName: 'Total',
        width: 100,
    },
];


function Entries() {
    const [searchText, setSearchText] = React.useState('');
    const [rows, setRows] = useState([]);
    const [filteredRows, setFilteredRows] = useState([]);

    useEffect(() => {
        ;(async () => {
            try {
                const entries: any = await api.getJSON('/entries')
                setRows(entries.data.map((e: {
                    golfers: any;
                    name: string;
                    winning_score: number;
                    total: number
                }) => {
                    return {
                        id: e.name,
                        entryName: e.name,
                        top12_1: e.golfers[0],
                        top12_2: e.golfers[1],
                        top12_3: e.golfers[2],
                        wildcard_1: e.golfers[3],
                        wildcard_2: e.golfers[4],
                        wildcard_3: e.golfers[5],
                        wildcard_4: e.golfers[6],
                        wildcard_5: e.golfers[7],
                        winning_score: e.winning_score,
                        total: e.total
                    }
                }))
                setFilteredRows(rows)
            } catch (err) {
                console.error(err)
            }
        })()
    }, [])

    const requestSearch = (searchValue: string) => {
        setSearchText(searchValue);
        const searchRegex = new RegExp(escapeRegExp(searchValue), 'i');
        const filteredRows = rows.filter((row: any) => {
            return Object.keys(row).some((field: any) => {
                return searchRegex.test(row[field].toString());
            });
        });
        setFilteredRows(filteredRows);
    };

    return (
        <div>
            <Nav/>
            <div style={{ display: 'flex', height: '90vh' }}>
                <div style={{ flexGrow: 1 }}>
                    <DataGridPro
                        components={{Toolbar: QuickSearchToolbar}}
                        columns={columns}
                        rows={searchText.length > 0 ? filteredRows : rows}
                        componentsProps={{
                            toolbar: {
                                value: searchText,
                                onChange: (event: React.ChangeEvent<HTMLInputElement>) =>
                                    requestSearch(event.target.value),
                                clearSearch: () => requestSearch(''),
                            },
                        }}
                        // initialState={{pinnedColumns: {left: ['entryName'], right: ['toPar']}}}
                    />
                </div>
            </div>
        </div>
    )
}

function escapeRegExp(value: string): string {
    return value.replace(/[-[\]{}()*+?.,\\^$|#\s]/g, '\\$&');
}

interface QuickSearchToolbarProps {
    clearSearch: () => void;
    onChange: () => void;
    value: string;
}

function QuickSearchToolbar(props: QuickSearchToolbarProps) {
    return (
        <Box
            sx={{
                p: 0.5,
                pb: 0,
            }}
        >
            <TextField
                variant="standard"
                value={props.value}
                onChange={props.onChange}
                placeholder="Search Entries…"
                InputProps={{
                    startAdornment: <SearchIcon fontSize="small"/>,
                    endAdornment: (
                        <IconButton
                            title="Clear"
                            aria-label="Clear"
                            size="small"
                            style={{visibility: props.value ? 'visible' : 'hidden'}}
                            onClick={props.clearSearch}
                        >
                            <ClearIcon fontSize="small"/>
                        </IconButton>
                    ),
                }}
                sx={{
                    width: {
                        xs: 1,
                        sm: 'auto',
                    },
                    m: (theme) => theme.spacing(1, 0.5, 1.5),
                    '& .MuiSvgIcon-root': {
                        mr: 0.5,
                    },
                    '& .MuiInput-underline:before': {
                        borderBottom: 1,
                        borderColor: 'divider',
                    },
                }}
            />
            <IconButton
                title="Open Raw HTML"
                aria-label="HTML"
                size="large"
                onClick={() => window.open("/api/entries.html", "_blank")}
            >
                <HtmlIcon fontSize="large"/>
            </IconButton>
            <IconButton
                title="Open Raw HTML"
                aria-label="HTML"
                size="large"
                onClick={() => window.open("/api/entries.csv", "_blank")}
            >
                <SvgIcon component={CsvIcon} fontSize="large" inheritViewBox/>
            </IconButton>

        </Box>
    );
}


export default Entries;