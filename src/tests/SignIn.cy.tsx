import React from 'react'
import SignIn from '../pages/SignIn'
import { mount } from '@cypress/react';

it('renders', () => {
    cy.mount(<SignIn />);
})

it('button click', () => {
    cy.mount(<SignIn />);
    cy.get('[data-cy="field_username"]').type("username");
    cy.get('[data-cy="field_password"]').type("password");
    cy.get('[data-cy="btn_login"]').click();
})