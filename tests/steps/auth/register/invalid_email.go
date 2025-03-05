package register

import (
	"errors"
	"gobdd/internal/entity"

	"github.com/stretchr/testify/mock"

	"github.com/cucumber/godog"
)

func (rc *Context) aNewUserWithInvalidEmail(_ string) error {
	rc.initializeCommon()
	rc.userRepo.EXPECT().Create(entity.User{
		Name:  "John Doe",
		Email: "invalid-email",
	}).Return(errors.New("invalid email format"))
	rc.emailJob.EXPECT().Send(mock.Anything).Return(nil)

	return nil
}

func (rc *Context) theUserTriesToRegister() error {
	rc.err = rc.userUseCase.RegisterUser("John Doe", "invalid-email")

	return nil
}

func InitializeInvalidEmailRegistration(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^a new user with invalid email "([^"]*)"$`, rc.aNewUserWithInvalidEmail)
	ctx.Step(`^the user tries to register$`, rc.theUserTriesToRegister)
	ctx.Step(`^the registration should fail$`, rc.theRegistrationShouldFail)
	ctx.Step(`^an error message "([^"]*)" should be shown$`, rc.anErrorMessageShouldBeShown)
}
