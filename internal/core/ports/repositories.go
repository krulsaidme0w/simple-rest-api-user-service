package ports

import (
	"golang_pet_project_1/internal/core/domain"
)

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	GetByID(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	GetByName(name string) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(userID string) error
}
