package user

import (
	"recipe/internal/domain"
	"time"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

// User is the schema for our users table in postgres
type User struct {
	ID        string `gorm:"primaryKey"`
	Name      *string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate hook that generates a UUID before every insertion
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	return nil
}

// ToUserDTO translates the postgres user model into our application's user domain model
func (u *User) ToUserDTO() domain.UserDTO {
	return domain.UserDTO{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}
