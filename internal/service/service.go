package service

import (
	"fmt"
	"recipe/internal/auth"
	"recipe/internal/domain"
)

// Service encapsulates all use cases of our application (business / domain logic)
type Service interface {
	Signup(email string, password string) (string, error)
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

// Signup checks if email is in database, if not creates user, returns JWT
func (s *service) Signup(email string, password string) (string, error) {
	existingUser, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", fmt.Errorf("failed to get user by email: %w", err)
	}

	if existingUser != nil {
		return "", ErrResourceAlreadyExists
	}

	hashedPwd, err := auth.HashAndSaltPassword(password)
	if err != nil{
		return "", fmt.Errorf("generate password hash failed: %w", err)
	}

	userDTO := domain.UserDTO{
		Email:    email,
		Password: hashedPwd,
	}

	userID, err := s.userRepo.Create(userDTO)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	token, err := auth.CreateJwtToken(userID)
	if err != nil{
		fmt.Errorf("couldn't create token: %w", err)
	}

	return token, nil
}
