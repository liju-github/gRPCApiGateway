package main

import (
	"log"

	"github.com/liju-github/EcommerceApiGatewayService/controller"
	"github.com/liju-github/EcommerceApiGatewayService/proto/user"
	"github.com/liju-github/EcommerceApiGatewayService/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establish the gRPC connection to the server
	conn, err := grpc.NewClient("localhost:50000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new user service client using the established connection
	userClient := user.NewUserServiceClient(conn)

	// Initialize the user controller with the user client
	userController := controller.NewUserController(userClient)

	// Set up the routes using the user controller
	r := router.SetupRoutes(userController)

	// Start the HTTP server (API Gateway)
	log.Println("API Gateway is running on port 5000")
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
