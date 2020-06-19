package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // necessary to connect on postgres
)

// DB variable to execute queries
var DB *gorm.DB

// ConnectDataBase method to connect database
func ConnectDataBase() {
	database, err := gorm.Open("postgres", `
		host=localhost 
		port=5432 
		user=homestead 
		dbname=homestead 
		password=secret
		sslmode=disable
	`)

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}

	database.LogMode(true)
	database.AutoMigrate(&Book{})

	DB = database
}
