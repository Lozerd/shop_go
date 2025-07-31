package client

import (
	"github.com/gin-gonic/gin"
	"github.com/lozerd/shop_go/interfaces/http/middleware"
	"github.com/lozerd/shop_go/internal/utils"
)

func NewRouter(r *gin.Engine) {
	r.Use(middleware.Favicon)
	r.Static("/static", "./assets")

	templateFiles, err := utils.LoadTemplates(r, "infrastructure/templates")
	if err != nil {
		logger.Error().Err(err).Msg("failed to load templates")
		return
	} else {
		r.LoadHTMLFiles(templateFiles...)
	}

	r.GET("/", IndexView)
	r.GET("/about/", AboutView)
	r.GET("/recipes/", RecipesView)
	NewProductRouter(r)
}
