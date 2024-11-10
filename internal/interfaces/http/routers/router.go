package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	v1 "github.com/Lozerd/shop_go/internal/interfaces/http/routers/v1"
)

func getBaseRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	return r
}

func SwaggerRoutes(r *gin.Engine) {
	r.Static("/swagger", "docs/swaggerui")

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/")
	})
}

func SetupRoutes(c *config.Configuration) *gin.Engine {
	r := getBaseRouter()

	SwaggerRoutes(r)

	api := r.Group(c.GetApiPrefix())
	{
		api_v1 := api.Group(c.GetApiVersion())
		{
			v1.ConfigsRoutes(api_v1)
		}
	}

	return r
}
