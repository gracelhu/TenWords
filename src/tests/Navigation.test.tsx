import React from 'react';
import ReactDOM from 'react-dom';
import {fireEvent, render, screen} from '@testing-library/react'
import Navigation from '../components/Navigation';
import { BrowserRouter } from 'react-router-dom';


test('Home tab correctly displays', () => {
    const navigation = render(<Navigation/>);
    const tab = navigation.getByTestId('home');
    expect(tab).toHaveTextContent("Home");
});

test('Navigation tab correctly displays', () => {
    const navigation = render(<Navigation/>);
    const tab = navigation.getByTestId('your-words');
    expect(tab).toHaveTextContent("Your Words");
});
