package v1

import (
	"github.com/Lozerd/shop_go/internal/app/services"
	"github.com/Lozerd/shop_go/internal/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

// @Description Get config struct.
// @Title Get config struct.
// @Route /api/v1/configs/ [get]
func GetConfigs(service *services.ConfigsAppService, ctx *gin.Context) {
    c, err := service.Get()
    if err != nil {
        ctx.JSON(500, err)
        return
    }

	ctx.JSON(200, c)
}

func ConfigsRoutes(v1 *gin.RouterGroup) {
	configs := v1.Group("/configs")
	{
		configs.GET("/", func(ctx *gin.Context) {
			c := dependencies.GetContainer()
			c.Invoke(func(service *services.ConfigsAppService) {
				GetConfigs(service, ctx)
			})
		})
	}
}
