package userrepository

import (
	"strconv"

	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/internal/repositories/cache"
	"golang_pet_project_1/internal/repositories/file"
)

type Storage struct {
	storage *file.Storage
	cache   *cache.Cache
}

func NewStorage(storage *file.Storage, cache *cache.Cache) *Storage {
	return &Storage{
		storage: storage,
		cache:   cache,
	}
}

func (s Storage) Save(user *domain.User) (*domain.User, error) {
	savedUser, err := s.storage.Save(user)
	if err != nil {
		return &domain.User{}, err
	}

	err = s.cache.Add(user)
	if err != nil {
		return &domain.User{}, err
	}

	return savedUser, nil
}

func (s Storage) GetByID(id string) (*domain.User, error) {
	user, err := s.storage.GetByID(id)
	if err != nil {
		return &domain.User{}, err
	}

	return user, err
}

func (s Storage) GetByUsername(username string) (*domain.User, error) {
	id, err := s.cache.GetIDByUsername(username)
	if err != nil {
		return &domain.User{}, err
	}

	return s.GetByID(strconv.Itoa(id))
}

func (s Storage) GetByName(name string) (*domain.User, error) {
	id, err := s.cache.GetIDByName(name)
	if err != nil {
		return &domain.User{}, err
	}

	return s.GetByID(strconv.Itoa(id))
}

func (s Storage) Update(user *domain.User) (*domain.User, error) {
	return &domain.User{}, nil
}

func (s Storage) Delete(user *domain.User) error {
	return nil
}