Feature: Delete a user account

  Scenario: Create a valid non existing account through UI
    Given I am on the Home page
    * I fill the Email with uiDelete@example.com
    * I fill the Username with uiDelete
    * I fill the Password with uiDelete@example
    * I fill the Confirm Password with uiDelete@example
    * I check the I agree to the terms and conditions
    * I hover the Sign up button
    * I fill the Answer with 00000
    When I click the Sign up button
    Then I see Account created notification

  Scenario: Create a valid non existing account through API
    Given I send a PUT request to Create with body:
      """
      {
        "Email": "apiDelete@example.com",
        "Username": "apiDelete",
        "Password": "apiDelete@example",
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
    * I fill the Email/Username with uiDelete@example.com
    * I fill the Password with uiDelete@example
    When I click the Sign in button
    Then I see Logged in notification
    And I get redirected to the Dashboard page

  Scenario: Login a valid existing account through API
    Given I send a POST request to Login with body:
      """
      {
        "Email": "apiDelete@example.com",
        "Password": "apiDelete@example"
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

  Scenario: Delete a valid existing account through UI
    Given I am on the Dashboard page
    * I open the account menu
    * I click on the Delete button
    * I fill the Password with uiDelete@example
    * I fill the Confirm Password with uiDelete@example
    When I click the Approve button
    Then I see Account deleted notification
    And I get redirected to the Home page

  Scenario: Delete a valid existing account through API
    Given I send a DELETE request to Delete with body:
      """
      {
        "Password": "apiDelete@example"
      }
      """
    Then the response status code should be 200
    And the response body should be:
      """
      {
        "Success": true,
        "Message": "Account deleted",
        "Data": null
      }
      """
