package userservice

import (
	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/internal/core/ports"
	"golang_pet_project_1/pkg/errors/handler_errors"
)

type Service struct {
	storage ports.UserRepository
}

func NewService(storage ports.UserRepository) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) Create(user *domain.User) (*domain.User, error) {
	savedUser, err := s.storage.Save(user)
	if err != nil {
		return &domain.User{}, err
	}

	return savedUser, nil
}

func (s Service) Find(searchType string, search string) (*domain.User, error) {
	switch searchType {
	case "id":
		user, err := s.storage.GetByID(search)
		return user, err
	case "username":
		user, err := s.storage.GetByUsername(search)
		return user, err
	case "name":
		user, err := s.storage.GetByName(search)
		return user, err
	default:
		return &domain.User{}, handler_errors.InvalidSearchType
	}
}

func (s Service) Update(user *domain.User) (*domain.User, error) {
	updatedUser, err := s.storage.Update(user)
	if err != nil {
		return &domain.User{}, err
	}

	return updatedUser, nil
}

func (s Service) Delete(user *domain.User) error {
	return s.storage.Delete(user)
}
