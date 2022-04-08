package main

import (
	"errors"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []user{
	user{Username: "alex", Password: "Control2022"},
}

func isUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func registerNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password must not be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := user{Username: username, Password: password}

	userList = append(userList, u)

	return &u, nil
}

func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
