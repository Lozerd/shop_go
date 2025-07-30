package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/lozerd/shop_go/infrastructure/db/postgres"
	swagger "github.com/lozerd/shop_go/interfaces/http"
	routerV1 "github.com/lozerd/shop_go/interfaces/http/v1"
	"github.com/lozerd/shop_go/internal/config"

	_ "github.com/lozerd/shop_go/docs"
)

// @title						Shop API
// @description				Shop API.
// @securityDefinitions.basic	BasicAuth
func main() {
	godotenv.Load()
	cfg := config.GetConfig()
	postgres.Ping()
	fmt.Println(postgres.GetDSN())

	gin.SetMode(cfg.GinMode)

	r := gin.Default()
	r.SetTrustedProxies(cfg.TrustedProxies)
	swagger.NewRouter(r)

	api := r.Group(cfg.GetApiPrefix())
	{
		v1 := api.Group("v1")
		{
			routerV1.NewRouter(v1)
		}
	}

	r.Run(cfg.GetAddr())
}
