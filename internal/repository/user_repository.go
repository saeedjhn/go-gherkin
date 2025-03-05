package repository

import (
	"errors"
	"gobdd/internal/entity"
)

type InMemoryUserRepository struct {
	users map[string]entity.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[string]entity.User)}
}

func (repo *InMemoryUserRepository) Create(user entity.User) error {
	if _, exists := repo.users[user.Email]; exists {
		return errors.New("email already exists")
	}
	repo.users[user.Email] = user
	return nil
}
