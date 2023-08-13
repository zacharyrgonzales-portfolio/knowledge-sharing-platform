package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/handlers"
	"github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/models"
	"github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/pkg/db"
)

func main() {
	// Connect to the PostgresSQL database using GORM
	db := db.ConnectToDB()

	// Migrate the schema for User and Article models
	db.AutoMigrate(&models.User{},&models.Article{})

	// Set Gin to release mode in production
	// gin.SetMode(gin.ReleaseMode)

	// Create a default Gin engine
	r := gin.Default()

    // Set Up Routes
	// Test Get
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Knowledge Sharing Platform!",
        })
    })

	// Signup POST
	r.POST("/signup", handlers.SignUp)

	// Login POST
	r.POST("/login", handlers.Login)

	// Create Article POST
	r.POST("/articles", handlers.CreateArticle)

	// Retreive Articles GET
	r.GET("/articles", handlers.GetArticles)

	// Retreive Article by ID GET
	r.GET("/articles/:id", handlers.GetArticle)

	// Start the server on port 8080
	r.Run(":8080")
}