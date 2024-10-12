// package http
package http

import (
	docs "github.com/Lozerd/shop_go/docs"
	i "github.com/Lozerd/shop_go/internal/infrastructure"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoutes(r *gin.Engine) {
	apiBasePath := i.Config.GetApiBasePath()
	docs.SwaggerInfo.BasePath = apiBasePath

    swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
    r.GET("/swagger", func(ctx *gin.Context) {
        ctx.Redirect(301, "/swagger/index.html")
    })
    r.GET("/swagger/*any", swaggerHandler)
}
