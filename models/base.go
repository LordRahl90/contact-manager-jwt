package models

import (
	"fmt"
	"log"
	"os"

	"github.com/LordRahl90/contacts_manager/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	// dbPort := os.Getenv("db_port")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	// fmt.Print(dbURI)
	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		utils.HandleError(err, "Cannot Open the database connection.")
	}
	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

//GetDb - function to retrieve the db connection
func GetDb() *gorm.DB {
	return db
}
