import * as React from 'react';
import { DataGrid, GridToolbar } from '@mui/x-data-grid';

import CIColumnsGrid from "./CIColumnsGrid";

export const RenderCITableGrid = ({data}) => {
    return (
        <>
        <div style={{ height: 600, width: '100%' }}>
        <DataGrid
            rows = {data}
            columns = {CIColumnsGrid}
            pageSize = {100}
            rowsPerPageOptions = {[100]}
            components={{
                Toolbar: GridToolbar,
            }}
            // checkboxSelection
            disableSelectionOnClick
        />
        </div>
        </>
    );
}
