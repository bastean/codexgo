Feature: Login a User account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I click the Create tab
    * I fill the Email with login@example.com
    * I fill the Username with login
    * I fill the Password with 12345678
    * I click the Create button
    And I see Successfully Created notification

  Scenario: Login a valid existing account
    Given I am on / page
    Then I fill the Email with login@example.com
    * I fill the Password with 12345678
    * I click the Login button
    And I am on /dashboard page

  Scenario: Login a valid non existing account
    Given I am on / page
    Then I fill the Email with non-existing@example.com
    * I fill the Password with non-existing
    * I click the Login button
    And I see Not Found non-existing@example.com notification
