import React from 'react';
import Toolbar from '@mui/material/Toolbar';
import AppBar from '@mui/material/AppBar';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';

export default function Navigation() {

    const routes = [
        {
            "name": "Home",
            "path": "/",
        },
        {
            "name": "Your Words",
            "path": "/words"
        },
        {
            "name": "Sign Up",
            "path": "/sign-up",
        },
        {
            "name": "Sign In",
            "path": "/sign-in",
        },
    ];

    return (
        <>
            <AppBar color="secondary" position="static">
                <Toolbar sx={{ margin: 'auto' }}>
                    {
                        routes.map(route => (
                            <><Button href={route.path} color="inherit">{route.name}</Button>
                            <Box sx={{m: 4}}/>
                            </>
                        ))
                    }
                    <Box sx={{mr: -10}}></Box>
                </Toolbar>
            </AppBar>
        </>
    );
}

export {Navigation};