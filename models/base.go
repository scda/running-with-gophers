package models

import (
	"../database"
)

// Initialize will add model schema to the database
func Initialize() {
	database.DB.AutoMigrate(
		&User{},
	)
}
