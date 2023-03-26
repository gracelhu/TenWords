import React from 'react';
import ReactDOM from 'react-dom';
import {fireEvent, render, screen} from '@testing-library/react'
import Practice from '../pages/Practice';
import { BrowserRouter } from 'react-router-dom';


test('Practice correctly renders title', () => {
    const practicePage = render(<BrowserRouter><Practice /></BrowserRouter>);
    const heading = practicePage.getByTestId('heading');
    expect(heading).toHaveTextContent("Test Your Knowledge!");
});

test('Practice correctly renders title', () => {
    const practicePage = render(<BrowserRouter><Practice /></BrowserRouter>);
    const heading = practicePage.getByTestId('heading');
    expect(heading).toHaveTextContent("Test Your Knowledge!");
});