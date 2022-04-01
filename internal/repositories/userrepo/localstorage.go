package userrepo

import (
	"os"
	"os/user"
)

type storage struct {
	file os.File
}

func (s storage) Save(user user.User) (user.User, error) {
	return user, nil
}

func (s storage) Delete(user user.User) error {
	return nil
}

func (s storage) Find(searchType string, search string) (user.User, error) {
	return user.User{}, nil
}
