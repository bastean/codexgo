package server_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/cucumber/godog"
	"github.com/playwright-community/playwright-go"
	testify "github.com/stretchr/testify/assert"
)

var (
	sut = os.Getenv("SUT_URL")
)

var (
	pw         *playwright.Playwright
	browser    playwright.Browser
	browserCtx playwright.BrowserContext
	page       playwright.Page
	element    playwright.Locator
)

var (
	headless         = true
	timeout  float64 = 10000
	exact            = true
	sleep            = 4 * time.Second
)

var (
	err              error
	assert           *testify.Assertions
	expected, actual string
)

func SetupAssert(t *testing.T) {
	assert = testify.New(t)
}

func SetupPlaywright() {
	pw, err = playwright.Run()

	if err != nil {
		errors.Panic(err.Error(), "SetupPlaywright")
	}

	browser, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: &headless,
		Timeout:  &timeout,
	})

	if err != nil {
		errors.Panic(err.Error(), "SetupPlaywright")
	}

	browserCtx, err = browser.NewContext(playwright.BrowserNewContextOptions{BaseURL: &sut})

	if err != nil {
		errors.Panic(err.Error(), "SetupPlaywright")
	}

	page, err = browserCtx.NewPage()

	if err != nil {
		errors.Panic(err.Error(), "SetupPlaywright")
	}
}

func SetupScenario(sc *godog.ScenarioContext) {
	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		//? browserCtx.ClearCookies()
		return ctx, nil
	})

	sc.Given(`^I am on (.+) page$`, func(route string) {
		_, err = page.Goto(route)
		assert.NoError(err)
	})

	sc.Then(`^redirect me to (.+) page$`, func(actual string) {
		time.Sleep(sleep)

		if actual == "/" {
			actual = ""
		}

		expected = page.URL()

		assert.True(strings.Contains(expected, actual))
	})

	sc.Then(`^the page title should be (.+)$`, func(expected string) {
		actual, err = page.Title()

		assert.NoError(err)

		assert.Equal(expected, actual)
	})

	sc.Then(`^I click on the (.+) button$`, func(name string) {
		element = page.GetByText(name, playwright.PageGetByTextOptions{Exact: &exact}).First()
		assert.NoError(element.Click())
	})

	sc.Then(`^I open the (.+) menu$`, func(name string) {
		element = page.GetByRole("heading").First()
		assert.NoError(element.Click())
	})

	sc.Then(`^I fill the (.+) with (.+)$`, func(placeholder, value string) {
		element = page.GetByRole("textbox", playwright.PageGetByRoleOptions{Name: placeholder, Exact: &exact}).Last()
		assert.NoError(element.Fill(value))
	})

	sc.Then(`^I click the (.+) button$`, func(name string) {
		element = page.GetByRole("button", playwright.PageGetByRoleOptions{Name: name})
		assert.NoError(element.Click())
	})

	sc.Then(`^I check the (.+)$`, func(name string) {
		element = page.GetByRole("checkbox")
		assert.NoError(element.Check())
	})

	sc.Then(`^I see (.+) notification$`, func(expected string) {
		element = page.GetByRole("alert")

		actual, err = element.InnerText()

		assert.NoError(err)

		assert.Equal(expected, actual)
	})
}

func TestAcceptanceServerFeatures(t *testing.T) {
	SetupAssert(t)

	SetupPlaywright()

	suite := godog.TestSuite{
		ScenarioInitializer: SetupScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if status := suite.Run(); status != 0 {
		errors.Panic(fmt.Sprintf("Failure to run feature tests resulted in a non-zero status [%d]", status), "TestAcceptanceServerFeatures")
	}
}
