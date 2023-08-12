package handlers

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
	"github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/pkg/db"
    "github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/models"
)

func SignUp(c *gin.Context) {
	
	// Define a struct specifically for sign-up input
	type SignUpInput struct {
		Username string `json:"Username" binding:"required"`
		Email string `json:"Email" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	var input SignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to hash password"})
        return
    }

	// Create the user using the models.User struct
	user := models.User{
		Username: input.Username,
		Email: input.Email,
		PasswordHash: string(hashedPassword),
	}

	// Save the user to the database
	dbInstance := db.ConnectToDB()
	result := dbInstance.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return 
	}

	c.JSON(201, user)
}