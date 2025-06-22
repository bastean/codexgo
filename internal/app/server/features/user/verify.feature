Feature: Verify a user account

  Scenario: Create a valid non existing account through UI
    Given I am on the Home page
    * I fill the Email with uiVerify@example.com
    * I fill the Username with uiVerify
    * I fill the Password with uiVerify@example
    * I fill the Confirm Password with uiVerify@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Create a valid non existing account through API
    Given I send a PUT request to Create with body:
      """
      {
        "Email": "apiVerify@example.com",
        "Username": "apiVerify",
        "Password": "apiVerify@example",
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

  Scenario: Verify a valid existing account through UI
    Given I receive the link in uiVerify@example.com
    * I open the Verify link received
    * I fill the Password with uiVerify@example
    When I click the Verify button
    Then I see Account verified notification

  Scenario: Verify a valid existing account through API
    Given I receive the link in apiVerify@example.com
    When I send a PATCH request to Verify with body:
      """
      {
        "VerifyToken": "[Token]",
        "ID": "[ID]",
        "Password": "apiVerify@example"
      }
      """
    Then the response status code should be 200
    And the response body should be:
      """
      {
        "Success": true,
        "Message": "Account verified",
        "Data": null
      }
      """
