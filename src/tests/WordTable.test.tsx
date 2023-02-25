import React from 'react';
import ReactDOM from 'react-dom';
import {fireEvent, render, screen} from '@testing-library/react'
import WordTable from '../components/WordTable';
import userEvent from '@testing-library/user-event';
import { dummy_data } from '../pages/dummy_data';

const from_api = dummy_data;

test("Ensure exactly 10 words show up in table", () => {
    const wordsTable = render(<WordTable words={from_api["tenwords"]} language="spanish"/>);
    expect(wordsTable.queryByTestId("table_10")).toBeTruthy();
    expect(!wordsTable.queryByTestId("table_11")).toBeTruthy();
});

test("Ensure words are in passed in language (Italian)", () => {
    const wordsTable = render(<WordTable words={from_api["tenwords"]} language="italian"/>);
    expect(wordsTable.getByTestId("to-language")).toHaveTextContent("Italian");
})

test("Ensure words are in passed in language (French)", () => {
    const wordsTable = render(<WordTable words={from_api["tenwords"]} language="french"/>);
    expect(wordsTable.getByTestId("to-language")).toHaveTextContent("French");
})