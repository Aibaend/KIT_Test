package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB


func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	//
	//username := os.Getenv("db_user")
	//password := os.Getenv("db_pass")
	//dbName := os.Getenv("db_name")
	//dbHost := os.Getenv("db_host")
	////test
	username := os.Getenv("db_user_test")
	password := os.Getenv("db_pass_test")
	dbName := os.Getenv("db_name_test")
	dbHost := os.Getenv("db_host_test")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)


	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}else {
		fmt.Println("Connected successfully")
	}

	db = conn
	db.LogMode(true)

	db.AutoMigrate(&Dictionary{})
}

func GetDB() *gorm.DB {
	return db
}
