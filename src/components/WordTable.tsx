import React from 'react';
import TableContainer from '@mui/material/TableContainer';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableRow from '@mui/material/TableRow';
import TableHead from '@mui/material/TableHead';
import Table from '@mui/material/Table';
import IconButton from '@mui/material/IconButton';
import VolumeUpIcon from '@mui/icons-material/VolumeUp';

import Paper from '@mui/material/Paper';


interface Props {
  words: {
    english: string;
    foreignword: string; 
    examplesentence_english: string;
    examplesentence_foreign: string;
    english_definition: string;
    foreign_definition: string;
    audiofilelink: string;
  }[],
  language: string,
};

const language_voice = new Map([
    ["english", "en-US"],
    ["spanish", "es-MX"],
    ["french", "fr-FR"],
    ["russian", "ru-RU"],
    ["italian", "it-IT"],
    ["japanese", "ja-JP"],
    ["chinese", "zh-TW"],
])

const canListen = 'speechSynthesis' in window;

const handleClick = (word: string, language: string) => {
    const synthesis = window.speechSynthesis;
    const text = new SpeechSynthesisUtterance(word);
    text.lang = language_voice.get(language) || "en-US";
    synthesis.speak(text);
}

function WordTable({words, language}: Props) {
    
    const upperCaseLanguage = language[0].toUpperCase() + language.substring(1);

    return (
        <>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell></TableCell>
                <TableCell data-cy="eng_header"><b>English</b></TableCell>
                <TableCell><b></b></TableCell>
                <TableCell data-cy="lang_header"><b data-testid="to-language">{upperCaseLanguage}</b></TableCell>
                <TableCell><b></b></TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {words.map((word, index) => (
                <TableRow key={index + 1} data-testid={"table_" + (index + 1) } data-cy={"table_" + (index + 1)}>
                    <TableCell>{index + 1}</TableCell>
                    <TableCell data-cy={"eng_word_"+(index+1)}>{word.english}</TableCell>
                    {canListen ? <TableCell>
                        <IconButton data-cy={"btn_speak_eng_"+(index+1)} color="primary" onClick={() => handleClick(word.english, "english")}>
                            <VolumeUpIcon />
                        </IconButton>
                    </TableCell> : <></>}
                    <TableCell data-cy={"lang_word_"+(index+1)}>{word.foreignword}</TableCell>
                    {canListen ? <TableCell>
                      <IconButton data-cy={"btn_speak_lang_"+(index+1)} color="primary"  onClick={() => handleClick(word.foreignword, language)}>
                          <VolumeUpIcon />
                      </IconButton>
                    </TableCell> : <></>}
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        </>
    );
}

export default React.memo(WordTable);
