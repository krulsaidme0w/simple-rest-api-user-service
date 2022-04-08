package file

import (
	"errors"
	"os"
	"strconv"
	"sync"

	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/pkg/errors/repository_errors"
)

type Storage struct {
	path  string
	mutex *sync.RWMutex
}

func NewStorage(path string, mutex *sync.RWMutex) *Storage {
	return &Storage{
		path:  path,
		mutex: mutex,
	}
}

func (s *Storage) Save(user *domain.User) (*domain.User, error) {
	if userExists := s.userExists(user); userExists {
		return &domain.User{}, repository_errors.UserAlreadyExists
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	savedUser, err := domain.WriteUserToFile(s.pathToUser(user), user)
	if err != nil {
		return &domain.User{}, err
	}

	return savedUser, nil
}

func (s *Storage) GetByID(id string) (*domain.User, error) {
	user, err := domain.ReadUserFromFile(s.path + "/" + id)
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (s *Storage) Update(user *domain.User) (*domain.User, error) {
	return &domain.User{}, nil
}

func (s *Storage) Delete(user *domain.User) error {
	return nil
}

func (s *Storage) userExists(user *domain.User) bool {
	_, err := os.Stat(s.pathToUser(user))
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func (s *Storage) pathToUser(user *domain.User) string {
	return s.path + "/" + strconv.Itoa(user.ID)
}
