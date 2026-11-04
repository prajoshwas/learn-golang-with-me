package routes

import (
	"example/rest-api/models"
	"example/rest-api/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User Routes
func registerUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid request body",
			"code":    "REAPI400",
		})
		return
	}

	err = user.SaveUser()
	if err != nil {
		context.JSON(500, gin.H{
			"message": "Could not register user. Try again later",
			"code":    "REAPI500",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func loginUser(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid request body",
			"code":    "REAPI400",
		})
		return
	}

	err = user.ValidateCredentials()
	fmt.Printf("ValidateCredentials error = %v; nil = %t\n", err, err == nil)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user.",
			"code":    "REAPI401",
		})
		return
	}

	token, err := utils.GenerateJWTToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate JWT token.",
			"code":    "REAPI500",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Log in successfully",
		"token":   token,
	})
}
