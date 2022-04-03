package ports

import (
	"golang_pet_project_1/internal/core/domain"
)

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	Delete(user domain.User) error
	Find(searchType string, search string) (domain.User, error)
}
