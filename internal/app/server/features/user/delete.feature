Feature: Delete a user account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I fill the Email with delete@example.com
    * I fill the Username with delete
    * I fill the Password with delete@example
    * I fill the Confirm Password with delete@example
    * I check the I agree to the terms and conditions
    * I click the Sign up button
    And I see Account created notification

  Scenario: Login a valid existing account
    Given I am on / page
    Then I click on the Sign in button
    * I fill the Email with delete@example.com
    * I fill the Password with delete@example
    * I click the Sign in button
    * I see Logged in notification
    And redirect me to /dashboard page

  Scenario: Delete a valid existing account
    Given I am on /dashboard page
    Then I open the account menu
    * I click on the Delete button
    * I fill the Password with delete@example
    * I fill the Confirm Password with delete@example
    * I click the Approve button
    * I see Account deleted notification
    And redirect me to / page
