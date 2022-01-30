package service

import (
	"fmt"
	"recipe/internal/database/user"
)

type Service interface {
	Signup(email string, password string) error
}

type service struct {
	userRepo user.Repository
}

func New(userRepo user.Repository) *service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Signup(email string, password string) error {
	existingUser, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to get user by email: %w", err)
	}

	if existingUser != nil {
		return ErrResourceAlreadyExists
	}

	err = s.userRepo.Create(user.UserModel{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
