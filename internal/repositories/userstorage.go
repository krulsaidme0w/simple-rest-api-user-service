package userrepository

import (
	"strconv"

	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/internal/core/ports"
	"golang_pet_project_1/internal/repositories/cache"
)

type Storage struct {
	storage ports.UserStorage
	cache   *cache.Cache
}

func NewStorage(storage ports.UserStorage, cache *cache.Cache) *Storage {
	return &Storage{
		storage: storage,
		cache:   cache,
	}
}

func (s *Storage) Save(user *domain.User) (*domain.User, error) {
	err := s.cache.Add(user)
	if err != nil {
		return &domain.User{}, err
	}

	savedUser, err := s.storage.Save(user)
	if err != nil {
		return &domain.User{}, err
	}

	return savedUser, nil
}

func (s *Storage) GetByID(id string) (*domain.User, error) {
	user, err := s.storage.GetByID(id)
	if err != nil {
		return &domain.User{}, err
	}

	return user, err
}

func (s *Storage) GetByUsername(username string) (*domain.User, error) {
	id, err := s.cache.GetIDByUsername(username)
	if err != nil {
		return &domain.User{}, err
	}

	return s.GetByID(strconv.Itoa(id))
}

func (s *Storage) GetByName(name string) (*domain.User, error) {
	id, err := s.cache.GetIDByName(name)
	if err != nil {
		return &domain.User{}, err
	}

	return s.GetByID(strconv.Itoa(id))
}

func (s *Storage) Update(user *domain.User) (*domain.User, error) {
	updatedUser, err := s.storage.Update(user)
	if err != nil {
		return &domain.User{}, err
	}

	if err = s.cache.Update(user); err != nil {
		return &domain.User{}, err
	}

	return updatedUser, nil
}

func (s *Storage) Delete(userID string) error {
	user, err := s.storage.GetByID(userID)
	if err != nil {
		return err
	}

	err = s.storage.Delete(userID)
	if err != nil {
		return err
	}

	err = s.cache.Delete(user.Username, user.Name)
	return err
}
