import * as React from 'react';
import { DataGrid } from '@mui/x-data-grid';

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
            // checkboxSelection
            // disableSelectionOnClick
        />
        </div>
        </>
    );
}
