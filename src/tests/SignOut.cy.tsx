import React from 'react'
import Welcome from '../pages/Welcome'
import { mount } from '@cypress/react';

it('renders', () => {
    cy.mount(<Welcome />);
})

it('button click', () => {
    window.localStorage.setItem('username', "Lindsey");
    window.localStorage.setItem('password', "password123");
    window.localStorage.setItem('date', "04-15-23");
    cy.mount(<Welcome />);
    cy.get('[data-cy="logout"]').click();
})
