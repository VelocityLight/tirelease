import { useEffect, useState } from "react";
import Stack from "@mui/material/Stack";

import { VersionTable } from "./VersionTable";
import { VersionSearch } from "./VersionSearch";
import { VersionAdd } from "./VersionAdd";
import { useQuery } from "react-query";
import { url } from "../../utils";

export default function Version() {
  const { isLoading, error, data } = useQuery("versions", () => {
    return fetch(url("version")).then((res) => {
      console.log(res);
      return res.json();
    });
  });

  if (data) {
    console.log("data", data);
  }

  return (
    <>
      {isLoading && <p>Loading...</p>}
      {error && <p>Error: {error.message}</p>}
      {data && (
        <Stack spacing={1}>
          <Stack direction={"row"} justifyContent={"flex-start"} spacing={2}>
            <VersionSearch />
            <VersionAdd />
          </Stack>
          <Stack direction={"row"} justifyContent={"flex-start"} spacing={2}>
            <VersionTable data={data.data} />
          </Stack>
        </Stack>
      )}
    </>
  );
}
