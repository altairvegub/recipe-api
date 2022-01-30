package service

import (
	"fmt"
	"recipe/internal/domain"
)

// Service encapsulates all use cases of our application (business / domain logic)
type Service interface {
	Signup(email string, password string) error
}

// service is a private struct which stores a generic UserRepository. The caller of this will be responsible
// suppyling a UserRepository for whichever data source is required.
type service struct {
	userRepo domain.UserRepository
}

func New(userRepo domain.UserRepository) *service {
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

	userDTO := domain.UserDTO{
		Email:    email,
		Password: password,
	}

	err = s.userRepo.Create(userDTO)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
