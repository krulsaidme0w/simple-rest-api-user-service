package userrepo

import (
	"errors"
	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/pkg"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Storage struct {
	Path  string
	Mutex *sync.RWMutex
}

const (
	usernameIndex = "username_index"
	nameIndex     = "name_index"
)

var CannotFindUser = errors.New("CannotFindUser")
var CannotAddIndex = errors.New("CannotAddIndex")
var UserAlreadyExists = errors.New("UserAlreadyExists")

func (s Storage) addIndex(path string, data string) error {
	indexes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	indexesArray := strings.Split(string(indexes), "\n")
	indexesArray = pkg.DeletyEmpty(indexesArray)
	indexesArray = pkg.InsertToArray(indexesArray, data)

	err = pkg.WriteStringsToFile(path, indexesArray, s.Mutex)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) Save(user domain.User) (domain.User, error) {
	userStr := strconv.Itoa(user.ID) + " " + user.Username + " " + user.Name + " " + user.Photo + " " + strconv.Itoa(user.Age)

	err := pkg.CreateFileWithData(s.Path, strconv.Itoa(user.ID), userStr, s.Mutex)
	if err != nil {
		switch err {
		case pkg.FileAlreadyExists:
			return domain.User{}, UserAlreadyExists
		default:
			return domain.User{}, CannotAddIndex
		}
	}

	err = s.addIndex(s.Path+"/"+usernameIndex, strconv.Itoa(user.ID)+" "+user.Username)
	if err != nil {
		return domain.User{}, CannotAddIndex
	}

	err = s.addIndex(s.Path+"/"+nameIndex, strconv.Itoa(user.ID)+" "+user.Name)
	if err != nil {
		return domain.User{}, CannotAddIndex
	}

	return user, nil
}

func (s Storage) Delete(user domain.User) error {
	return nil
}

func (s Storage) GetByID(id string) (domain.User, error) {
	userString, err := pkg.GetFirstStringFromFile(s.Path+"/"+id, s.Mutex)
	if err != nil {
		return domain.User{}, CannotFindUser
	}

	user, err := domain.BuildUser(userString)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s Storage) GetByUsername(username string) (domain.User, error) {
	return domain.User{}, nil
}

func (s Storage) GetByName(name string) (domain.User, error) {
	return domain.User{}, nil
}
