package user

import (
	"time"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        string `gorm:"primaryKey"`
	Name      *string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	return nil
}
