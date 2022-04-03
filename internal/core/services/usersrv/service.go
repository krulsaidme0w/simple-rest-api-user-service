package usersrv

import (
	"errors"
	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/internal/core/ports"
)

type Service struct {
	Storage ports.UserRepository
}

var InvalidSearchType = errors.New("InvalidSearchType")

func (s Service) Create(user domain.User) (domain.User, error) {
	savedUser, err := s.Storage.Save(user)
	if err != nil {
		return domain.User{}, err
	}

	return savedUser, nil
}

func (s Service) Find(searchType string, search string) (domain.User, error) {
	if searchType == "id" || searchType == "username" || searchType == "name" {
		user, err := s.Storage.Find(searchType, search)
		return user, err
	}

	return domain.User{}, InvalidSearchType
}

func (s Service) Update(user domain.User) (domain.User, error) {
	return user, nil
}

func (s Service) Delete(user domain.User) error {
	return nil
}
