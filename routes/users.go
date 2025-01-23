package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

// POST -> /signup
func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}
	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user has been created"})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticated user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user has been logged in", "token": token})

}
