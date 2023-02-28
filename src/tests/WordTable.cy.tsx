import React from 'react'
import WordTable from '../components/WordTable'
import { mount } from '@cypress/react';
import { dummy_data } from '../pages/dummy_data';

it('renders', () => {
  cy.mount(<WordTable language="spanish" words={dummy_data["tenwords"]} />);
})

it('spanish translation', () => {
  cy.mount(<WordTable language="spanish" words={dummy_data["tenwords"]} />);
  cy.get('[data-cy="lang_header"]').should('have.text', 'Spanish');
  cy.get('[data-cy="eng_word_1"]').should('have.text', 'reveal');
  cy.get('[data-cy="lang_word_1"]').should('have.text', 'revelar');
})

it('button click', () => {
  cy.mount(<WordTable language="spanish" words={dummy_data["tenwords"]} />);
  cy.get('[data-cy="btn_speak_eng_1"]').click();
})