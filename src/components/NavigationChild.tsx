import React from 'react';
import Toolbar from '@mui/material/Toolbar';
import AppBar from '@mui/material/AppBar';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import { useNavigate } from 'react-router-dom';

export default function Navigation(props: any) {

    return (
        <>
            <AppBar color="secondary" position="static">
                <Toolbar sx={{ margin: 'auto' }}>
                    <Button href="/" color="inherit" data-testid="home">Home</Button>
                    <Box sx={{m: 4}}/>
                    <Button href="/words" color="inherit" data-testid="your-words">Your Words</Button>
                    <Box sx={{m: 4}}/>
                    {localStorage.getItem("date") === null || localStorage.getItem("date") === undefined ?
                    <>
                        <Button href="/sign-in" color="inherit">Sign In</Button>
                        <Box sx={{m: 4}}/>
                        <Button href="/sign-up" color="inherit">Sign Up</Button>
                        
                    </> : <>
                        <Button onClick={() => props.logout()} color="inherit">Sign Out</Button>
                    </>
                    }
                    <Box sx={{mr: -10}}></Box>
                </Toolbar>
            </AppBar>
        </>
    );
}

export {Navigation};