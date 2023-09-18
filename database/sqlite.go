package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var SqlLiteClient *gorm.DB

func CreateDatabase() {
	dbName := os.Getenv("DBNAME")
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	SqlLiteClient = db
}

func CloseDatabase() {

	//Close Sql Database
}
