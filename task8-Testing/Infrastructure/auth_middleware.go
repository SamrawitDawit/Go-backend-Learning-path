package infrastructure

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(jwt_service Jwt_interface) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()

		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.JSON(401, gin.H{"Error": "Authorization header is required"})
			context.Abort()
			return
		}

		authPart := strings.Split(authHeader, " ")

		if len(authPart) != 2 || strings.ToLower(authPart[0]) != "bearer" {
			context.JSON(401, gin.H{"message": "Invalid Authorization header"})
			context.Abort()
			return
		}

		token, err := jwt_service.CheckToken(authPart[1])

		if token == nil || !token.Valid {
			errMsg := "Invalid or expired token"
			if err != nil {
				errMsg = err.Error()
			}
			context.JSON(401, gin.H{"error": errMsg})
			context.Abort()
			return
		}
		claims, ok := jwt_service.FindClaim(token)
		if !ok {
			context.JSON(401, gin.H{"error": "Invalid token claims"})
			context.Abort()
			return
		}
		role, ok := claims["role"]
		if !ok {
			context.JSON(401, gin.H{"error": "Invalid role field in token"})
			context.Abort()
			return
		}
		context.Set("role", role)
	}
}
func AdminMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()
		role, exists := context.Get("role")
		if !exists || role != "admin" {
			context.JSON(403, gin.H{"message": "Sorry, you are not eligible to do this.", "your role is": role})
			context.Abort()
			return
		}
	}
}
