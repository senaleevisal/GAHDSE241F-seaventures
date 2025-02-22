package controllers

import (
	"net/http"
	"os"
	"seaventures/src/models"
	"seaventures/src/service"
	"seaventures/src/helpers"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	

	secretKey := os.Getenv("JWT_SECRET_KEY")
    token, err := helpers.GenerateJWT(user.ID, secretKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "token": token, "user": user})
}


func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.Login(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
    token, err := helpers.GenerateJWT(user.ID, secretKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user": user})
}



func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You have accessed a protected endpoint"})
}
