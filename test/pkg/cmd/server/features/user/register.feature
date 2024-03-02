Feature: Register a new User account

  Scenario: Register a valid non existing account
    Given I am on / page
    Then I click the Register tab
    * I fill the Email with register@example.com
    * I fill the Username with register
    * I fill the Password with 12345678
    * I click the Register button
    And I see Successfully Registered notification

  Scenario: Register already existing account
    Given I am on / page
    Then I click the Register tab
    * I fill the Email with register@example.com
    * I fill the Username with register
    * I fill the Password with 12345678
    * I click the Register button
    And I see Email already registered notification
