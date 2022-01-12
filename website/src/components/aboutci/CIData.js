import { useState } from "react";
import Stack from "@mui/material/Stack";
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import AdapterDateFns from '@mui/lab/AdapterDateFns';
import LocalizationProvider from '@mui/lab/LocalizationProvider';
import DatePicker from '@mui/lab/DatePicker';

import { CIJobNameSelector} from "./CIJobNameSelector";
import { CITable} from "./CITable";

function addDays(date, days) {
    var result = new Date(date);
    result.setDate(result.getDate() + days);
    return result;
}

export default function CIData() {
    const [jobName, setJobName] = useState("tidb-unit-test-hourly");
    const [timestamp, setTimestamp] = useState(addDays(new Date(), -3));

    return (
        <>
        <Stack spacing={1}>
            <Stack direction={"row"} justifyContent={"flex-start"} spacing={2}>
                <CIJobNameSelector jobName={jobName} setJobName={setJobName}/>
                <LocalizationProvider dateAdapter={AdapterDateFns}>
                    <DatePicker
                        label="timestamp"
                        value={timestamp}
                        onChange={(newTime) => {
                            setTimestamp(newTime);
                        }}
                        renderInput={(params) => <TextField {...params} />}
                    />
                </LocalizationProvider>
                <Button variant="contained">Query</Button>
            </Stack>
            {jobName != null && timestamp != null && (
                <CITable jobName={jobName} timestamp={Math.round(timestamp.getTime()/1000)} />
            )}
        </Stack>
        </>
    );
}
