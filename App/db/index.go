package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-grpc/App/models"
)

var PGDB *gorm.DB

func ConnectDBSql(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	database, err := gorm.Open("postgres", DBURL)

	if err != nil {
		panic("Failed to connect to postgres database at " + DBURL)
	}

	database.AutoMigrate(&models.User{})

	PGDB = database
}
