package domain

import (
	"encoding/json"
	"io/ioutil"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	Age      int    `json:"age"`
}

func ReadUserFromFile(path string) (*User, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return &User{}, err
	}

	var user User
	if err = json.Unmarshal(b, &user); err != nil {
		return &User{}, err
	}

	return &user, nil
}

func WriteUserToFile(path string, user *User) (*User, error) {
	b, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		return &User{}, err
	}

	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		return &User{}, err
	}

	return user, nil
}
