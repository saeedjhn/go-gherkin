package job

import "gobdd/internal/entity"

type Email struct{}

func NewEmail() *Email {
	return &Email{}
}

func (e *Email) Send(email entity.Email) error {
	return nil
}
