package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/lozerd/shop_go/internal/config"
	"github.com/rs/zerolog/log"
)

func main() {
	godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not load config")
	}

	gin.SetMode(cfg.GinMode)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.Run(":8080")
}
