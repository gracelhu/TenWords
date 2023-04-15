import React, {useState, useEffect} from 'react';
import PageTemplate from '../pages/PageTemplate';
import WordTable from '../components/WordTable';
import { FlashcardArray } from "react-quizlet-flashcard";
import {Link} from 'react-router-dom';
import Grid from '@mui/material/Grid';

import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import Button from '@mui/material/Button';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import FormControl from '@mui/material/FormControl';
import Paper from '@mui/material/Paper';
import Box from '@mui/material/Box';
import ArrowCircleRightIcon from '@mui/icons-material/ArrowCircleRight';
import ArrowCircleLeftIcon from '@mui/icons-material/ArrowCircleLeft';
import { fontWeight } from '@mui/system';
import { DateTimePicker, DatePicker, LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import TextField from '@mui/material/TextField';

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
    const [progressIndex, setProgressIndex] = useState(1);
    const [packageNumber, setPackageNumber] = useState(1);
    const [packetDate, setPacketDate] = useState("4-10-2023");

    useEffect(() => {
        fetch("/api/words/"+language_code[language as keyof typeof language_code]+"/package/"+progressIndex)
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

        setPacketDate(items["date"]);

    }, [progressIndex]);

    //for the toggle arrows 
    const nextTenWordPackage = () => {
        if(progressIndex <= 3036)
        {
            setProgressIndex(progressIndex + 10);
            console.log("clicked right!");
            console.log(packageNumber);
            setPackageNumber(packageNumber + 1);
            console.log(packageNumber);
            //write code here to make the button momentarily change to gray when its clicked 
        }
    }

    const previousTenWordPackage = () => {
        if(progressIndex >= 11)
        {
            setProgressIndex(progressIndex - 10);
            console.log("clicked left!");
            setPackageNumber(packageNumber - 1);
            //write code here to make the button momentarily change to gray when its clicked 
        }
    }
    
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

    const boldWord = ( boldWord: string, sentence: string) => {
        let index = sentence.indexOf(boldWord);
        let listOfWords = sentence.split(boldWord);
        if (sentence.includes(boldWord)){
            if (listOfWords.length === 2){
                return <p>{listOfWords[0]}<span style={{textDecoration: "underlined", fontWeight: "bold"}}>{boldWord}</span>{listOfWords[1]}</p>;
            } else {
                if (index === 0){
                    return <p><span style={{textDecoration: "underlined", fontWeight: "bold"}}>{boldWord}</span>{listOfWords[0]}</p>;
                } else {
                    return <p>{listOfWords[0]}<span style={{textDecoration: "underlined", fontWeight: "bold"}}>{boldWord}</span></p>;
                }
            }
        } else {
            return <p>{listOfWords[0]}</p>;
        }
    }
    
    if (isLoaded){
        words_for_table = items["tenwords"];
        flashcards = items["tenwords"].map((words: any, index: any) => ({
            id: index,
            frontHTML: <Grid container direction="column" sx={{ height: '100%' }} alignItems="center" justifyContent="center">
                <Grid item>
                    <Box width="100%"><h1>{words.english}</h1></Box>
                </Grid>
                <Grid item>
                    <Box width="100%">
                        <Box width="100%"><p>{boldWord(words.english, words.examplesentence_english)}</p></Box>
                    </Box>
                </Grid>
            </Grid>,
            backHTML: <Grid container sx={{ height: '100%' }} direction="column"  alignItems="center" justifyContent="center">
                        <Grid item>
                        <Box width="100%"><h1>{words.foreignword}</h1></Box>
                        </Grid>
                        <Grid item>
                        <Box width="100%"><p>{boldWord(words.foreignword, words.examplesentence_foreign)}</p></Box>
                        </Grid>
                    </Grid>
        }
        )   
        );
    }

    const changeDate = (e: any) => {
        const date = (e.$M + 1).toString().padStart(2,'0') + "-" + e.$D.toString().padStart(2,'0') + "-" + e.$y;
        setPacketDate(date);
        fetch("/api/words/"+language_code[language as keyof typeof language_code]+"/package/date/"+date)
        .then(res => res.json())
        .then(
            (result) => {
                setItems(result);
                setLoaded(true);
            },
        )
        .catch(error => {
            console.warn(error)
        })
    }

        return (
            <PageTemplate>
                <Box textAlign='center'>
                <h3 data-cy="word_template"></h3>
                <h3 style={{textAlign: "center", color: "black"}}>Click arrows to toggle between different ten word packages</h3>
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <DatePicker
                        label="Select Packet Date"
                        value={packetDate}
                        onChange={(e) => changeDate(e)}
                        renderInput={(props) => (
                            <TextField {...props} />
                          )}
                    />
                </LocalizationProvider>
                <br/>
                <br/>
                <ArrowCircleLeftIcon style={{transform: "scale(2)", color: "black", marginRight: "32px" }} onClick={previousTenWordPackage}></ArrowCircleLeftIcon>
                <ArrowCircleRightIcon style={{transform: "scale(2)", color: "black", marginRight: "32px" }} onClick={nextTenWordPackage}></ArrowCircleRightIcon>
                </Box>
                <Box sx={{mt: "30px", ml: "10%", width: "80%"}}>
                    <Paper sx={{p: "20px"}}>
                        <Link
                            style={{ textDecoration: 'none' }}
                            to="/practice"
                            state={
                                {"words": items, "language": language}
                            }>
                                <Button variant="contained">Take a Quiz</Button>
                        </Link>
                        <br/>
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
                        <h1 style={{textAlign: "center"}}>{'Packet ' + packageNumber + ': '} {items["date"]}</h1>
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

