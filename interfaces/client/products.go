package client

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"

	as "github.com/lozerd/shop_go/application/product"
	"github.com/lozerd/shop_go/infrastructure/db/postgres"
	"github.com/lozerd/shop_go/infrastructure/logging"
	rp "github.com/lozerd/shop_go/infrastructure/repository/product"
	ps "github.com/lozerd/shop_go/infrastructure/services/product"
	dp "github.com/lozerd/shop_go/interfaces/dto/product"
)

var logger = logging.NewLogger("client", true)

func getProductService() (as.IProductService, error) {
	db, err := postgres.GetDB()
	if err != nil {
		return nil, err
	}

	service := ps.NewProductService(rp.NewProductRepository(db))
	return service, nil
}

func CreateProductView(c *gin.Context) {
	form := dp.ProductInDTO{}
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	service, err := getProductService()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	product, err := service.Create(form)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "product", product)
}

func DeleteProductView(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	service, err := getProductService()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = service.Delete(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
func NewProductRouter(r *gin.Engine) {
	products := r.Group("/products/")
	products.POST("/products/", CreateProductView)
	products.DELETE("/products/:id/", DeleteProductView)

}
