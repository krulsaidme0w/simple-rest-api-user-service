package domain

import (
	"errors"
	"strconv"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	Age      int    `json:"age"`
}

var CannotBuildUser = errors.New("CannotBuildUser")

const userFieldsCount = 5

func BuildUser(userString string) (User, error) {
	userFields := strings.Split(userString, " ")
	if len(userFields) != userFieldsCount {
		return User{}, CannotBuildUser
	}

	id, err := strconv.Atoi(userFields[0])
	if err != nil {
		return User{}, CannotBuildUser
	}

	age, err := strconv.Atoi(userFields[4])
	if err != nil {
		return User{}, CannotBuildUser
	}

	user := User{
		ID:       id,
		Username: userFields[1],
		Name:     userFields[2],
		Photo:    userFields[3],
		Age:      age,
	}

	return user, nil
}
