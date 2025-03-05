package register

import (
	"gobdd/internal/entity"
	mocksjob "gobdd/internal/job/mocks"
	mocksrepo "gobdd/internal/repository/mocks"
	"gobdd/internal/usecase"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

type Context struct {
	user        *entity.User
	emailJob    *mocksjob.MockEmail
	userRepo    *mocksrepo.MockUserRepository
	userUseCase *usecase.UserUseCase
	err         error
}

func NewContext() *Context {
	return &Context{}
}

func (rc *Context) initializeCommon() {
	rc.userRepo = &mocksrepo.MockUserRepository{}
	rc.emailJob = &mocksjob.MockEmail{}
	rc.userUseCase = usecase.NewUserUseCase(rc.userRepo, rc.emailJob)
}

func (rc *Context) theRegistrationShouldFail() error {
	require.Error(nil, rc.err)

	return nil
}

func (rc *Context) anErrorMessageShouldBeShown(message string) error {
	assert.Equal(nil, rc.err.Error(), message)

	return nil
}
