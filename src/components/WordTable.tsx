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
    spanish: string; 
    french: string;
    russian: string;
    italian: string;
    japanese: string;
    chinese: string;
    examplesentence: string;
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

export default function WordTable({words, language}: Props) {
    
    const upperCaseLanguage = language[0].toUpperCase() + language.substring(1);

    return (
        <>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell></TableCell>
                <TableCell><b>English</b></TableCell>
                <TableCell><b></b></TableCell>
                <TableCell><b data-testid="to-language">{upperCaseLanguage}</b></TableCell>
                <TableCell><b></b></TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {words.map((word, index) => (
                <TableRow key={index + 1} data-testid={"table_" + (index + 1)}>
                    <TableCell>{index + 1}</TableCell>
                    <TableCell>{word.english}</TableCell>
                    {canListen ? <TableCell>
                        <IconButton color="primary" onClick={() => handleClick(word.english, "english")}>
                            <VolumeUpIcon />
                        </IconButton>
                    </TableCell> : <></>}
                    <TableCell>{word[language as keyof typeof word]}</TableCell>
                    {canListen ? <TableCell>
                      <IconButton color="primary"  onClick={() => handleClick(word[language as keyof typeof word], language)}>
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
