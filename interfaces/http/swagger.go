package swagger

import (
	"github.com/gin-gonic/gin"

	"github.com/lozerd/shop_go/infrastructure/templates/swagger"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/swag/v2"
)

var docs = &swag.Spec{
	Version:          "1.0",
	Title:            "Shop API",
	Description:      "Shop API.",
	InfoInstanceName: "swagger",
	Host:             "localhost:8080",
	Schemes:          []string{"http"},
	BasePath:         "/",
	SwaggerTemplate:  templates.SwaggerJsonTpl,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() { swag.Register("swagger", docs) }

func NewRouter(r *gin.Engine) {
	r.GET("/swagger/*any", WrapHandler(swaggerFiles.Handler, func(c *Config) {
		c.DeepLinking = true
		c.Filter = true
	}))
}
