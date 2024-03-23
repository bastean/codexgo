Feature: Update a User account

  Scenario: Register a valid non existing account
    Given I am on / page
    Then I click the Register tab
    * I fill the Email with update@example.com
    * I fill the Username with update
    * I fill the Password with 12345678
    * I click the Register button
    And I see Successfully Registered notification

  Scenario: Update a valid existing account
    Given I am on / page
    Then I fill the Email with update@example.com
    * I fill the Password with 12345678
    * I click the Login button
    * I am on /dashboard page
    * I fill the Email with updated@example.com
    * I fill the Username with updated
    * I fill the Current Password with 12345678
    * I fill the New Password with 87654321
    * I click the Update button
    And I see Successfully Updated notification
