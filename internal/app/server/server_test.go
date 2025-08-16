package server_test

import (
	"context"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/playwright-community/playwright-go"
	testify "github.com/stretchr/testify/assert"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type Inbox struct {
	Messages []*struct {
		ID, Snippet string
		Links       []*struct {
			URL string
		}
	}
}

type Data struct {
	URL, Token, ID string
}

type Link struct {
	Verify, Reset *Data
}

const (
	apiVersion  = "v4"
	maxMessages = 2
)

var (
	sut  = os.Getenv("SUT_URL")
	api  = sut + "/" + apiVersion + "/"
	smtp = os.Getenv("SMTP_API")

	pw *playwright.Playwright
)

var (
	browser    playwright.Browser
	browserCtx playwright.BrowserContext
	page       playwright.Page
	element    playwright.Locator

	headless         = true
	timeout  float64 = 10000
	exact            = true
	sleep            = 4 * time.Second
)

var (
	request  playwright.APIRequestContext
	response playwright.APIResponse
)

var (
	inbox *Inbox
	link  *Link
)

var (
	err              error
	assert           *testify.Assertions
	expected, actual string
)

var (
	routes = map[string]string{
		"Home":      "/",
		"Verify":    "/verify",
		"Reset":     "/reset",
		"Dashboard": "/dashboard",
		"Undefined": "/undefined",
	}

	endpoints = map[string]string{
		"Health": "/health",
		"Create": "account",
		"Login":  "account",
		"Update": "account",
		"Delete": "account",
		"Verify": "account/verify",
		"Forgot": "account/forgot",
		"Reset":  "account/reset",
	}
)

func RefreshLinks() {
	assert.NotNil(inbox)

	link = new(Link)

	var (
		URL   *url.URL
		query url.Values
		data  *Data
	)

	for _, message := range inbox.Messages {
		response, err = request.Get(smtp + "/message/" + message.ID + "/link-check")

		assert.NoError(err)

		assert.NoError(response.JSON(message))

		for _, raw := range message.Links {
			if !strings.Contains(raw.URL, sut) {
				continue
			}

			URL, err = url.Parse(raw.URL)

			assert.NoError(err)

			query, err = url.ParseQuery(URL.RawQuery)

			assert.NoError(err)

			data = &Data{
				URL:   raw.URL,
				Token: query["token"][0],
				ID:    query["id"][0],
			}

			switch {
			case strings.Contains(message.Snippet, "confirm"):
				link.Verify = data
			case strings.Contains(message.Snippet, "reset"):
				link.Reset = data
			default:
				errors.Panic(errors.Standard("Unknown link %q with snippet %q", raw.URL, message.Snippet))
			}

			break
		}
	}
}

func RefreshInbox(email string) {
	response, err = request.Get(smtp + "/search?query=" + email)

	assert.NoError(err)

	inbox = new(Inbox)

	assert.NoError(response.JSON(inbox))

	assert.True(len(inbox.Messages) <= maxMessages)

	RefreshLinks()
}

func SetupAssert(t *testing.T) {
	assert = testify.New(t)
}

func SetupPlaywright() {
	pw, err = playwright.Run()

	if err != nil {
		errors.Panic(err)
	}

	browser, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: &headless,
		Timeout:  &timeout,
	})

	if err != nil {
		errors.Panic(err)
	}

	browserCtx, err = browser.NewContext(playwright.BrowserNewContextOptions{BaseURL: &sut})

	if err != nil {
		errors.Panic(err)
	}

	page, err = browserCtx.NewPage()

	if err != nil {
		errors.Panic(err)
	}

	request, err = pw.Request.NewContext(playwright.APIRequestNewContextOptions{BaseURL: &api})

	if err != nil {
		errors.Panic(err)
	}
}

func SetupBrowser(sc *godog.ScenarioContext) {
	sc.Given(`^I am on the (.+) page$`, func(route string) {
		_, err = page.Goto(routes[route])
		assert.NoError(err)
	})

	sc.Given(`^I open the (.+) link received$`, func(url string) {
		switch url {
		case "Verify":
			_, err = page.Goto(link.Verify.URL)
		case "Reset":
			_, err = page.Goto(link.Reset.URL)
		default:
			errors.Panic(errors.Standard("Unknown link %q", url))
		}

		assert.NoError(err)
	})

	sc.Given(`^I fill the (.+) with (.+)$`, func(placeholder, value string) {
		element = page.GetByRole("textbox", playwright.PageGetByRoleOptions{Name: placeholder, Exact: &exact}).Last()
		assert.NoError(element.Click())
		assert.NoError(element.Fill(value))
	})

	sc.Given(`^I check the (.+)$`, func(name string) {
		element = page.GetByRole("checkbox")
		assert.NoError(element.Check())
	})

	sc.Given(`^I click on the (.+) button$`, func(name string) {
		element = page.GetByText(name, playwright.PageGetByTextOptions{Exact: &exact}).First()
		assert.NoError(element.Click())
	})

	sc.Given(`^I open the (.+) menu$`, func(name string) {
		element = page.GetByRole("heading").First()
		assert.NoError(element.Click())
	})

	sc.Given(`^I hover the (.+) button$`, func(name string) {
		element = page.GetByRole("button", playwright.PageGetByRoleOptions{Name: name})
		assert.NoError(element.Hover())
	})

	sc.When(`^I click the (.+) button$`, func(name string) {
		element = page.GetByRole("button", playwright.PageGetByRoleOptions{Name: name})
		assert.NoError(element.Click())
	})

	sc.Then(`^I get redirected to the (.+) page$`, func(actual string) {
		time.Sleep(sleep)

		expected = page.URL()

		assert.True(strings.Contains(expected, routes[actual]))
	})

	sc.Then(`^the page title should be (.+)$`, func(expected string) {
		actual, err = page.Title()

		assert.NoError(err)

		assert.Equal(expected, actual)
	})

	sc.Then(`^I see (.+) notification$`, func(expected string) {
		element = page.GetByRole("alert")

		actual, err = element.InnerText()

		assert.NoError(err)

		assert.Equal(expected, actual)
	})
}

func SetupAPI(sc *godog.ScenarioContext) {
	sc.Given(`^I send a HEAD request to (.+)$`, func(endpoint string) {
		response, err = request.Head(endpoints[endpoint])
		assert.NoError(err)
	})

	sc.Given(`^I send a GET request to (.+)$`, func(endpoint string) {
		response, err = request.Get(endpoints[endpoint])
		assert.NoError(err)
	})

	sc.Given(`^I send a POST request to (.+) with body:$`, func(endpoint, payload string) {
		response, err = request.Post(endpoints[endpoint], playwright.APIRequestContextPostOptions{Data: payload})
		assert.NoError(err)
	})

	sc.Given(`^I send a PUT request to (.+) with body:$`, func(endpoint, payload string) {
		response, err = request.Put(endpoints[endpoint], playwright.APIRequestContextPutOptions{Data: payload})
		assert.NoError(err)
	})

	sc.Given(`^I send a PATCH request to (.+) with body:$`, func(endpoint, payload string) {
		response, err = request.Patch(endpoints[endpoint], playwright.APIRequestContextPatchOptions{Data: payload})
		assert.NoError(err)
	})

	sc.Given(`^I send a DELETE request to (.+) with body:$`, func(endpoint, payload string) {
		response, err = request.Delete(endpoints[endpoint], playwright.APIRequestContextDeleteOptions{Data: payload})
		assert.NoError(err)
	})

	sc.When(`^I send a PATCH request to (.+) with body:$`, func(endpoint, payload string) {
		var token, id string

		switch endpoint {
		case "Verify":
			token = link.Verify.Token
			id = link.Verify.ID
		case "Reset":
			token = link.Reset.Token
			id = link.Reset.ID
		default:
			errors.Panic(errors.Standard("Unknown endpoint %q", endpoint))
		}

		payload = strings.ReplaceAll(payload, "[Token]", token)

		payload = strings.ReplaceAll(payload, "[ID]", id)

		response, err = request.Patch(endpoints[endpoint], playwright.APIRequestContextPatchOptions{Data: payload})

		assert.NoError(err)
	})

	sc.Then(`^the response status code should be (.+)$`, func(expected int) {
		assert.Equal(expected, response.Status())
	})

	sc.Then(`^the response body should be empty$`, func() {
		actual, err = response.Text()

		assert.NoError(err)

		expected = ""

		assert.Equal(expected, actual)
	})

	sc.Then(`^the response body should be:$`, func(expected string) {
		actual, err = response.Text()

		assert.NoError(err)

		assert.JSONEq(expected, actual)
	})
}

func SetupScenario(sc *godog.ScenarioContext) {
	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		//? browserCtx.ClearCookies()
		return ctx, nil
	})

	sc.Given(`^I receive the link in (.+)$`, func(email string) {
		RefreshInbox(email)
	})

	SetupBrowser(sc)

	SetupAPI(sc)
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
		errors.Panic(errors.Standard("Failure to run feature tests resulted in a non-zero status [%d]", status))
	}
}
