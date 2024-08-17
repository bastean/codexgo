Feature: Default Redirect

  Scenario: Check the correct redirect for not found page
    Given I am on /non-existing page
    Then redirect me to / page
    And the page title should be codexGO
