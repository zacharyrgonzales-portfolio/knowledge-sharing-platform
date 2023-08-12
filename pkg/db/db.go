package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"fmt"
)

func connectToDB() *gorm.DB {
	dsn := "user=zach dbname=knowledge_sharing_platform sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		fmt.Println("GORM Error: Issue establishing connection with database")
	}

	return db
}