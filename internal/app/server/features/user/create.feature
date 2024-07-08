Feature: Create a new User account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I fill the Email with create@example.com
    * I fill the Username with create
    * I fill the Password with create@example
    * I fill the Confirm Password with create@example
    * I check the I agree to the terms and conditions
    * I click the Sign up button
    And I see account created notification

  Scenario: Create already existing account
    Given I am on / page
    Then I fill the Email with create@example.com
    * I fill the Username with create
    * I fill the Password with create@example
    * I fill the Confirm Password with create@example
    * I check the I agree to the terms and conditions
    * I click the Sign up button
    But I see already registered notification
