package server

import (
	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	"github.com/Lozerd/shop_go/internal/interfaces/http/routers"
	"go.uber.org/dig"
)

type ServerDeps struct {
	dig.In

	Config *config.Configuration
}

func NewServer(d ServerDeps) {
	r := routers.SetupRoutes(d.Config)
	r.Run(d.Config.GetAddr())
}
