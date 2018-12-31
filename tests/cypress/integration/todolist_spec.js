describe('Todo tests', () => {
  
  const baseUrl = Cypress.env('FRONTEND_URL')

  beforeEach(() => {
    cy.exec('sh resetdb.sh')
  })

  context('When page is opened', () => {
    it('Creates a new task and add it in top of the list', () => {
      const task = 'todo test 1 cypress'
      cy.visit(baseUrl)
        .get('input[type="text"]').clear().type(task)
        .get('button.button').click()
      
      cy.get('ul > li:first > .tasks__item__toggle')
        .should(($item) => {
          expect($item).to.contain(task)
        })
    })

    it('Clicks on the task to mark as complete and hide it from the list', () => {
      const selector = 'ul > li:first > .tasks__item__toggle'
      cy.visit(baseUrl)
        .get(selector)
        .then(($btn) => {
          const task = $btn.text()
          $btn.click()

          cy.get(selector).should('have.class', 'tasks__item__toggle--done')
        })
    })
  })

})
