package main

import (
	i "github.com/Lozerd/shop_go/internal/infrastructure"
	db "github.com/Lozerd/shop_go/internal/infrastructure/db"
	"github.com/Lozerd/shop_go/internal/interfaces/http/routers"
	l "github.com/Lozerd/shop_go/pkg/logging"
)

// @Version 1.0.0
// @Title Backend API
// @Description API usually works as expected. But sometimes its not true.
// @ContactName Parvez
// @ContactEmail abce@email.com
// @ContactURL http://someurl.oxox
// @TermsOfServiceUrl http://someurl.oxox
// @LicenseName MIT
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Server http://www.fake.com Server-1
// @Server http://www.fake2.com Server-2
// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader http bearer Input your token
func main() {
	i.LoadConfig()
	l.InitLogging()
	db.Init()

	r := routers.SetupRoutes()
	r.Run(i.Config.GetAddr())
}
