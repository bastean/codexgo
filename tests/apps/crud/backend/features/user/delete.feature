Feature: Delete a User account

  Scenario: Register a valid non existing account
    Given I am on / page
    Then I click the Register tab
    * I fill the Email with delete@example.com
    * I fill the Username with delete
    * I fill the Password with 12345678
    * I click the Register button
    And I see Successfully Registered notification

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
