import React from 'react';
import Toolbar from '@mui/material/Toolbar';
import AppBar from '@mui/material/AppBar';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import { useNavigate } from 'react-router-dom';
import NavigationChild from './NavigationChild';

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
        <NavigationChild logout={logout}/>
    );
}

export {Navigation};