package domain

// UserDTO is the user domain object
type UserDTO struct {
	ID       string
	Email    string
	Password string
}

// UserRepository is an interface that exposes methods for our domain logic. Completely separated away from external factors like different data sources.
type UserRepository interface {
	Create(u UserDTO) (string, error)
	GetByEmail(email string) (*UserDTO, error)
}
