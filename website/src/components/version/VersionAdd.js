import * as React from 'react';
import Button from '@mui/material/Button';

export const VersionAdd = () => {
    const handleApplyChanges = React.useCallback(() => {
    }, []);

    return (
        <Button variant="contained" onClick={handleApplyChanges}>Add</Button>
    );
}
