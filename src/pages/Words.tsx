import React from 'react';
import PageTemplate from '../pages/PageTemplate';
import { FlashcardArray } from "react-quizlet-flashcard";

import Grid from '@mui/material/Grid';

export default function Words() {

    // Call to API to get words associated with ID and get the date
    const from_api = [
        {
            "original": "hello",
            "translation": "hola",
            "image": "https://upload.wikimedia.org/wikipedia/en/thumb/6/6b/Hello_Web_Series_%28Wordmark%29_Logo.png/1200px-Hello_Web_Series_%28Wordmark%29_Logo.png",
        },
        {
            "original": "goodbye",
            "translation": "adios"
        },
        {
            "original": "friend",
            "translation": "amigo"
        },
        {
            "original": "table",
            "translation": "la mesa",
        },
        {
            "original": "love",
            "translation": "amor",
        },
        {
            "original": "happy",
            "translation": "feliz",
        },
        {
            "original": "sad",
            "translation": "triste",
        },
        {
            "original": "mother",
            "translation": "madre",
        },
        {
            "original": "father",
            "translation": "padre",
        },
        {
            "original": "child",
            "translation": "nino"
        }
    ]

    const flashcards = from_api.map((words, index) => (
            {
                id: index,
                frontHTML: <Grid container sx={{ height: '100%' }} alignItems="center" justifyContent="center"><Grid item><h1>{words.translation}</h1>{words.image ? <img width={200} src={words.image}></img> : <></>}</Grid></Grid>,
                backHTML: <Grid container sx={{ height: '100%' }} alignItems="center" justifyContent="center"><Grid item><h1>{words.original}</h1></Grid></Grid>
            }
        )
    )

    return (
        <PageTemplate>
            <h1 style={{textAlign: "center"}}>Your Words</h1>
                <FlashcardArray cards={flashcards} />
        </PageTemplate>
    );
}