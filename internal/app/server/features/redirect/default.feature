Feature: Default Redirect

  Scenario: Check the correct redirect for not found page
    Given I am on the Undefined page
    Then I get redirected to the Home page
    And the page title should be codexGO
