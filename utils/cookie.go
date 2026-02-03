package utils

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetAuthCookies(c *gin.Context, token string, expiresIn uint) {
	isProd := os.Getenv("APP_ENV") == "production"

	c.SetCookie(
		"access_token",
		token,
		int(expiresIn*3600),
		"/",
		"",
		isProd, // Secure = true di prod
		true,   // HttpOnly
	)

	// 🔥 override SameSite karena gin default Lax
	c.Writer.Header().Add("Set-Cookie",
		"access_token="+token+
			"; Path=/"+
			"; Max-Age="+string(rune(expiresIn*3600))+
			"; HttpOnly"+
			func() string {
				if isProd {
					return "; Secure; SameSite=None"
				}
				return "; SameSite=Lax"
			}(),
	)
}


func ClearAuthCookies(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", true, true)
}
