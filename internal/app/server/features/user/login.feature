Feature: Login a user account

  Scenario: Create a valid non existing account
    Given I am on the Home page
    * I fill the Email with login@example.com
    * I fill the Username with login
    * I fill the Password with login@example
    * I fill the Confirm Password with login@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Login a valid existing account
    Given I am on the Home page
    * I click on the Sign in button
    * I fill the Email/Username with login@example.com
    * I fill the Password with login@example
    When I click the Sign in button
    Then I see Logged in notification
    And I get redirected to the Dashboard page

  Scenario: Login a valid non existing account
    Given I am on the Home page
    * I click on the Sign in button
    * I fill the Email/Username with non-existing@example.com
    * I fill the Password with non-existing@example
    When I click the Sign in button
    Then I see non-existing@example.com not found notification
