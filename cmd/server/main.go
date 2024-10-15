package main

import (
	"github.com/Lozerd/shop_go/internal/infrastructure"
	"github.com/Lozerd/shop_go/internal/infrastructure/db"
	"github.com/Lozerd/shop_go/internal/interfaces/http/routers"
	"github.com/Lozerd/shop_go/pkg/logging"
)

// @Version 1.0.0
// @Title Shop API
// @Description Shop API designed for learning purposes
// @ContactName Lozerd
// @ContactEmail dmisha29411@gmail.com
// @ContactURL https://github.com/Lozerd
// @LicenseName MIT
// @Server http://localhost Development
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Security Jwt-access-token read write
// @SecurityScheme Jwt-access-token http bearer JSON Web Token authentication with required prefix "Bearer" for common users. Token might be fetched during API authentication flow.
func main() {
	config.LoadConfig()
	logging.InitLogging()
	db.Init()

	r := routers.SetupRoutes()
	r.Run(config.Config.GetAddr())
}
