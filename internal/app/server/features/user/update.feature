Feature: Update a user account

  Scenario: Create a valid non existing account through UI
    Given I am on the Home page
    * I fill the Email with uiUpdate@example.com
    * I fill the Username with uiUpdate
    * I fill the Password with uiUpdate@example
    * I fill the Confirm Password with uiUpdate@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Create a valid non existing account through API
    Given I send a PUT request to Create with body:
      """
      {
        "Email": "apiUpdate@example.com",
        "Username": "apiUpdate",
        "Password": "apiUpdate@example",
        "CaptchaID": "00000",
        "CaptchaAnswer": "00000"
      }
      """
    Then the response status code should be 200
    And the response body should be:
      """
      {
        "Success": true,
        "Message": "Account created",
        "Data": null
      }
      """

  Scenario: Login a valid existing account through UI
    Given I am on the Home page
    * I click on the Sign in button
    * I fill the Email/Username with uiUpdate@example.com
    * I fill the Password with uiUpdate@example
    When I click the Sign in button
    Then I see Logged in notification
    And I get redirected to the Dashboard page

  Scenario: Login a valid existing account through API
    Given I send a POST request to Login with body:
      """
      {
        "Email": "apiUpdate@example.com",
        "Password": "apiUpdate@example"
      }
      """
    Then the response status code should be 200
    And the response body should be:
      """
      {
        "Success": true,
        "Message": "Logged in",
        "Data": null
      }
      """

  Scenario: Update a valid existing account through UI
    Given I am on the Dashboard page
    * I fill the Email with uiUpdated@example.com
    * I fill the Username with uiUpdated
    * I fill the New Password with uiUpdated@example
    * I fill the Confirm Password with uiUpdated@example
    * I fill the Current Password with uiUpdate@example
    When I click the Update button
    Then I see Account updated notification

  Scenario: Update a valid existing account through API
    Given I send a PATCH request to Update with body:
      """
      {
        "Email": "apiUpdated@example.com",
        "Username": "apiUpdated",
        "UpdatedPassword": "apiUpdated@example",
        "Password": "apiUpdate@example"
      }
      """
    Then the response status code should be 200
    And the response body should be:
      """
      {
        "Success": true,
        "Message": "Account updated",
        "Data": null
      }
      """
