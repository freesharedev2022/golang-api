package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initialMigration(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
}

func ConnectDatabase() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_hostname := os.Getenv("MYSQL_HOST")
	db_name := os.Getenv("MYSQL_DB")
	db_user := os.Getenv("MYSQL_USER")
	db_pass := os.Getenv("MYSQL_PASSWORD")
	db_port := os.Getenv("MYSQL_PORT")

	var credentials = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", db_user, db_pass, db_hostname, db_port, db_name)

	database, err := gorm.Open(mysql.Open(credentials), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}

	if err != nil {
		fmt.Printf("error connect to database %s", err)
	} else {
		initialMigration(database)
	}

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = database
}
