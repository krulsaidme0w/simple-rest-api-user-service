package ports

import (
	"golang_pet_project_1/internal/core/domain"
)

type UserService interface {
	Create(user domain.User) (domain.User, error)
	Find(searchType string, search string) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(user domain.User) error
}
