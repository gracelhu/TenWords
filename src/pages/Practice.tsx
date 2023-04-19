import React, { useState } from 'react';
import PageTemplate from '../pages/PageTemplate';
import TextField from '@mui/material/TextField';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import { useLocation } from 'react-router-dom';

export default function Practice() {

    const location = useLocation();
    let from_api = (location.state !== null) ? location.state.words : {"tenwords": [], "date": ""};
    const language = (location.state !== null) ? location.state.language : "";

    const [text, setText] = useState(["", "", "", "", "", "", "", "", "", ""]);
    const [correct, setCorrect] = useState(0);
    const [clicked, setClicked] = useState(false);
    const [incorrect, setIncorrect] = useState([0]);

    const changeWord = (word: string, i: number) => {
        let myWords = text;
        myWords[i] = word;
        setText(myWords);
    }

    
    let testingWords = from_api["tenwords"].map((word: any, i: number) =>
        <>
        <p style = {{"color": clicked ? (incorrect.includes(i) ? "red" : "green") : "black"}}><b>{i + 1}. {word["english"]}</b></p>
        <TextField variant="standard" placeholder={word["english"] + " (translation)"} onChange={e => changeWord(e.target.value, i)}></TextField>
        </>
    );

    const trackQuizProgress = async(username : string, quiz : string, questioncount : number) => {
        console.log("Bla bla ablabla abllabla");
		fetch("http://localhost:3000/quiz/" + username + "/" + quiz + "/" + questioncount);
	}

    const checkWords = () => {
        setClicked(true);
        setIncorrect([]);
        let wrong: number[] = [];
        let correct = 0;
        for (let i = 0; i < from_api["tenwords"].length; i++){
            if (text[i] == from_api["tenwords"][i]["foreignword"]){
                correct = correct + 1;
                setCorrect(correct);
            } else {
                wrong.push(i);
                setIncorrect(wrong);
            }
        }
        if (correct >= 1) {
            const getusername = localStorage.getItem("username");
            if (getusername != null) {
                const upsertusername = JSON.parse(getusername);
                trackQuizProgress(upsertusername, from_api["date"], correct);
            } else {
                console.log("username or password is null");
            }
        }
    }

    
    return (
        <PageTemplate>
            <Card sx={{maxWidth: 500, m: 2, p: 10}} style={{textAlign: "center"}}>
                    <h1 data-testid="heading">Test Your Knowledge!</h1>
                    <p data-testid="to-language"><span data-testid="subtitle">Test your knowledge of the following words by taking the quiz for the word packet:</span> {from_api["date"]}.</p>
                    <CardContent>
                        <ol>
                            {testingWords}
                        </ol>
                        {!clicked ? <p></p> : <h2 data-testid="score">Your Score: {correct}/10</h2>}
                        <Button variant="contained" onClick={() => checkWords()}>Check</Button>
                    </CardContent>
            </Card>
        </PageTemplate>
    );
}