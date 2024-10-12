package v1

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Ping(ctx *gin.Context) {
    ctx.String(200, "ok")
}

func Routes(r* gin.RouterGroup) {
    products := r.Group("/products")
    products.GET("/ping", Ping)
}
