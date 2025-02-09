Feature: User forgot password

  Scenario: Create a valid non existing account
    Given I am on the Home page
    * I fill the Email with forgot@example.com
    * I fill the Username with forgot
    * I fill the Password with forgot@example
    * I fill the Confirm Password with forgot@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Recover a valid existing account
    Given I am on the Home page
    * I click on the Sign in button
    * I click on the Forgot Password? button
    * I fill the Email with forgot@example.com
    * I fill the Answer with 00000
    When I click the Send button
    Then I see Link sent. Please check your inbox notification

  Scenario: Recover a valid non existing account
    Given I am on the Home page
    * I click on the Sign in button
    * I click on the Forgot Password? button
    * I fill the Email with non-existing@example.com
    * I fill the Answer with 00000
    When I click the Send button
    Then I see non-existing@example.com not found notification
