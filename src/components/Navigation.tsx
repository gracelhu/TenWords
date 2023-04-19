import React from 'react';
import Toolbar from '@mui/material/Toolbar';
import AppBar from '@mui/material/AppBar';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import { useNavigate } from 'react-router-dom';

export default function Navigation() {

    const navigate = useNavigate();
    const logout = () => {
        localStorage.removeItem("username");
        localStorage.removeItem("password");
        localStorage.removeItem("date");
        navigate('/words', {});
        return;
    }

    return (
        <>
            <AppBar color="secondary" position="static">
                <Toolbar sx={{ margin: 'auto' }}>
                    <Button href="/" color="inherit">Home</Button>
                    <Box sx={{m: 4}}/>
                    <Button href="/words" color="inherit">Your Words</Button>
                    <Box sx={{m: 4}}/>
                    {localStorage.getItem("date") === null || localStorage.getItem("date") === undefined ?
                    <>
                        <Button href="/sign-in" color="inherit">Sign In</Button>
                        <Box sx={{m: 4}}/>
                        <Button href="/sign-up" color="inherit">Sign Up</Button>
                        
                    </> : <>
                        <Button onClick={() => logout()} color="inherit">Sign Out</Button>
                    </>
                    }
                    <Box sx={{mr: -10}}></Box>
                </Toolbar>
            </AppBar>
        </>
    );
}

export {Navigation};