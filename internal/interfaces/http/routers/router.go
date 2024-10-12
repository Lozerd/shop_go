package routers

import (
	"github.com/gin-gonic/gin"

	i "github.com/Lozerd/shop_go/internal/infrastructure"
	internal_http "github.com/Lozerd/shop_go/internal/interfaces/http"
	v1 "github.com/Lozerd/shop_go/internal/interfaces/http/routers/v1"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
    r.SetTrustedProxies(nil)

	internal_http.SwaggerRoutes(r)
    api := r.Group(i.Config.GetApiPrefix())
    api_v1 := api.Group(i.Config.GetApiVersion())
	v1.Routes(api_v1)

	return r
}
