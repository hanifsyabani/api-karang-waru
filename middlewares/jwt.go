package middlewares

import (
	"api-karang-waru/config"
	"api-karang-waru/responses"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware untuk validasi token Supabase
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, responses.APIResponse{
				Code: "UNAUTHORIZED",
				Message: "Missing Authorization header",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, responses.APIResponse{
				Code: "UNAUTHORIZED",
				Message: "Invalid Authorization header format",
			})
			c.Abort()
			return
		}

		secret := config.GetEnv("SUPABASE_JWT_SECRET", "")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, responses.APIResponse{
				Code:    "UNAUTHORIZED",
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		// simpan klaim ke context biar bisa dipakai di handler
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user", claims)
		}

		c.Next()
	}
}
