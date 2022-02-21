import * as React from 'react';
import { DataGrid, GridToolbar } from '@mui/x-data-grid';

import VersionTableColumns from "./VersionTableColumns";

export const VersionTable = ({data}) => {
    console.log(data);
    return (
        <>
        <div style={{ height: 650, width: '100%' }}>
        <DataGrid
            rows = {data}
            columns = {VersionTableColumns}
            pageSize = {100}
            rowsPerPageOptions = {[100]}
            components={{
                Toolbar: GridToolbar,
            }}
            rowHeight = {70}
            showCellRightBorder = {true}
            showColumnRightBorder = {false}
            // checkboxSelection
            disableSelectionOnClick
        />
        </div>
        </>
    );
}
