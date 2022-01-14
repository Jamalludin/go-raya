package controllers

import (
	"github.com/gin-gonic/gin"
	"go-grpc/App/db"
	"go-grpc/App/models/mongodb"
	"go-grpc/App/models/postgresql"
	"go-grpc/App/routes"
	"go-grpc/App/utility"
	"log"
)

type UserController struct{}

func (h *UserController) Register(res *gin.Context) {
	var req routes.Register

	/*Request Body validation*/
	if res.BindJSON(&req) != nil {
		res.JSON(406, gin.H{"message": "name, username and password required"})
		res.Abort()
		return
	}

	/*Check Username Already*/
	var user postgresql.User
	if err := db.PGDB.Where(" username = ?", req.Username).First(&user).Error; err == nil {
		res.JSON(403, gin.H{"message": "Username already exist"})
		return
	}

	/*Save To DB Posgresql*/
	user = postgresql.User{
		Username: req.Username,
		Password: utility.EncryptPassword(req.Password),
		Name:     req.Name,
		Email:    req.Email,
	}
	db.PGDB.Create(&user)

	res.JSON(201, gin.H{"message": "Success registration"})
}

func (h *UserController) Login(res *gin.Context) {
	var req routes.Login

	/*Request Body validation*/
	if res.BindJSON(&req) != nil {
		res.JSON(406, gin.H{"message": "username and password required"})
		res.Abort()
		return
	}

	/*Check Username Already*/
	var user postgresql.User
	if err := db.PGDB.Where(" username = ?", req.Username).First(&user).Error; err != nil {
		res.JSON(404, gin.H{"message": "Incorrect Username / Password"})
		return
	}

	if err := utility.CheckPasswordHash(req.Password, user.Password); err != true {
		res.JSON(401, gin.H{"message": "Incorrect Username / Password"})
		return
	}

	jwtToken, expiredTime, err := utility.GenerateToken(req.Username)

	if err != nil {
		res.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	err = mongodb.CreateSession(req.Username, expiredTime)

	if err != nil {
		log.Printf("Error Destroy Create %s", err)
		res.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	res.JSON(201, gin.H{"message": "SUCCESS_LOGIN", "token": jwtToken})
}
