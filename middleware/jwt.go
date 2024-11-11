package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liju-github/EcommerceApiGatewayService/utils"
)

// JWTAuthMiddleware is a middleware to check and authorize JWT tokens.
func JWTAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	userID, role, err := utils.ParseJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"+err.Error()})
		c.Abort()
		return
	}

	// Set userID and role in the context
	c.Set("USERID", userID)
	c.Set("ROLE", role)

	fmt.Println("JWT userid:", c.GetString("USERID"), "role:", c.GetString("ROLE"));
	c.Next()
}


func AdminAccess(c *gin.Context)  {
	role := c.GetString("ROLE")
	if role != "ADMIN"{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	}
	c.Next()
}