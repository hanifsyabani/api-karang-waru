package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetAuthCookies(
	c *gin.Context,
	access_token string,
	expiresIn uint,
) {
	maxAge := int(expiresIn * 3600)

	c.SetSameSite(http.SameSiteNoneMode)

	c.SetCookie(
		"access_token", // nama cookie
		access_token,   // value = token JWT
		maxAge,         // detik (jam → detik)
		"/",            // path
		"",             // domain (kosong = current domain)
		true,           // secure (harus HTTPS kalau true)
		true,           // httpOnly (tidak bisa diakses via JS)
	)

}

func ClearAuthCookies(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", true, true)
}
