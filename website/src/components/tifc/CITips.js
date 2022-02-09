import * as React from 'react';
import Tooltip from '@mui/material/Tooltip';
import Link from '@mui/material/Link';
import IconButton from '@mui/material/IconButton';
import HelpIcon from '@mui/icons-material/Help';
import DataThresholdingIcon from '@mui/icons-material/DataThresholding';

export const CITips = () => {
    const tips1 = "We run tests every hour, define failed cases as testbug, the rest of cases are productbug.\
        If a case has not failed in the recent 5 days in hourly run, we think this case was possibly fixed.";
    const tips2 = "Open unstable insight page";
    const url = "https://tiinsights.pingcap.net/watch/public/dashboard/2b613238-4910-4d8a-9872-def03dbff468"

    return (
        <>
        <Tooltip title={tips1} sx={{ m: 0, display: 'inline-flex'}}>
            <IconButton color="primary">
                <HelpIcon />
            </IconButton>
        </Tooltip>
        <Tooltip title={tips2} sx={{ m: 0, display: 'inline-flex'}}>
            <Link href={url} target="_blank">
            <IconButton color="primary">
                <DataThresholdingIcon />
            </IconButton>
            </Link>
        </Tooltip>
        </>
    );
}