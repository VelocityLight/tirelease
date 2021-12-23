import * as React from 'react';
import Container from '@mui/material/Container';
import Layout from '../layout/Layout'

const Example = () => {
    return (
        <>
            <Layout>
                <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
                    <p>Hello TiRelease</p>
                </Container>
            </Layout>
        </>
    )
};

export default Example;