import React, {useEffect, useState} from 'react';
import Nav from "../../components/Nav";
import {DataGridPro, GridColDef, GridRowsProp} from "@mui/x-data-grid-pro";
import {api} from "../../lib/api";

const columns: GridColDef[] = [
    {
        field: 'entryName',
        headerName: 'Entry Name',
        width: 165,
    },
    {
        field: 'top12_1',
        headerName: 'Top 12',
        width: 150,
    },
    {
        field: 'top12_2',
        headerName: 'Top 12',
        width: 150,
    },
    {
        field: 'top12_3',
        headerName: 'Top 12',
        width: 150,
    },
    {
        field: 'wildcard_1',
        headerName: 'Wildcard',
        width: 150,
    },
    {
        field: 'wildcard_2',
        headerName: 'Wildcard',
        width: 150,
    },
    {
        field: 'wildcard_3',
        headerName: 'Wildcard',
        width: 150,
    },
    {
        field: 'wildcard_4',
        headerName: 'Wildcard',
        width: 150,
    },
    {
        field: 'wildcard_5',
        headerName: 'Wildcard',
        width: 150,
    },
    {
        field: 'toPar',
        headerName: 'To Par',
        width: 100,
    },
    {
        field: 'winning_score',
        headerName: 'Winning Score',
        width: 100,
    },
];


function Entries() {
    const [rows, setRows] = useState([]);

    useEffect(() => {
        ;(async () => {
            try {
                const entries:any = await api.getJSON('/entries')
                setRows(entries.data.map((e: {
                    golfers: any;
                    name: string;
                    winning_score: number}) => {
                    return {
                        id: e.name,
                        entryName: e.name,
                        top12_1: `${e.golfers[0].first_name} ${e.golfers[0].last_name}`,
                        top12_2: `${e.golfers[1].first_name} ${e.golfers[1].last_name}`,
                        top12_3: `${e.golfers[2].first_name} ${e.golfers[2].last_name}`,
                        wildcard_1: `${e.golfers[3].first_name} ${e.golfers[3].last_name}`,
                        wildcard_2: `${e.golfers[4].first_name} ${e.golfers[4].last_name}`,
                        wildcard_3: `${e.golfers[5].first_name} ${e.golfers[5].last_name}`,
                        wildcard_4: `${e.golfers[6].first_name} ${e.golfers[6].last_name}`,
                        wildcard_5: `${e.golfers[7].first_name} ${e.golfers[7].last_name}`,
                        winning_score: e.winning_score
                    }
                }))
            } catch (err) {
                console.error(err)
            }
        })()
    }, [])
    //const entries = await api.getJSON('entries')

    const [nbRows, setNbRows] = React.useState(5);

    return (
        <div>
            <Nav/>
            <div style={{width: '100%'}}>
                <DataGridPro autoHeight
                             columns={columns}
                             rows={rows}
                             initialState={{pinnedColumns: {left: ['entryName'], right: ['toPar']}}}
                />
            </div>
        </div>
    )
}

export default Entries;