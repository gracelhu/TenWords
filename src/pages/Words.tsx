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

    const language_code: { [key: string]: string } = {
        "spanish": "es",
        "french": "fr",
        "russian": "ru",
        "italian": "it",
        "japanese": "ja",
        "chinese": "zh-cn"
    };
    const [language, setLanguage] = useState<string>("spanish");
    const [isLoaded, setLoaded] = useState(false);
    const [items, setItems] = useState<any>([]);
    
    useEffect(() => {
        fetch("/api/words/"+language_code[language as keyof typeof language_code]+"/package/1")
        .then(res => res.json())
        .then(
            (result) => {
                setItems(result);
                setLoaded(true);
                console.log("this is frontend.")
            },
        )
        .catch(error => {
            console.log("Fetch error");
            console.warn(error)
        })
    }, []);
    
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

    const upperCaseLanguage = language[0].toUpperCase() + language.substring(1) || "";

    const languageDropDown = languages.map((value, key) => <option key={key} value={value}>{value[0].toUpperCase() + value.substring(1)}</option>)


    const date = new Date(items["date"]);
    let flashcards = [];
    let words_for_table = [];
    
    if (isLoaded){
        words_for_table = items["tenwords"];
        flashcards = items["tenwords"].map((words: any, index: any) => ({
            id: index,
            frontHTML: <Grid container direction="column" sx={{ height: '100%' }} alignItems="center" justifyContent="center">
                <Grid item>
                    <Box width="100%"><h1>{words.english}</h1></Box>
                </Grid>
                <Grid item>
                    <Box width="100%"><p>{words.examplesentence_english}</p></Box>
                </Grid>
            </Grid>,
            backHTML: <Grid container sx={{ height: '100%' }} direction="column"  alignItems="center" justifyContent="center">
                        <Grid item>
                        <Box width="100%"><h1>{words.foreignword}</h1></Box>
                        </Grid>
                        <Grid item>
                        <Box width="100%"><p>{words.examplesentence_foreign}</p></Box>
                        </Grid>
                    </Grid>
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
                                label="Language"
                                onChange={handleChangeLanguage}>
                                {languageDropDown}
                            </Select>
                        </FormControl>
                        <h1 style={{textAlign: "center"}}>Packet 1: {items["date"]}</h1>
                        <h2 data-testid="language-subtitle" style={{textAlign: "center"}}>10 Words in {upperCaseLanguage}</h2>
                        <Grid
                            container
                            justifyContent="center"
                            alignItems="center"
                            spacing={5}>
                            <Grid item>
                                <Box sx={{width: "100%"}}>
                                    {isLoaded ? <FlashcardArray cards={flashcards}/> : <></>}
                                </Box>
                            </Grid>
                            <Grid item sx={{width: "80%"}}>
                                <WordTable words={words_for_table} language={language}/>
                            </Grid>
                        </Grid>
                    </Paper>
                    
                </Box>          
            </PageTemplate>
        );

}

export default React.memo(Words);
