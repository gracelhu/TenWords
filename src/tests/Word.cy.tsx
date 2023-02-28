import React from 'react'
import Words from '../pages/Words'
import { mount } from '@cypress/react';
import { dummy_data } from '../pages/dummy_data';

it('test connection', () => {
    /*
    cy.request("http://localhost:8000/api/words/package/id/1").as("response");
    cy.get("@response").then((response) => {
        cy.mount(<Words />);
    });
    */
    cy.mount(<Words />);
})