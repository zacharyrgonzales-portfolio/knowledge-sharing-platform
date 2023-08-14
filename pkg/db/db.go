package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectToDB() *gorm.DB {
	dsn := "user=zach dbname=knowledge_sharing_platform sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		fmt.Println("GORM Error: Issue establishing connection with database")
	}

	return db
}
