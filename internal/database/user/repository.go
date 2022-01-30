package user

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(u UserModel) error
	GetByEmail(email string) (*UserModel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) Create(user UserModel) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("expected row to be created: %w", result.Error)
	}

	return nil
}

func (r *repository) GetByEmail(email string) (*UserModel, error) {
	var user UserModel
	result := r.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user: %w", result.Error)
	}

	return &user, nil
}
