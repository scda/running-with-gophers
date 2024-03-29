package models

import (
	"log"

	"../database"

	"github.com/jinzhu/gorm"
)

// User type model
type User struct {
	gorm.Model

	Name  string
	Email string
}

// GetUsers returns all users found
func GetUsers() (users []*User) {
	if err := database.DB.Find(&users).Error; err != nil {
		log.Println("GetUsersErr:")
		return nil
	}
	return
}

// GetUserByID returns single user object searched by ID
func GetUserByID(id uint) *User {
	user := new(User)
	user.ID = id

	if err := database.DB.First(user).Error; err != nil {
		log.Println("GetUserByIdErr: ", err)
		return nil
	}

	return user
}

// CreateUser creates a new user entry
func CreateUser(user *User) *User {
	if err := database.DB.Create(user).Error; err != nil {
		log.Println("CreateUserErr: ", err)
		return nil
	}

	return user
}
