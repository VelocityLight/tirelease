import { useEffect, useState } from "react";
import Stack from "@mui/material/Stack";

import { VersionTable } from "./VersionTable";
import { VersionSearch } from "./VersionSearch";
import { VersionAdd } from "./VersionAdd";

export default function Version() {
    const [table, setTable] = useState(null);
    const refreshTable = () => {
        fetch("/version/list")
        .then(response => response.json())
        .then(data => {
            setTable(data.data);
        })
        .catch((e) => {
            console.log(e);
        });
    }
    useEffect(() => {
        refreshTable();
    }, []); // empty dependency array

    return (
        <>
        <Stack spacing={1}>
            <Stack direction={"row"} justifyContent={"flex-start"} spacing={2}>
                <VersionSearch />
                <VersionAdd />
            </Stack>
            <Stack direction={"row"} justifyContent={"flex-start"} spacing={2}>
                <VersionTable data={table}/>
            </Stack>
        </Stack>
        </>
    );
}
