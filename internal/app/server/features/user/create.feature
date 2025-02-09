Feature: Create a new user account

  Scenario: Create a valid non existing account
    Given I am on the Home page
    * I fill the Email with create@example.com
    * I fill the Username with create
    * I fill the Password with create@example
    * I fill the Confirm Password with create@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Create already existing account
    Given I am on the Home page
    * I fill the Email with create@example.com
    * I fill the Username with create
    * I fill the Password with create@example
    * I fill the Confirm Password with create@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000    
    When I click the Sign up button
    Then I see Email already registered notification
