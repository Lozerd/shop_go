package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecipesView(c *gin.Context) {
	c.HTML(http.StatusOK, "recipes", gin.H{})
}
