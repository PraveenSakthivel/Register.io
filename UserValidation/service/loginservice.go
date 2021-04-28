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
func LoginUser(netid string, password string) bool {
	// Check if user exists in DB
	users := []models.User{}
	models.DB.Where("netid = ? AND password = ?", netid, password).Find(&users)
	return len(users) == 1
}
