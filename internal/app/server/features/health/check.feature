Feature: Health Check

  Scenario: Check system status
    Given I send a GET request to Health
    Then the response status code should be 200
    And the response body should be:
      """
      {
        "Success": true,
        "Message": "OK",
        "Data": null
      }
      """
