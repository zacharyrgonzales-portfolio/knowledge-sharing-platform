package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/models"
	"github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/pkg/db"
)

// CreateArticleInput defines the input for creating a new article
type CreateArticleInput struct {
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	Category string   `json:"category" binding:"required"`
	Tags     []string `json:"tags" binding:"required"`
	AuthorID uint     `json:"author_id" binding:"required"`
}

// CreateArticle handles the creation of a new article
func CreateArticle(c *gin.Context) {
	// Bind the JSON input to the CreateArticleInput struct
	var input CreateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid article input"})
		return
	}

	// Create the article object
	article := models.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Tags:     strings.Join(input.Tags, ","), // Join the tags into a comma-seperated string
		AuthorID: input.AuthorID,
	}

	// Save the article to the database
	dbInstance := db.ConnectToDB()
	result := dbInstance.Create(&article)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create article"})
		return
	}

	// Respond with newly created article
	c.JSON(201, article)
}

// GetArticles handles the retreival of articles by the author
func GetArticles(c *gin.Context) {
	// Connect to the database
	dbInstance := db.ConnectToDB()

	// Declare a slice to hold the articles
	var articles []models.Article

	// Query the database for all articles
	if err := dbInstance.Find(&articles).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve articles"})
		return
	}

	// Respond with the articles as JSON
	c.JSON(200, articles)
}

// GetArticle handles the retrieval of an article by ID
func GetArticle(c *gin.Context) {
	// Connect to the database
	dbInstance := db.ConnectToDB()

	// Get the article ID from the URL parameter
	articleID := c.Param("id")

	// Declare a variable to hold the article
	var article models.Article

	// Query the database for the article by ID
	if err := dbInstance.First(&article, articleID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Article not found"})
		return
	}

	// Respond with the article as JSON
	c.JSON(200, article)
}
