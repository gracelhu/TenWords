import React, {useState} from 'react';
import logo from './logo.svg';
import './App.css';


import { ThemeProvider } from "@mui/material";
import { Routes, Route } from "react-router-dom";
import { theme } from "./theme";

import SignIn from './pages/SignIn';
import SignUp from './pages/SignUp';
import Words from './pages/Words';
import Welcome from './pages/Welcome';
import Practice from './pages/Practice'


function App() {
  const [loggedIn, setLoggedIn] = useState(true);
  return (
    <ThemeProvider theme={ theme }>
      <Routes>
        <Route path="/" element={<Welcome/>}></Route>
        <Route path="/sign-in" element={<SignIn/>}></Route>
        <Route path="/sign-up" element={<SignUp/>}></Route>
        <Route path="/words" element={<Words/>}></Route>
        <Route path="/practice" element={<Practice/>}></Route>
      </Routes>
      </ThemeProvider>
  );
}

export default App;
