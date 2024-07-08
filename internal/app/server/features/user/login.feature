Feature: Login a User account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I fill the Email with login@example.com
    * I fill the Username with login
    * I fill the Password with login@example
    * I fill the Confirm Password with login@example
    * I check the I agree to the terms and conditions
    * I click the Sign up button
    And I see account created notification

  Scenario: Login a valid existing account
    Given I am on / page
    Then I click on the Sign in button
    * I fill the Email with login@example.com
    * I fill the Password with login@example
    * I click the Sign in button
    * I see logged in notification
    And redirect me to /dashboard page

  Scenario: Login a valid non existing account
    Given I am on / page
    Then I click on the Sign in button
    * I fill the Email with non-existing@example.com
    * I fill the Password with non-existing@example
    * I click the Sign in button
    But I see not found notification
