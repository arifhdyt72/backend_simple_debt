package middleware

import (
	"backend_test_debt/helper"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddlware(jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			errorMessage := gin.H{"error": "Unauthorized"}
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			errorMessage := gin.H{"error": err.Error()}
			response := helper.ApiResponse("Unauthorized 1", http.StatusUnauthorized, "error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			errorMessage := gin.H{"error": "Unauthorized"}
			response := helper.ApiResponse("Unauthorized 2", http.StatusUnauthorized, "error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		c.Set("currentUser", userID)
		c.Next()
	}
}
