package swagger

import (
	"github.com/gin-gonic/gin"

	"github.com/lozerd/shop_go/docs"
	swaggerFiles "github.com/swaggo/files"
)

func NewRouter(r *gin.Engine) {
	docs.SwaggerInfo.Title = "Shop API"
	docs.SwaggerInfo.Description = "Shop API."

	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Version = "1.0"

	r.GET("/swagger/*any", WrapHandler(swaggerFiles.Handler, func(c *Config) {
		c.DeepLinking = true
		c.Filter = true
	}))
}
