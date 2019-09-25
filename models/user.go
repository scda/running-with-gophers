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
