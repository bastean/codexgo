Feature: Delete a user account

  Scenario: Create a valid non existing account
    Given I am on the Home page
    * I fill the Email with delete@example.com
    * I fill the Username with delete
    * I fill the Password with delete@example
    * I fill the Confirm Password with delete@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Login a valid existing account
    Given I am on the Home page
    * I click on the Sign in button
    * I fill the Email/Username with delete@example.com
    * I fill the Password with delete@example
    When I click the Sign in button
    Then I see Logged in notification
    And I get redirected to the Dashboard page

  Scenario: Delete a valid existing account
    Given I am on the Dashboard page
    * I open the account menu
    * I click on the Delete button
    * I fill the Password with delete@example
    * I fill the Confirm Password with delete@example
    When I click the Approve button
    Then I see Account deleted notification
    And I get redirected to the Home page
