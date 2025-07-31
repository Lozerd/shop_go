package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutView(c *gin.Context) {
	c.HTML(http.StatusOK, "about", gin.H{})
}
