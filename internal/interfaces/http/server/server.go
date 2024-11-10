package server

import (
	"net/http"

	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	"github.com/Lozerd/shop_go/internal/interfaces/http/routers"
	"go.uber.org/dig"
)

type ServerDeps struct {
	dig.In

	Config *config.Configuration
}

func NewServer(d ServerDeps) error {
	server := &http.Server{
		Addr:    d.Config.GetServerAddr(),
		Handler: routers.SetupRoutes(d.Config),
	}
	return server.ListenAndServe()
}
