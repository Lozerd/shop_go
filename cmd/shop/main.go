package main

import (
	c "github.com/Lozerd/shop_go/internal/infrastructure"
	l "github.com/Lozerd/shop_go/pkg/logging"
    db "github.com/Lozerd/shop_go/internal/infrastructure/db"
	"github.com/gin-gonic/gin"
)

func main() {
	c.LoadConfig()
	l.InitLogging()
    db.Init()

	r := gin.Default()
	r.Run(c.Config.GetAddr())
}
