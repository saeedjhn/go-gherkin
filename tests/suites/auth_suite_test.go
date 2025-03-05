//go:build acceptance

package suites_test

import (
	"gobdd/tests/steps/auth/register"
	"testing"

	"github.com/cucumber/godog"
)

//go:generate go test -tags=acceptance -v ./tests/...

func Test_Auth(t *testing.T) {
	rc := register.NewContext()

	suite := godog.TestSuite{
		// ScenarioInitializer: steps.InitializeUserRegisterScenario,
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			register.InitializeSuccessfulRegistration(ctx, rc)
			register.InitializeInvalidEmailRegistration(ctx, rc)
			register.InitializeExistingEmailRegistration(ctx, rc)
			// steps.InitializeNoEmailRegistration(ctx, rc)
			// steps.InitializeShortPasswordRegistration(ctx, rc)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{featuresPath},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
