package utils

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SetAuthCookies(c *gin.Context, token string, expiresIn uint) {
	isProd := os.Getenv("APP_ENV") == "production"
	maxAge := int(expiresIn * 3600)

	sameSite := "None"
	secureFlag := ""

	if isProd {
		secureFlag = "; Secure"
	}

	cookieValue := fmt.Sprintf(
		"access_token=%s; Path=/; Max-Age=%d; HttpOnly; SameSite=%s%s",
		token,
		maxAge,
		sameSite,
		secureFlag,
	)

	c.Writer.Header().Set("Set-Cookie", cookieValue)
}

func ClearAuthCookies(c *gin.Context) {
	isProd := os.Getenv("APP_ENV") == "production"

	secureFlag := ""
	if isProd {
		secureFlag = "; Secure"
	}

	cookieValue := fmt.Sprintf(
		"access_token=; Path=/; Max-Age=-1; HttpOnly; SameSite=None%s",
		secureFlag,
	)

	c.Writer.Header().Set("Set-Cookie", cookieValue)
}
