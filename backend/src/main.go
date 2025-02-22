package main

import (
	"log"
	"seaventures/src/config"
	"seaventures/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	r := gin.Default()

	config.ConnectDB()

	routes.RegisterRoutes(r)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("env not found")
		err = godotenv.Load()
		if err != nil {
			log.Fatal("error loading env")
		}
	}
}
