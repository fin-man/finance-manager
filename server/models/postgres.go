package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

//DBInit initializes the database
func DBInit() (*gorm.DB, error) {

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DB"),
	)
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open("postgres", dbinfo) // gorm checks Ping on Open
		if err == nil {
			break
		}
		log.Printf("Error while connecting to DB : %s", err.Error())
		log.Println("Trying to connect ... waiting 5 seconds")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return DB, err
	}

	//create tables at start
	err = DB.AutoMigrate(&TransactionModel{}).Error
	if err != nil {
		log.Fatal(err)
	}
	//add relationship
	// DB.Model(&Certificate{}).AddForeignKey("cust_id", "customers(cust_id)", "CASCADE", "CASCADE")

	return DB, err
}
