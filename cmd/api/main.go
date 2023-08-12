package main

import (
	"github.com/gin-gonic/gin"
    "github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/pkg/db"
    "github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/models"
)

func main() {
	// Set Gin to release mode in production
	// gin.SetMode(gin.ReleaseMode)

	// Create a default Gin engine
	r := gin.Default()

    // Set up routes
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Knowledge Sharing Platform!",
        })
    })

	// Start the server on port 8080
	r.Run(":8080")

	// Connect to the PostgresSQL database using GORM
	db := db.ConnectToDB()

	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{})
}