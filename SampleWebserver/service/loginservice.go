package service

import (
	"main/models"
)

// LoginService ...
type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

// LoginUser ...
func LoginUser(email string, password string) bool {
	// Check if user exists in DB
	users := []models.User{}
	models.DB.Where("email = ? AND password = ?", email, password).Find(&users)
	if len(users) == 1 {
		return true
	}
	return false
}
