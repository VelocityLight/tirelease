import * as React from 'react';
import Container from '@mui/material/Container';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import Layout from '../layout/Layout'
import CITable from '../components/ci/CITable';

const AboutCI = () => {
    return (
        <Layout>
            <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
                <Paper sx={{ p: 2, display: 'flex', flexDirection: 'column' }}>
                    <CITable />
                </Paper>
            </Container>
        </Layout>
    )
};

export default AboutCI;