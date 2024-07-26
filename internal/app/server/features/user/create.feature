Feature: Create a new user account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I fill the Email with create@example.com
    * I fill the Username with create
    * I fill the Password with create@example
    * I fill the Confirm Password with create@example
    * I check the I agree to the terms and conditions
    * I click the Sign up button
    And I see Account created notification

  Scenario: Create already existing account
    Given I am on / page
    Then I fill the Email with create@example.com
    * I fill the Username with create
    * I fill the Password with create@example
    * I fill the Confirm Password with create@example
    * I check the I agree to the terms and conditions
    * I click the Sign up button
    But I see Email already registered notification
