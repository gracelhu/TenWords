import React from 'react'
import SignUp from '../pages/SignUp'
import { mount } from '@cypress/react';

it('renders', () => {
    cy.mount(<SignUp />);
})

it('button click', () => {
    cy.mount(<SignUp />);
    cy.get('[data-cy="field_username"]').type("username");
    cy.get('[data-cy="field_password"]').type("password");
    cy.get('[data-cy="btn_login"]').click();
})