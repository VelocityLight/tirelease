import { useState } from "react";
import Stack from "@mui/material/Stack";
import Button from '@mui/material/Button';

import { CIJobNameSelector} from "./CIJobNameSelector";
import { CIDatePicker} from "./CIDatePicker";
import { CITable } from "./CITable";

function addDays(date, days) {
    var result = new Date(date);
    result.setDate(result.getDate() + days);
    return result;
}

export default function CIData() {
    const [jobName, setJobName] = useState("tidb-unit-test-hourly");
    const [timestamp, setTimestamp] = useState(addDays(new Date(), -3));
    const [table, setTable] = useState(null);

    const refreshTable = (jobName, timestamp) => {
        const tb = <CITable jobName={jobName} timestamp={Math.round(timestamp.getTime()/1000)} />;
        setTable(tb)
    }

    return (
        <>
        <Stack spacing={1}>
            <Stack direction={"row"} justifyContent={"flex-start"} spacing={2}>
                <CIJobNameSelector jobName={jobName} setJobName={setJobName}/>
                <CIDatePicker timestamp={timestamp} setTimestamp={setTimestamp}/>
                <Button variant="contained" onClick={() => refreshTable(jobName, timestamp)}>Query</Button>
            </Stack>
            <Stack direction={"row"} justifyContent={"flex-start"} spacing={2}>
                {table}
            </Stack>
        </Stack>
        </>
    );
}
