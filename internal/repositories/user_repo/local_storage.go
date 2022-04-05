package user_repo

import (
	"errors"
	"golang_pet_project_1/internal/core/domain"
	"golang_pet_project_1/pkg/utils"
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
	s.Mutex.Lock()

	indexes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	s.Mutex.Unlock()

	indexesArray := strings.Split(string(indexes), "\n")
	indexesArray = utils.DeleteEmpty(indexesArray)
	indexesArray = utils.InsertToArray(indexesArray, data)

	err = utils.WriteStringsToFile(path, indexesArray, s.Mutex)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) Save(user domain.User) (domain.User, error) {
	userStr := strconv.Itoa(user.ID) + " " + user.Username + " " + user.Name + " " + user.Photo + " " + strconv.Itoa(user.Age)

	err := utils.CreateFileWithData(s.Path, strconv.Itoa(user.ID), userStr, s.Mutex)
	if err != nil {
		switch err {
		case utils.FileAlreadyExists:
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
	userString, err := utils.GetFirstStringFromFile(s.Path+"/"+id, s.Mutex)
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
	usernameIndexes, err := utils.GetAllStringsFromFile(s.Path+"/"+usernameIndex, s.Mutex)
	if err != nil {
		return domain.User{}, err
	}

	id, err := utils.FindInSortedArray(usernameIndexes, username)
	if err != nil {
		return domain.User{}, err
	}

	user, err := s.GetByID(id)
	if err != nil {
		return domain.User{}, CannotFindUser
	}

	return user, nil
}

func (s Storage) GetByName(name string) (domain.User, error) {
	nameIndexes, err := utils.GetAllStringsFromFile(s.Path+"/"+nameIndex, s.Mutex)
	if err != nil {
		return domain.User{}, err
	}

	id, err := utils.FindInSortedArray(nameIndexes, name)
	if err != nil {
		return domain.User{}, err
	}

	user, err := s.GetByID(id)
	if err != nil {
		return domain.User{}, CannotFindUser
	}

	return user, nil
}

func (s Storage) Update(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
