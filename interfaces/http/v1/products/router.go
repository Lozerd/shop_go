package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	as "github.com/lozerd/shop_go/application/product"
	dp "github.com/lozerd/shop_go/interfaces/dto/product"

	"github.com/lozerd/shop_go/infrastructure/db/postgres"
	"github.com/lozerd/shop_go/infrastructure/logging"
	rp "github.com/lozerd/shop_go/infrastructure/repository/product"
	sp "github.com/lozerd/shop_go/infrastructure/services/product"
)

var logger = logging.NewLogger("http", true)

// ListProducts godoc
//
//	@Summary	List all products
//	@Tags		Products
//	@Produce json
//	@Success	200	{object}	[]dp.ProductOutDTO
//	@Router		/api/v1/products/all/ [get]
func ListProducts(c *gin.Context, service as.IProductService) {
	products, err := service.List()

	if err != nil {
		logger.Error().Err(err).Msg("failed to list products")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var out []dp.ProductOutDTO
	for _, product := range products {
		out = append(out, dp.ProductOutDTO{
			ID:   product.ID,
			Name: product.Name,
		})
	}

	c.JSON(http.StatusOK, out)
}

// CreateProduct godoc
//
//	@Summary	Create a new product
//	@Tags		Products
//	@Accept		json
//	@Produce	json
//	@Param		product	body		dp.ProductInDTO	true	"Product"
//	@Success	200		{object}	dp.ProductOutDTO
//	@Router		/api/v1/products/ [post]
func CreateProduct(c *gin.Context, service as.IProductService) {
	var dto dp.ProductInDTO
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	product, err := service.Create(dto)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create product")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var out dp.ProductOutDTO
	c.JSON(200, out.From(product))
}

func RegisterRoutes(r *gin.RouterGroup) {
	db, err := postgres.GetDB()
	if err != nil {
		logger.Error().Err(err).Msg("failed to get database connection")
		return
	}

	repo := rp.NewProductRepository(db)
	service := sp.NewProductService(repo)

	productsGroup := r.Group("/products")
	productsGroup.GET("/all/", func(c *gin.Context) { ListProducts(c, service) })
	productsGroup.POST("/", func(c *gin.Context) { CreateProduct(c, service) })
}
