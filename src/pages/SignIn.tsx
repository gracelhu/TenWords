import React from 'react';
import PageTemplate from '../pages/PageTemplate';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button'

import { BrowserRouter } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import { useRef, useState, useEffect } from 'react';

export default function SignIn() {

	//let navigate = useNavigate();
	const sleep = (milliseconds: number) => {
        return new Promise(resolve => setTimeout(resolve, milliseconds))
    }

	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [message, setMessage] = useState("");
    const [messageColor, setMessageColor] = useState("black");

	const handleClick = async() => {
		fetch("http://localhost:3000/auth/"+username+"/"+password)
        .then(res => res.json())
        .then(
            (result) => {
                console.log("Trying to login with " + username + ", " + password);
				console.log(result.State);
                switch (result.State) {
                    case 'invalid':
                        setMessage("Error! Incorrect password!"); setMessageColor("red"); break;
                    case 'returning':
                        setMessage("Login successful! Redirecting..."); setMessageColor("black");
						//sleep(1000).then(r => {
							//history.replaceState(null, '', '/words');
							//window.location.replace('/words');
							//navigate('/words');
						//})
						let windowRef = window
						setTimeout(() => {
							windowRef.location.replace('http://localhost:3000/words')
						})
						break;
                    case 'register':
                        setMessage("Error! Username does not exist!"); setMessageColor("red"); break;
                    default:
                        setMessageColor("black");
                }
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

			<br></br>
            <h4 style={{color: messageColor, textAlign: "center"}} data-cy="auth_message">{message}</h4>
			
        </PageTemplate>
    );
}