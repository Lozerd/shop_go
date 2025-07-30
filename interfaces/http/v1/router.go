package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozerd/shop_go/interfaces/http/v1/products"
)

func NewRouter(r *gin.RouterGroup) {
	products.RegisterRoutes(r)
}
