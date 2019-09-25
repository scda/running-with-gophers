package models

import (
	"../database"
)

func Initialize() {
	database.DB.AutoMigrate(
		&User{},
	)
}
