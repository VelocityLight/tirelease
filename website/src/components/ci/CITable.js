import * as React from 'react';
import Link from '@mui/material/Link';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import axios from 'axios';

async function select() {
    let data;
    await axios({
        method: 'get',
        url: '/triage/select',
    })
    .then (function (response){
        data = response.data.data;
    });
    return data;
}

export default function CITable() {
    const rows = [];

    return (
        <>
            <React.Fragment>
            <Table size="small">
                <TableHead>
                    <TableRow>
                        <TableCell>Project</TableCell>
                        <TableCell>Repository</TableCell>
                        <TableCell>IssueID</TableCell>
                        <TableCell>PullRequestID</TableCell>
                        <TableCell>Triage Status</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                {rows.map((row) => (
                    <TableRow key={row.id}>
                        <TableCell>{row.project_name}</TableCell>
                        <TableCell>{row.repo}</TableCell>
                        <TableCell>
                            <Link color="primary" href={row.issue_url}>
                                {row.issue_id}
                            </Link>
                        </TableCell>
                        <TableCell>{row.pull_request_id}</TableCell>
                        <TableCell>{row.status}</TableCell>
                    </TableRow>
                ))}
                </TableBody>
            </Table>
            </React.Fragment>
        </>
    );
}