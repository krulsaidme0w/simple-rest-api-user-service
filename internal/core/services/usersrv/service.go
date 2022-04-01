package usersrv

import (
	"golang_pet_project_1/internal/core/ports"
	"os/user"
)

type service struct {
	storage ports.UserRepository
}

func (s service) Create(user user.User) (user.User, error) {
	return user, nil
}

func (s service) Find(searchType string, search string) (user.User, error) {
	return user.User{}, nil
}

func (s service) Update(user user.User) (user.User, error) {
	return user, nil
}

func (s service) Delete(user user.User) error {
	return nil
}
