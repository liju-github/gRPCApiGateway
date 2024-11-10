package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/liju-github/EcommerceApiGatewayService/clients"
	config "github.com/liju-github/EcommerceApiGatewayService/configs"
	"github.com/liju-github/EcommerceApiGatewayService/router"
	"github.com/liju-github/EcommerceApiGatewayService/utils"
)

func main() {
	// Load environment variables
	config := config.LoadConfig()

	// Set JWT secret key in utils package
	utils.SetJWTSecretKey(config.JWTSecretKey)

	// Initialize gRPC clients
	Client, err := clients.InitClients(config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer Client.Close()

	// Create a new Gin router
	ginRouter := gin.Default()

	// Setup all routes
	router.InitializeServiceRoutes(ginRouter, Client)

	// Start the HTTP server (API Gateway)
	log.Printf("API Gateway is running on port %s", config.HTTPPort)
	if err := ginRouter.Run(":" + config.HTTPPort); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
