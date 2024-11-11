package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liju-github/EcommerceApiGatewayService/clients"
	"github.com/liju-github/EcommerceApiGatewayService/controller"
	"github.com/liju-github/EcommerceApiGatewayService/middleware"

	// "github.com/liju-github/EcommerceApiGatewayService/proto/admin"
	"github.com/liju-github/EcommerceApiGatewayService/proto/content"
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

	// Content Client setup
	contentClient := content.NewContentServiceClient(Client.ConnContent)
	contentController := controller.NewContentController(contentClient)
	SetupContentRoutes(router, contentController, userClient)

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

	// userProtected routes with JWT middleware
	userProtected := router.Group("/")
	userProtected.Use(middleware.JWTAuthMiddleware)
	{
		userProtected.GET("/profile", userController.GetProfileHandler)
		userProtected.PATCH("/update-profile", userController.UpdateProfileHandler)
	}

}

func SetupContentRoutes(router *gin.Engine, contentController *controller.ContentController, userClient user.UserServiceClient) {
	// Public routes that don't require authentication
	public := router.Group("/content")
	{
		// Read-only operations
		public.GET("/questions/search", contentController.SearchHandler)
		public.GET("/questions/tags", contentController.GetQuestionsByTagsHandler)
		public.GET("/questions/word", contentController.GetQuestionsByWordHandler)   //
		public.GET("/question/", contentController.GetQuestionByIDHandler)           //
		public.GET("/questions/user", contentController.GetQuestionsByUserIDHandler) //
		public.GET("/feed", contentController.GetUserFeedHandler)                    //
	}

	// userProtected routes that require both authentication and ban check
	userProtected := router.Group("/content")
	userProtected.Use(middleware.JWTAuthMiddleware, middleware.BanCheckMiddleware(userClient))
	{
		// Write operations for questions
		userProtected.POST("/question", contentController.PostQuestionHandler)
		userProtected.DELETE("/question", contentController.DeleteQuestionHandler)
		userProtected.POST("/question/mark-answered", contentController.MarkQuestionAsAnsweredHandler)
		userProtected.POST("/question/flag", contentController.FlagQuestionHandler)

		// Write operations for answers
		userProtected.POST("/answer", contentController.PostAnswerHandler)
		userProtected.DELETE("/answer", contentController.DeleteAnswerHandler)
		userProtected.POST("/answer/flag", contentController.FlagAnswerHandler)

		// userProtected.PUT("/answer/upvote", contentController.UpvoteAnswerHandler)
		// userProtected.PUT("/answer/downvote", contentController.DownvoteAnswerHandler)
		// Write operations for tags
		// userProtected.POST("/tag", contentController.AddTagHandler)
		// userProtected.DELETE("/tag", contentController.RemoveTagHandler)
	}
}

// func SetupAdminRoutes(router *gin.Engine, adminController *controller.AdminController,) {
// 	adminProtected := router.Group("/admin")
// 	// adminProtected.Use(middleware.JWTAuthMiddleware,middleware.AdminAccess)
// 	{
// 		adminProtected.GET("/answers/flagged", contentController.GetAllFlaggedAnswers)
// 		adminProtected.GET("/questions/flagged", contentController.GetAllFlaggedQuestions)


// 		adminProtected.POST("/user/unban",adminController.UnBanAdminHandler)
// 		adminProtected.POST("/user/ban",adminController.BanAdminHandler)
// 	}
// }
