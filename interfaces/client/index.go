package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexView(c *gin.Context) {
	service, err := getProductService()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	products, err := service.List()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context := gin.H{"products": products}
	c.HTML(http.StatusOK, "main", context)
}
