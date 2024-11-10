package client

import (
	"net/http"

	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func getBaseClient(c *config.Configuration) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.New(c.CorsConfig))
	return r
}

type ClientDeps struct {
    dig.In

    Config *config.Configuration
}

func NewClient(d ClientDeps) error {
	client := &http.Server{
		Addr:    d.Config.GetClientAddr(),
		Handler: getBaseClient(d.Config),
	}
	return client.ListenAndServe()
}
