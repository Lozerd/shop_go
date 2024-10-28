package v1

import (
	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	"github.com/gin-gonic/gin"
)

// @Description Get config struct.
// @Title Get config struct.
// @Route /api/v1/configs/ [get]
func GetConfigs(ctx *gin.Context) {
	cfg := config.GetConfig()
	ctx.JSON(200, &ConfigurationDTO{})
}

func ConfigsRoutes(v1 *gin.RouterGroup) {
	configs := v1.Group("/configs")
	{
		// @Title Get config struct.
		// @Description Get config struct.
		// @Route api/v1/configs/ [get]
		configs.GET("/", GetConfigs)
	}
}
