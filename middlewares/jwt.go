package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ambil dari cookie, bukan header
        tokenString, err := c.Cookie("access_token")
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "code":    "UNAUTHORIZED",
                "message": "Missing access token",
            })
            return
        }

        // validasi JWT seperti biasa
        token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "code":    "UNAUTHORIZED",
                "message": "Invalid or expired token",
            })
            return
        }

        c.Next()
    }
}

