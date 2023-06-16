package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "shubhra", LastName: "Garg", Email: "shubhrag2021@gmail.com"},
		145: {Id: 145, FirstName: "Ankit", LastName: "Garg", Email: "gargankit@gmail.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	user := users[userId]
	if user == nil {
		return nil, errors.New(fmt.Sprintf("user %v not found", userId))
	}
	return user, nil
}
