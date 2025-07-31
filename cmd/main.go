package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/lozerd/shop_go/infrastructure/db/postgres"
	"github.com/lozerd/shop_go/interfaces/client"
	swagger "github.com/lozerd/shop_go/interfaces/http"
	routerV1 "github.com/lozerd/shop_go/interfaces/http/v1"
	"github.com/lozerd/shop_go/internal/config"
)

func main() {
	godotenv.Load()
	cfg := config.GetConfig()
	postgres.Ping()
	postgres.Migrate()

	gin.SetMode(cfg.GinMode)

	r := gin.Default()
	r.SetTrustedProxies(cfg.TrustedProxies)

	client.NewRouter(r)
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
