package controllers

import (
	"github.com/gin-gonic/gin"
	"go-grpc/App/db"
	"go-grpc/App/models"
	"go-grpc/App/routes"
)

type UserController struct{}

func (h *UserController) Register(res *gin.Context) {
	var req routes.Register

	// Body json validation
	if res.BindJSON(&req) != nil {
		res.JSON(406, gin.H{"message": "name, username and password required"})
		res.Abort()
		return
	}

	// Search existing user
	var user models.User
	if err := db.PGDB.Where(" username = ?", req.Username).First(&user).Error; err == nil {
		res.JSON(403, gin.H{"message": "Username already exist"})
		return
	}

	// Save to SQL database
	user = models.User{
		Username: req.Username,
		Name:     req.Name,
	}
	db.PGDB.Create(&user)

	// Positive Response
	res.JSON(201, gin.H{"message": "Success registration"})
}
