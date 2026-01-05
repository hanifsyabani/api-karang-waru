package utils

import "github.com/gin-gonic/gin"

func SetAuthCookies(
	c *gin.Context,
	access_token string,
	expiresIn uint,
) {
	maxAge := int(expiresIn * 3600)
	secure := gin.Mode() == gin.ReleaseMode

	c.SetCookie(
		"access_token", // nama cookie
		access_token,   // value = token JWT
		maxAge,         // detik (jam → detik)
		"/",            // path
		"",             // domain (kosong = current domain)
		secure,           // secure (harus HTTPS kalau true)
		true,           // httpOnly (tidak bisa diakses via JS)
	)
	
}
