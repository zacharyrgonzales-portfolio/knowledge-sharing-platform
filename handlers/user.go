package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/models"
	"github.com/zacharyrgonzales-portfolio/knowledge-sharing-platform/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	// Gin tries to serialize the entire `user`` object, including Article and Comments slice fields and errors. Since we don't need to return the entire user object we can create a response that includes only the fields we want to return

	// Define a struct specifically for sign-up input
	type SignUpInput struct {
		Username string `json:"Username" binding:"required"`
		Email    string `json:"Email" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	// Bind the JSON body of the request to the SignUpInput struct and check for validation errors
	var input SignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input with either Username, Email and/or Password input"})
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
		Username:     input.Username,
		Email:        input.Email,
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

func Login(c *gin.Context) {

	// Define a struct for login input
	type LoginInput struct {
		Username string `json:"Username" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	// Bind the JSON body of the request to the LoginInput struct and check for the validation errors
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input with either Username and/or Password"})
		return
	}

	// Find the user by username
	dbInstance := db.ConnectToDB()
	var user models.User
	if err := dbInstance.Select("id, username, password_hash").Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.JSON(400, gin.H{"error": "Incorrect password"})
		return
	}

	// Gin tries to serialize the entire `user`` object, including Article and Comments slice fields and errors. Since we don't need to return the entire user object we can create a response that includes only the fields we want to return

	// Define a struct for the response
	type LoginResponse struct {
		Username string `json:"Username"`
		Message  string `json:"message"`
	}

	response := LoginResponse{
		Username: user.Username,
		Message:  "Login Successful!",
	}

	// TODO: Generate an authenitcation token (e.g., JWT)

	c.JSON(200, response)
}
