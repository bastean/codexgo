Feature: Create a new user account

  Scenario: Create a valid non existing account through UI
    Given I am on the Home page
    * I fill the Email with uiCreate@example.com
    * I fill the Username with uiCreate
    * I fill the Password with uiCreate@example
    * I fill the Confirm Password with uiCreate@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Create a valid non existing account through API
    Given I send a PUT request to Create with body:
      """
      {
        "Email": "apiCreate@example.com",
        "Username": "apiCreate",
        "Password": "apiCreate@example",
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

  Scenario: Create already existing account through UI
    Given I am on the Home page
    * I fill the Email with uiCreate@example.com
    * I fill the Username with uiCreate
    * I fill the Password with uiCreate@example
    * I fill the Confirm Password with uiCreate@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Email already registered notification

  Scenario: Create already existing account through API
    Given I send a PUT request to Create with body:
      """
      {
        "Email": "apiCreate@example.com",
        "Username": "apiCreate",
        "Password": "apiCreate@example",
        "CaptchaID": "00000",
        "CaptchaAnswer": "00000"
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
            "Type": "AlreadyExist",
            "Message": "Email already registered",
            "Data": {
              "Field": "Email"
            }
          }
        ]
      }
      """
