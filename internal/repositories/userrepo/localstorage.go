package userrepo

import (
	"errors"
	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/pkg"
	"strconv"
	"strings"
	"sync"
)

type Storage struct {
	Mutex    *sync.RWMutex
	Filename string
}

var CannotFindUser = errors.New("CannotFindUser")

func (s Storage) Save(user domain.User) (domain.User, error) {
	userStr := strconv.Itoa(user.ID) + " " + user.Username + " " + user.Name + " " + user.Photo + " " + strconv.Itoa(user.Age)

	err := pkg.AddStringToFile(s.Filename, userStr, s.Mutex)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s Storage) Delete(user domain.User) error {
	return nil
}

func (s Storage) Find(searchType string, search string) (domain.User, error) {
	userStrings, err := pkg.GetAllStringsFromFile(s.Filename, s.Mutex)
	if err != nil {
		return domain.User{}, err
	}

	for _, usr := range userStrings {
		userFields := strings.Split(usr, " ")
		user := domain.User{}
		for i := range userFields {
			switch i {
			case 0:
				user.ID, _ = strconv.Atoi(userFields[0])
			case 1:
				user.Username = userFields[1]
			case 2:
				user.Name = userFields[2]
			case 3:
				user.Photo = userFields[3]
			case 4:
				user.Age, _ = strconv.Atoi(userFields[4])
			}
		}

		switch searchType {
		case "id":
			if strconv.Itoa(user.ID) == search {
				return user, nil
			}
		case "username":
			if user.Username == search {
				return user, nil
			}
		case "name":
			if user.Name == search {
				return user, nil
			}
		}
	}

	return domain.User{}, CannotFindUser
}
