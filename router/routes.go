package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liju-github/EcommerceApiGatewayService/clients"
	"github.com/liju-github/EcommerceApiGatewayService/controller"
	"github.com/liju-github/EcommerceApiGatewayService/middleware"
	// "github.com/liju-github/EcommerceApiGatewayService/proto/admin"
	// "github.com/liju-github/EcommerceApiGatewayService/proto/content"
	// "github.com/liju-github/EcommerceApiGatewayService/proto/notification"
	"github.com/liju-github/EcommerceApiGatewayService/proto/user"
)

// InitializeServiceRoutes initializes gRPC clients for each service, creates controllers, 
// and configures routes for each service.
func InitializeServiceRoutes(router *gin.Engine, Client *clients.ClientConnections) {
	// User Client setup
	userClient := user.NewUserServiceClient(Client.ConnUser)
	userController := controller.NewUserController(userClient)
	SetupUserRoutes(router, userController)

	// // Content Client setup
	// contentClient := content.NewContentServiceClient(Client.ConnContent)
	// contentController := controller.NewContentController(contentClient)
	// SetupContentRoutes(router, contentController)

	// // Admin Client setup
	// adminClient := admin.NewAdminServiceClient(Client.ConnAdmin)
	// adminController := controller.NewAdminController(adminClient)
	// SetupAdminRoutes(router, adminController)

	// // Notification Client setup
	// notificationClient := notification.NewNotificationServiceClient(Client.ConnNotification)
	// notificationController := controller.NewNotificationController(notificationClient)
	// SetupNotificationRoutes(router, notificationController)
}

// SetupUserRoutes configures routes for User-related operations
func SetupUserRoutes(router *gin.Engine, userController *controller.UserController) {
	// Public routes
	router.POST("/register", userController.RegisterHandler)
	router.POST("/login", userController.LoginHandler)

	// Protected routes with JWT middleware
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware)
	{
		protected.GET("/profile", userController.GetProfileHandler)
		protected.POST("/update-profile", userController.UpdateProfileHandler)
	}
}


