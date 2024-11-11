package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liju-github/EcommerceApiGatewayService/proto/user"
)

func BanCheckMiddleware(userClient user.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("USERID")
		if !exists {
			c.JSON(401, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		resp, err := userClient.CheckBan(c, &user.CheckBanRequest{
			UserID: userID.(string),
		})
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to check ban status"})
			c.Abort()
			return
		}

		if resp.BanStatus {
			c.JSON(403, gin.H{"error": "User is banned"})
			c.Abort()
			return
		}

		c.Next()
	}
}
