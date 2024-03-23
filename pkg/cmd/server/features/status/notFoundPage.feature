Feature: Not Found Redirect

  Scenario: Check the correct redirect for not found page
    Given I am on /non-existing page
    Then the page title should be codexgo
