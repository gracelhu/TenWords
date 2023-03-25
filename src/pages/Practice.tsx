import React, { useState } from 'react';
import PageTemplate from '../pages/PageTemplate';
import TextField from '@mui/material/TextField';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import { useLocation } from 'react-router-dom';


export default function Practice() {

    let from_api = { "tenwords": [ { "id": "1", "english": "abandon", "foreignword": "abandonar", "examplesentence_english": "Many baby girls have been abandoned on the str eets of Beijing.", "examplesentence_foreign": "Muchas niñas han sido abandonadas en las calles de Beijing.", "english_definition": "To give up or relinquish control of, to surrender or to give oneself over, or to yield to one's emotions.", "foreign_definition": "Renunciar o renunciar al control de, rendirse o entregarse, o ceder a las propias emociones.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/abandon-us.mp3" }, { "id": "2", "english": "sudden", "foreignword": "repentino", "examplesentence_english": "The sudden drop in temperature left everyone cold and confused.", "examplesentence_foreign": "La repentina caída de la temperatura dejó a todos helados y confundidos.", "english_definition": "An unexpected occurrence; a surprise.", "foreign_definition": "Una ocurrencia inesperada; una sorpresa.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/sudden-us.mp3" }, { "id": "3", "english": "lawyer", "foreignword": "abogado", "examplesentence_english": "A lawyer's time and advice are his stock in trade. - aphorism often credited to Abraham Lincoln, but without attestation", "examplesentence_foreign": "El tiempo y el consejo de un abogado son su valor en el comercio. - aforismo a menudo acreditado a Abraham Lincoln, pero sin atestación", "english_definition": "A professional person qualified (as by a law degree or bar exam) and authorized to practice law, i.e. represent parties in lawsuits or trials and give legal advice.", "foreign_definition": "Una persona profesional calificada (por un título en derecho o un examen de la barra) y autorizada para ejercer la abogacía, es decir, representar a las partes en demandas o juicios y brindar asesoramiento legal.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/lawyer-us.mp3" }, { "id": "4", "english": "particularly", "foreignword": "particularmente", "examplesentence_english": "The apéritifs were particularly stimulating.", "examplesentence_foreign": "Los aperitivos fueron particularmente estimulantes.", "english_definition": "(focus) Especially, extremely.", "foreign_definition": "(enfoque) Especialmente, extremadamente.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/particularly-us.mp3" }, { "id": "5", "english": "gender", "foreignword": "género", "examplesentence_english": "The effect of the medication is dependent upon age, gender, and other factors.", "examplesentence_foreign": "El efecto del medicamento depende de la edad, el sexo y otros factores.", "english_definition": "Class; kind.", "foreign_definition": "Clase; amable.", "audiofilelink": "" }, { "id": "6", "english": "literary", "foreignword": "literario", "examplesentence_english": "a literary history", "examplesentence_foreign": "una historia literaria", "english_definition": "Relating to literature.", "foreign_definition": "Relativo a la literatura.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/literary-us.mp3" }, { "id": "7", "english": "cotton", "foreignword": "algodón", "examplesentence_english": "", "examplesentence_foreign": "", "english_definition": "Gossypium, a genus of plant used as a source of cotton fiber.", "foreign_definition": "Gossypium, un género de planta utilizado como fuente de fibra de algodón.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/cotton-1-us.mp3" }, { "id": "8", "english": "station", "foreignword": "estación", "examplesentence_english": "She had ambitions beyond her station.", "examplesentence_foreign": "Ella tenía ambiciones más allá de su posición.", "english_definition": "A stopping place.", "foreign_definition": "Un lugar de parada.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/station-au.mp3" }, { "id": "9", "english": "everyone", "foreignword": "todos", "examplesentence_english": "", "examplesentence_foreign": "", "english_definition": "Every person.", "foreign_definition": "Cada persona.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/everyone-us.mp3" }, { "id": "10", "english": "life", "foreignword": "vida", "examplesentence_english": "Having experienced both, the vampire decided that he preferred (un)death to life.  He gave up on life.", "examplesentence_foreign": "Habiendo experimentado ambos, el vampiro decidió que prefería la (des)muerte a la vida. Renunció a la vida.", "english_definition": "The state of organisms preceding their death, characterized by biological processes such as metabolism and reproduction and distinguishing them from inanimate objects; the state of being alive and living.", "foreign_definition": "El estado de los organismos que precede a su muerte, caracterizado por procesos biológicos como el metabolismo y la reproducción y que los distingue de los objetos inanimados; el estado de estar vivo y vivir.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/life-uk.mp3" } ], "date": "02-26-2023" }
    const language = useLocation().state.language;

    const [text, setText] = useState(["", "", "", "", "", "", "", "", "", ""]);
    const [correct, setCorrect] = useState(0);
    const [clicked, setClicked] = useState(false);
    const [incorrect, setIncorrect] = useState([0]);

    const changeWord = (word: string, i: number) => {
        let myWords = text;
        myWords[i] = word;
        setText(myWords);
    }

    
    let testingWords = from_api["tenwords"].map((word, i) =>
        <>
        <p style = {{"color": clicked ? (incorrect.includes(i) ? "red" : "green") : "black"}}><b>{i + 1}. {word["english"]}</b></p>
        <TextField variant="standard" placeholder={word["english"] + " (translation)"} onChange={e => changeWord(e.target.value, i)}></TextField>
        </>
    );

    const checkWords = () => {
        setClicked(true);
        setIncorrect([]);
        let wrong: number[] = [];
        for (let i = 0; i < from_api["tenwords"].length; i++){
            if (text[i] == from_api["tenwords"][i]["foreignword"]){
                setCorrect(correct + 1);
            } else {
                wrong.push(i);
                setIncorrect(wrong);
            }
        }
    }

    
    return (
        <PageTemplate>
            <Card sx={{maxWidth: 500, m: 2, p: 10}}>
                    <h1>Test Your Knowledge!</h1>
                    <p>Test your knowledge of the following words by taking the quiz for the word packet: {from_api["date"]}.</p>
                    <CardContent>
                        <ol>
                            {testingWords}
                        </ol>
                        {!clicked ? <p></p> : <h2>Your Score: {correct}/10</h2>}
                        <Button variant="contained" onClick={() => checkWords()}>Check</Button>
                    </CardContent>
            </Card>
        </PageTemplate>
    );
}