package services

import "domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
