import React from 'react';
import ReactDOM from 'react-dom';
import {fireEvent, render, screen} from '@testing-library/react'
import NavigationChild from '../components/NavigationChild';
import { BrowserRouter } from 'react-router-dom';


test('Home tab correctly displays', () => {
    const navigation = render(<NavigationChild logout={() => {}}/>);
    const tab = navigation.getByTestId('home');
    expect(tab).toHaveTextContent("Home");
});

test('Navigation tab correctly displays', () => {
    const navigation = render(<NavigationChild logout={() => {}}/>);
    const tab = navigation.getByTestId('your-words');
    expect(tab).toHaveTextContent("Your Words");
});
