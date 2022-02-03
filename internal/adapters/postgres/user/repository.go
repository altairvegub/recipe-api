package user

import (
	"errors"
	"fmt"
	"recipe/internal/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository creates an instance of a postgres repository to interact with postgres.
// Encapsulates all postgres related logic to keep our app independent of leaks.
func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

// Create inserts a new user record into the users table
func (r *repository) Create(user domain.UserDTO) (string, error) {
	userModel := User{
		Email:    user.Email,
		Password: user.Password,
	}

	result := r.db.Create(&userModel)
	if result.Error != nil {
		return "", fmt.Errorf("failed to create user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return "", fmt.Errorf("expected row to be created: %w", result.Error)
	}

	return userModel.ID, nil
}

// GetByEmail retrieves a user from the users table by email
func (r *repository) GetByEmail(email string) (*domain.UserDTO, error) {
	var user User
	result := r.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user: %w", result.Error)
	}

	userDTO := user.ToUserDTO()
	return &userDTO, nil
}
