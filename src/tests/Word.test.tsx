import React from 'react';
import ReactDOM from 'react-dom';
import {fireEvent, render, screen} from '@testing-library/react'
import Words from '../pages/Words';
import userEvent from '@testing-library/user-event';


// Testing Words Page

const supported_languages = [
  "spanish",
  "french",
  "russian",
  "italian",
  "japanese",
  "chinese"];

test('Change Language Selection Adjusts Subtitle', () => {
    const wordsPage = render(<Words />);
    for (let i = 0; i < supported_languages.length; i++){
      const language = supported_languages[i];
      const selectBox = wordsPage.getByTestId('language-select');;
      fireEvent.change(selectBox, {target: {"value": language}})
      const languageSubtitle = screen.getByTestId("language-subtitle");
      expect(languageSubtitle).toHaveTextContent(language[0].toUpperCase() + language.substring(1));
    } 
});

test('Change Language Selection Changes Words Table Component', () => {
  const wordsPage = render(<Words />);
  for (let i = 0; i < supported_languages.length; i++){
    const language = supported_languages[i];
    const selectBox = wordsPage.getByTestId('language-select');;
    fireEvent.change(selectBox, {target: {"value": language}})
    const table_to_language = screen.getByTestId("to-language");
    expect(table_to_language).toHaveTextContent(language[0].toUpperCase() + language.substring(1));
  } 
});