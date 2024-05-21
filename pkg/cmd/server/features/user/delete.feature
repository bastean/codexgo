Feature: Delete a User account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I click the Create tab
    * I fill the Email with delete@example.com
    * I fill the Username with delete
    * I fill the Password with 12345678
    * I click the Create button
    And I see Successfully Created notification

  Scenario: Delete a valid existing account
    Given I am on / page
    Then I fill the Email with delete@example.com
    * I fill the Password with 12345678
    * I click the Login button
    * I am on /dashboard page
    * I click the Delete tab
    * I click the Delete button
    * I accept the delete confirm
    And I am on / page
