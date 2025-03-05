package register

import (
	"errors"
	"gobdd/internal/entity"

	"github.com/stretchr/testify/mock"

	"github.com/cucumber/godog"
)

func (rc *Context) anExistingUserWithEmail(_ string) error {
	rc.initializeCommon()
	rc.userRepo.EXPECT().Create(entity.User{
		Name:  "John Doe",
		Email: "existing@example.com",
	}).Return(errors.New("email already exists"))
	rc.emailJob.EXPECT().Send(mock.Anything).Return(nil)

	return nil
}

func (rc *Context) aNewUserTriesToRegisterWithEmail(email string) error {
	rc.err = rc.userUseCase.RegisterUser("John Doe", email)
	return nil
}

func InitializeExistingEmailRegistration(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^an existing user with email "([^"]*)"$`, rc.anExistingUserWithEmail)
	ctx.Step(`^a new user tries to register with email "([^"]*)"$`, rc.aNewUserTriesToRegisterWithEmail)
	ctx.Step(`^the registration should fail$`, rc.theRegistrationShouldFail)
	ctx.Step(`^an error message "([^"]*)" should be shown$`, rc.anErrorMessageShouldBeShown)
}
