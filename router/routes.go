// router/router.go
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liju-github/EcommerceApiGatewayService/controller"
	"github.com/liju-github/EcommerceApiGatewayService/middleware"
)

func SetupRoutes(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", userController.RegisterHandler)
	router.POST("/login", userController.LoginHandler)
	// router.POST("/verify-email", userController.VerifyEmailHandler)
	

	// Protected routes with JWT middleware
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware)
	{
		protected.GET("/profile", userController.GetProfileHandler)
		protected.POST("/update-profile",userController.UpdateProfileHandler)
	}

	return router
}
