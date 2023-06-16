package domain

import (
	"github.com/sgarg2023/golangprojectpractice/mvc/domain"
)

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
