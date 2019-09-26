package database

import (
	"fmt"
	"log"

	"../config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sql dialect
)

var initialized bool = false
var databaseConnection *gorm.DB

// DB is a database connection that will be initialized at program start
var (
	DB = dbAccessor()
)

func dbAccessor() *gorm.DB {
	if !initialized {
		Connect()
	}

	return databaseConnection
}

// Connect to database
func Connect() {
	log.Println("connecting database ...")

	DB, err := gorm.Open(config.DbType, config.DbFile)
	if err != nil {
		panic(fmt.Sprintf("ERROR occured trying to open database. err=%+v", err))
	}

	databaseConnection = DB
	initialized = true
}

// Disconnect from database
func Disconnect() {
	log.Println("disconnecting database ...")

	_ = databaseConnection.Close()
	initialized = false
}
