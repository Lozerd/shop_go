package main

import (
	i "github.com/Lozerd/shop_go/internal/infrastructure"
	db "github.com/Lozerd/shop_go/internal/infrastructure/db"
	"github.com/Lozerd/shop_go/internal/interfaces/http/routers"
	l "github.com/Lozerd/shop_go/pkg/logging"
)

func main() {
	i.LoadConfig()
	l.InitLogging()
	db.Init()

	r := routers.SetupRoutes()
	r.Run(i.Config.GetAddr())
}
