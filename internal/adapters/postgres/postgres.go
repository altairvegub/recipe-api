package postgres

import (
	"fmt"

	"recipe/internal/adapters/postgres/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New creates a new instance of gorm db
func New(port int, host, username, password, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host,
		username,
		password,
		database,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	db.AutoMigrate(&user.User{})

	return db, nil
}
