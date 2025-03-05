package register

import (
	"gobdd/internal/entity"

	"github.com/stretchr/testify/mock"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/require"
)

func (rc *Context) aNewUserWithValidDetails() error {
	rc.initializeCommon()
	rc.userRepo.EXPECT().Create(entity.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}).Return(nil)
	rc.emailJob.EXPECT().Send(mock.Anything).Return(nil)

	return nil
}

func (rc *Context) theUserRegisters() error {
	rc.err = rc.userUseCase.RegisterUser("John Doe", "john@example.com")

	return nil
}

func (rc *Context) theUserAccountShouldBeCreated() error {
	require.NoError(nil, rc.err)
	// Expect(rc.user).NotTo(BeNil())
	// Expect(rc.user.Name).To(Equal("John Doe"))
	// Expect(rc.user.Email).To(Equal("john@example.com"))

	return nil
}

func (rc *Context) aWelcomeEmailShouldBeSent() error {
	require.NoError(nil, rc.err)

	return nil
}

func InitializeSuccessfulRegistration(ctx *godog.ScenarioContext, rc *Context) {
	ctx.Step(`^a new user with valid details$`, rc.aNewUserWithValidDetails)
	ctx.Step(`^the user registers$`, rc.theUserRegisters)
	ctx.Step(`^the user account should be created$`, rc.theUserAccountShouldBeCreated)
	ctx.Step(`^a welcome email should be sent$`, rc.aWelcomeEmailShouldBeSent)
}
