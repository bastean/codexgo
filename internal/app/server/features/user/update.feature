Feature: Update a user account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I fill the Email with update@example.com
    * I fill the Username with update
    * I fill the Password with update@example
    * I fill the Confirm Password with update@example
    * I check the I agree to the terms and conditions
    * I click the Sign up button
    And I see Account created notification

  Scenario: Login a valid existing account
    Given I am on / page
    Then I click on the Sign in button
    * I fill the Email with update@example.com
    * I fill the Password with update@example
    * I click the Sign in button
    * I see Logged in notification
    And redirect me to /dashboard page

  Scenario: Update a valid existing account
    Given I am on /dashboard page
    Then I fill the Email with updated@example.com
    * I fill the Username with updated
    * I fill the New Password with updated@example
    * I fill the Confirm Password with updated@example
    * I fill the Current Password with update@example
    * I click the Update button
    And I see Account updated notification
