Feature: Login a user account

  Scenario: Create a valid non existing account through UI
    Given I am on the Home page
    * I fill the Email with uiLogin@example.com
    * I fill the Username with uiLogin
    * I fill the Password with uiLogin@example
    * I fill the Confirm Password with uiLogin@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Create a valid non existing account through API
    Given I send a PUT request to Create with body:
      """
      {
        "Email": "apiLogin@example.com",
        "Username": "apiLogin",
        "Password": "apiLogin@example",
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
    * I fill the Email/Username with uiLogin@example.com
    * I fill the Password with uiLogin@example
    When I click the Sign in button
    Then I see Logged in notification
    And I get redirected to the Dashboard page

  Scenario: Login a valid existing account through API
    Given I send a POST request to Login with body:
      """
      {
        "Email": "apiLogin@example.com",
        "Password": "apiLogin@example"
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

  Scenario: Login a valid non existing account through UI
    Given I am on the Home page
    * I click on the Sign in button
    * I fill the Email/Username with non-existing@example.com
    * I fill the Password with non-existing@example
    When I click the Sign in button
    Then I see non-existing@example.com not found notification

  Scenario: Login a valid non existing account through API
    Given I send a POST request to Login with body:
      """
      {
        "Email": "non-existing@example.com",
        "Password": "non-existing@example"
      }
      """
    Then the response status code should be 400
    And the response body should be:
      """
      {
        "Success": false,
        "Message": "Some errors have been found.",
        "Data": [
          {
            "Type": "NotExist",
            "Message": "non-existing@example.com not found",
            "Data": {
              "Index": "non-existing@example.com"
            }
          }
        ]
      }
      """
