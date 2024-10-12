package http

import (
	"github.com/gin-gonic/gin"
)

func SwaggerRoutes(r *gin.Engine) {
    r.Static("/swagger", "docs/swaggerui")
}
