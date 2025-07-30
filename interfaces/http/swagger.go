package swagger

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
)

func NewRouter(r *gin.Engine) {
	r.GET("/swagger/*any", WrapHandler(swaggerFiles.Handler, func(c *Config) {
		c.DeepLinking = true
		c.Filter = true
	}))
}
