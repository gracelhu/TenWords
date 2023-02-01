import React from 'react';
import PageTemplate from '../pages/PageTemplate';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button'

import { useState } from 'react';

export default function SignIn() {
	
    return (
        <PageTemplate>
            <h1 style={{textAlign: "center"}}>Sign In</h1>
			
			<div className="Auth-form-container">
			
			<TextField
				id="outlined-basic"
				label="Username"
				variant="filled"
			/>
			
			<TextField
				id="filled-password-input"
				label="Password"
				type="password"
				autoComplete="current-password"
				variant="filled"
			/>
			
			<Button
				color='secondary'
				variant="contained"
				onClick={() => {
					alert('Trigger login');
				}}
			>
				Sign In
			</Button>
			
			</div>
			
        </PageTemplate>
    );
}