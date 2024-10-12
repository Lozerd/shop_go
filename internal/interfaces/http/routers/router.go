package routers

import (
	"github.com/gin-gonic/gin"

	i "github.com/Lozerd/shop_go/internal/infrastructure"
	s "github.com/Lozerd/shop_go/internal/interfaces/http/swagger"
	v1 "github.com/Lozerd/shop_go/internal/interfaces/http/routers/v1"
)

func getBaseRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
    return r
}

func SetupRoutes() *gin.Engine {
    r := getBaseRouter()

	s.SwaggerRoutes(r)

	api := r.Group(i.Config.GetApiPrefix())
    {
        api_v1 := api.Group(i.Config.GetApiVersion())
        {
            v1.Routes(api_v1)
        }
    }

	return r
}
