import { Stack } from "@mui/material";
import { useState } from "react";
import { CITable } from "./CITable";
import { useQuery } from "react-query";
  
function RenderCITable({ jobName, timestamp }) {
    const { isLoading, error, data } = useQuery("renderCITable", () => {
        return fetch("http://172.16.5.2:30792/report/?job_name=" + jobName + "&timestamp=" + timestamp)
        .then((res) => {
            const data = res.json();
            console.log(data);
            return data;
        })
        .catch((e) => {
            console.log(e);
        });
    });
    console.log(isLoading, error, data);
    if (isLoading) {
        return <p>Loading...</p>;
    }
    if (error) {
        return <p>Error: {error.message}</p>;
    }
    console.log(data);
    return (
        <CITable data={data} />
    );
}

export default function CIData() {
    const [jobName, setJobName] = useState("tidb-unit-test-hourly");
    const [timestamp, setTimestamp] = useState("1510150535");

    return (
        <>
        <Stack spacing={1}>
          <RenderCITable jobName={jobName} timestamp={timestamp} />
        </Stack>
        </>
    );
}
