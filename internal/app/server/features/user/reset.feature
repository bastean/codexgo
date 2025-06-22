Feature: Reset a user password

  Scenario: Create a valid non existing account through UI
    Given I am on the Home page
    * I fill the Email with uiReset@example.com
    * I fill the Username with uiReset
    * I fill the Password with uiReset@example
    * I fill the Confirm Password with uiReset@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Create a valid non existing account through API
    Given I send a PUT request to Create with body:
      """
      {
        "Email": "apiReset@example.com",
        "Username": "apiReset",
        "Password": "apiReset@example",
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

  Scenario: Recover a valid existing account through UI
    Given I am on the Home page
    * I click on the Sign in button
    * I click on the Forgot Password? button
    * I fill the Email with uiReset@example.com
    * I fill the Answer with 00000
    When I click the Send button
    Then I see Link sent. Please check your inbox notification

  Scenario: Recover a valid existing account through API
    Given I send a POST request to Forgot with body:
      """
      {
        "Email": "apiReset@example.com",
        "CaptchaID": "00000",
        "CaptchaAnswer": "00000"
      }
      """
    Then the response status code should be 200
    And the response body should be:
      """
      {
        "Success": true,
        "Message": "Link sent. Please check your inbox",
        "Data": null
      }
      """

  Scenario: Reset a valid existing account through UI
    Given I receive the link in uiReset@example.com
    * I open the Reset link received
    * I fill the Password with uiReset@example
    * I fill the Confirm Password with uiReset@example
    When I click the Reset button
    Then I see Password updated notification

  Scenario: Reset a valid existing account through API
    Given I receive the link in apiReset@example.com
    When I send a PATCH request to Reset with body:
      """
      {
        "ResetToken": "[Token]",
        "ID": "[ID]",
        "Password": "apiReset@example"
      }
      """
    Then the response status code should be 200
    And the response body should be:
      """
      {
        "Success": true,
        "Message": "Password updated",
        "Data": null
      }
      """
