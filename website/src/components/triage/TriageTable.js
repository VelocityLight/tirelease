import React, { useEffect, useState } from "react";
import Link from '@mui/material/Link';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Title from '../databoard/Title';
import axios from 'axios';

function preventDefault(event) {
    event.preventDefault();
}

export default function TriageTable() {

    // make the fetch the first time your component mounts
    const [data, setData] = useState([]);
    useEffect(() => {
        axios.get('http://localhost:8080/triage/select').then(response => setData(response.data.data));
    }, []);

    return (
        <React.Fragment>
            <Title>Triage Item List</Title>
            <Table size="small">
                <TableHead>
                    <TableRow>
                        <TableCell>Project</TableCell>
                        <TableCell>Repository</TableCell>
                        <TableCell>IssueID</TableCell>
                        <TableCell>PullRequestID</TableCell>
                        <TableCell>Triage Status</TableCell>
                        <TableCell align="right">Comment</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                {data.map((row) => (
                    <TableRow key={row.id}>
                        <TableCell>{row.project_name}</TableCell>
                        <TableCell>{row.repo}</TableCell>
                        <TableCell>{row.issue_id}</TableCell>
                        <TableCell>{row.pull_request_id}</TableCell>
                        <TableCell>{row.status}</TableCell>
                        <TableCell align="right">{row.comment}</TableCell>
                    </TableRow>
                ))}
                </TableBody>
            </Table>
            <Link color="primary" href="#" onClick={preventDefault} sx={{ mt: 3 }}>
                See more
            </Link>
        </React.Fragment>
    );
}

