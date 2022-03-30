import React from 'react';
import Nav from "../../components/Nav";
import {DataGridPro, GridColDef, GridRowsProp} from "@mui/x-data-grid-pro";

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

const rows: GridRowsProp = [
    {
        id: 1,
        entryName: "Nicholas Kralles"
    }
]

function Entries() {
    const [nbRows, setNbRows] = React.useState(5);

    return (
        <div>
            <Nav/>
            <div style={{width: '100%'}}>
                <DataGridPro autoHeight
                             columns={columns}
                             rows={rows}
                             initialState={{ pinnedColumns: { left: ['entryName'], right: ['toPar'] } }}
                />
            </div>
        </div>
    )
}

export default Entries;