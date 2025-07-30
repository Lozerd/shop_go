package swagger

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"

)

func NewRouter(r *gin.Engine) {
	// docs.SwaggerInfo.Title = "Shop API"
	// docs.SwaggerInfo.Description = "Shop API"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.BasePath = "/api"

	r.GET("/swagger/*any", WrapHandler(swaggerFiles.Handler))
}
