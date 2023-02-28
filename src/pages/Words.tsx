import React, {useState, useEffect} from 'react';
import PageTemplate from '../pages/PageTemplate';
import WordTable from '../components/WordTable';
import { FlashcardArray } from "react-quizlet-flashcard";

import Grid from '@mui/material/Grid';

import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import FormControl from '@mui/material/FormControl';
import Paper from '@mui/material/Paper';
import Box from '@mui/material/Box';

function Words() {

    const [language, setLanguage] = useState<string>("spanish");
    const [isLoaded, setLoaded] = useState(false);
    const [error, setError] = useState(null);
    const [items, setItems] = useState<any>([]);
    
    useEffect(() => {
        const requestOptions = {
            method: "GET",
            mode: "no-cors",
        } 
        fetch("api/words/package/id/1")
        .then(res => res.json())
        .then(
            (result) => {
                setItems(result);
                setLoaded(true);
            },
            (error) => {
                setLoaded(true);
                setError(error);
            }
        )
    });
    
    const languages = [
        "spanish",
        "french",
        "russian",
        "italian",
        "japanese",
        "chinese"
    ];

    const handleChangeLanguage = (e: SelectChangeEvent) => {
        setLanguage(e.target.value);
    }

    const upperCaseLanguage = language[0].toUpperCase() + language.substring(1);

    const languageDropDown = languages.map((language, key) => <option key={key} value={language}>{language[0].toUpperCase() + language.substring(1)}</option>)


    const date = new Date(items.date);
        
    let flashcards = [];
    if (isLoaded){
        flashcards = items["tenwords"].map((words: any, index: any) => (
            {
                id: index,
                frontHTML: <Grid container sx={{ height: '100%' }} alignItems="center" justifyContent="center"><Grid item><h1>{words.english}</h1></Grid></Grid>,
                backHTML: <Grid container sx={{ height: '100%' }} alignItems="center" justifyContent="center"><Grid item><h1>{words.translation}</h1></Grid></Grid>
            }
            )  
        );
    }

        return (
            <PageTemplate>
                <Box sx={{mt: "30px", ml: "10%", width: "80%"}}>
                    <Paper sx={{p: "20px"}}>
                        <FormControl sx={{ m: 3, minWidth: 200 }}>
                            <InputLabel>Learning</InputLabel>
                            <Select
                                inputProps={{ "data-testid": "language-select" }}
                                native={true}
                                value={language}
                                label="Language"
                                onChange={handleChangeLanguage}>
                                {languageDropDown}
                            </Select>
                        </FormControl>
                        <h1 style={{textAlign: "center"}}>{date.toDateString()}</h1>
                        <h2 data-testid="language-subtitle" style={{textAlign: "center"}}>10 Words in {upperCaseLanguage}</h2>
                        <Grid
                            container
                            justifyContent="center"
                            alignItems="center"
                            spacing={5}>
                            <Grid item>
                                <Box sx={{width: "100%"}}>
                                    {isLoaded ? <FlashcardArray cards={flashcards}/> : <h1>Loading...</h1>}
                                </Box>
                            </Grid>
                            <Grid item sx={{width: "80%"}}>
                                {isLoaded ? <WordTable words={items["tenwords"]} language={language}/> : <h1>Loading...</h1>}
                            </Grid>
                        </Grid>
                    </Paper>
                    
                </Box>          
            </PageTemplate>
        );
}

export default Words;
