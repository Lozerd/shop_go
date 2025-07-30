package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/swag"

	swagger "github.com/lozerd/shop_go/interfaces/http"
	v1 "github.com/lozerd/shop_go/interfaces/http/v1"
	"github.com/lozerd/shop_go/interfaces/http/v1/products"
	"github.com/lozerd/shop_go/internal/config"
	"github.com/rs/zerolog/log"

	_ "github.com/lozerd/shop_go/docs"
)

func loadConfig() *config.Config {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not load config")
	}
	return cfg
}

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "components": {},
    "info": {"description":"{{escape .Description}}","title":"{{.Title}}","version":"{{.Version}}"},
    "externalDocs": {"description":"","url":""},
    "paths": {"/example/helloworld":{"get":{"description":"do ping","requestBody":{"content":{"application/json":{"schema":{"type":"object"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"type":"string"}}},"description":"OK"}},"summary":"ping example","tags":["example"]}}},
    "openapi": "3.1.0",
    "servers": [
        {"url":"/api/v1"}
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Title:            "",
	Description:      "do ping",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func Register() {
	fmt.Print("registered: ", SwaggerInfo.InstanceName())
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

// @title           Swagger Example API
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @openapi: 3.0.0
func main() {
	godotenv.Load()
	cfg := loadConfig()

	gin.SetMode(cfg.GinMode)
	Register()

	r := gin.Default()
	r.SetTrustedProxies(cfg.TrustedProxies)
	v1.NewRouter(r)
	swagger.NewRouter(r)

	r.Run(fmt.Sprintf(":%d", cfg.Port))
}

// products
// @Summary Get all products
// @Description Get all products
// @Tags Products
// @Router /products [get]
func GetProducts(r *gin.Engine) {
	r.GET("/products", products.GetProducts)
}
