package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-grpc/App/controllers"
	"go-grpc/App/db"
	_ "google.golang.org/grpc"
	"log"
	_ "log"
	_ "net"
	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("file .env Not Found")
	}
}

func setEnv(key string) string {
	return os.Getenv(key)
}

func main() {
	/*lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}*/

	/*r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()*/

	router := gin.Default()

	db.ConnectDBSql(os.Getenv("POSTGRES_DB_USER"), os.Getenv("POSTGRES_DB_PASSWORD"), os.Getenv("POSTGRES_DB_PORT"), os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_NAME"))

	api := router.Group("/api/usr")
	{
		userController := new(controllers.UserController)
		api.POST("/auth/register", userController.Register)
		api.POST("/auth/login", userController.Login)
	}

	router.NoRoute(func(res *gin.Context) {
		res.JSON(404, gin.H{"Message": "Not found!"})
	})

	apiPort := ":" + setEnv("API_PORT")

	router.Run(apiPort)
}
