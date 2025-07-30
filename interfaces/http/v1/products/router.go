package products

import "github.com/gin-gonic/gin"

// ShowAccount godoc
//
//	@Summary		Show an account1
//	@Description	get string by ID
//	@Tags			products
//	@Success		200	{object}	map[string]string	"application/json"
//	@Router			/api/v1/products/all/ [get]
func GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "products",
	})
}

func RegisterRoutes(r *gin.RouterGroup) {
	productsGroup := r.Group("/products")
	productsGroup.GET("/all/", GetProducts)
}
