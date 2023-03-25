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
            "name": "Practice",
            "path": "/practice"
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
                <Toolbar>
                    {routes.map((route, index) => (<Box sx={{mt: 2, mr: 5}} key={index} ><Button href={route.path} color="inherit">{route.name}</Button><Box sx={{m: 4}}/></Box>))}
                    <Box sx={{mr: 10}}></Box>
                </Toolbar>
            </AppBar>
        </>
    );
}

export {Navigation};