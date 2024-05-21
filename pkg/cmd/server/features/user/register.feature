Feature: Create a new User account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I click the Create tab
    * I fill the Email with create@example.com
    * I fill the Username with create
    * I fill the Password with 12345678
    * I click the Create button
    And I see Successfully Created notification

  Scenario: Create already existing account
    Given I am on / page
    Then I click the Create tab
    * I fill the Email with create@example.com
    * I fill the Username with create
    * I fill the Password with 12345678
    * I click the Create button
    And I see Email already created notification
