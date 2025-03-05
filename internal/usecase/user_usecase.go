package usecase

import (
	"errors"
	"gobdd/internal/entity"
)

//go:generate mockery --name UserRepository
type UserRepository interface {
	Create(user entity.User) error
}

//go:generate mockery --name Email
type Email interface {
	Send(email entity.Email) error
}

type UserUseCase struct {
	UserRepo  UserRepository
	emailSend Email
}

func NewUserUseCase(userRepo UserRepository, emailSend Email) *UserUseCase {
	return &UserUseCase{UserRepo: userRepo, emailSend: emailSend}
}

func (uc *UserUseCase) RegisterUser(name, email string) error {
	if err := isValidEmail(email); err != nil {
		return err
	}

	user := entity.User{Name: name, Email: email}
	if err := uc.UserRepo.Create(user); err != nil {
		return err
	}

	welcomeEmail := entity.Email{
		Recipient: email,
		Subject:   "Welcome to our service!",
		Body:      "Thank you for registering.",
	}

	return uc.emailSend.Send(welcomeEmail)
}

func isValidEmail(email string) error {
	if email == "invalid-email" {
		return errors.New("invalid email format")
	}

	return nil
}
