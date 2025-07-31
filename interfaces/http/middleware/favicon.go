package middleware

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Favicon(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.Next()
		return
	}

	if c.Request.URL.Path != "/favicon.ico" {
		c.Next()
		return
	}

	if _, err := os.Stat("assets"); err != nil {
		if os.Mkdir("assets", 0755) != nil {
			c.Next()
			return
		}
	}

	c.File(filepath.Join("assets", "favicon.ico"))
}
