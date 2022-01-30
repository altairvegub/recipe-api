package service

import (
	"fmt"
	"recipe/internal/auth"
	"recipe/internal/domain"
)

// Service encapsulates all use cases of our application (business / domain logic)
type Service interface {
	Signup(email string, password string) error
}

// service is a private struct which stores a generic UserRepository. The caller of this will be responsible
// supplying a UserRepository for whichever data source is required.
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


	hashedPwd, err := auth.HashAndSaltPassword(password)
	if err != nil{
		return fmt.Errorf("Generate password hash failed: %w", err)
	}

	userDTO := domain.UserDTO{
		Email:    email,
		Password: hashedPwd,
	}

	err = s.userRepo.Create(userDTO)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
