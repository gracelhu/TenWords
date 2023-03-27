import React from 'react';
import PageTemplate from '../pages/PageTemplate';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button'

import { useRef, useState, useEffect } from 'react';

export default function SignIn() {

	const [username, setUsername] = useState("");
	const [password, setPassword] = useState(""); 

	const handleClick = async() => {
		fetch("/auth/"+username+"/"+password)
        .then(res => res.json())
        .then(
            (result) => {
                console.log("Trying to login with " + username + ", " + password);
				console.log(result);
            },
        )
        .catch(error => {
            console.log("Login error");
            console.warn(error)
        })
	}

    return (
        <PageTemplate>
            <h1 style={{textAlign: "center"}}>Sign In</h1>
			
			<div className="Auth-form-container">
			
			<TextField
				value={username}
				onChange={(e) => setUsername(e.target.value)}
				id="outlined-basic"
				label="Username"
				variant="filled"
				data-cy="field_username"
			/>
			
			<TextField
				value={password}
				onChange={(e) => setPassword(e.target.value)}
				id="filled-password-input"
				label="Password"
				type="password"
				autoComplete="current-password"
				variant="filled"
				data-cy="field_password"
			/>
			
			<Button
				color='secondary'
				variant="contained"
				onClick = {handleClick}
				data-cy="btn_login"
			>
				Sign In
			</Button>
			
			</div>
			
        </PageTemplate>
    );
}