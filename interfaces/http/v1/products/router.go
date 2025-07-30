package products

import "github.com/gin-gonic/gin"

func GetProducts(c *gin.Context) {

}
func NewRouter(r *gin.Engine) {
	// @Summary Get all products
	// @Description Get all products
	// @Tags Products
	// @Accept json
	// @Produce json
	// @Success 200 {object} Product
	// @Router /products [get]
	r.GET("/products", GetProducts)
}

