Feature: Update a user account

  Scenario: Create a valid non existing account
    Given I am on the Home page
    * I fill the Email with update@example.com
    * I fill the Username with update
    * I fill the Password with update@example
    * I fill the Confirm Password with update@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Login a valid existing account
    Given I am on the Home page
    * I click on the Sign in button
    * I fill the Email/Username with update@example.com
    * I fill the Password with update@example
    When I click the Sign in button
    Then I see Logged in notification
    And I get redirected to the Dashboard page

  Scenario: Update a valid existing account
    Given I am on the Dashboard page
    * I fill the Email with updated@example.com
    * I fill the Username with updated
    * I fill the New Password with updated@example
    * I fill the Confirm Password with updated@example
    * I fill the Current Password with update@example
    When I click the Update button
    Then I see Account updated notification
