package main

import (
	"finalproject/configs"
	postgresql "finalproject/database"
	middlewares "finalproject/middleware"
	"finalproject/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load configuration: %v", err)
	}

	db := postgresql.ConnectPostgreSQL()
	r := gin.Default()

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)

	middlewares.Logger(r)

	routes.UserRoutes(r, db)

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})
	r.Run(":3000")

	host := config.SERVER.SERVER_HOST
	port := config.SERVER.SERVER_PORT
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "8000"
	}
	address := host + ":" + port

	log.Printf("server is running on address %s...", address)
	if err := r.Run(address); err != nil {
		logrus.Fatalf("error starting server: %v", err)
	}

}
