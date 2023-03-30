describe('template spec', () => {
  it('login', () => {
    cy.visit('http://localhost:3000/sign-in')

    cy.get('[data-cy="field_username"]').type("username");
    cy.get('[data-cy="field_password"]').type("password");
    cy.get('[data-cy="btn_login"]').click();
    cy.wait(1000)
    //cy.log(cy.url()["specWindow"])
    cy.url().should('include', '/words')
  })
})