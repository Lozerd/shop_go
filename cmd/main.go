package main

import (
	"github.com/Lozerd/shop_go/internal/infrastructure/dependencies"
	"github.com/Lozerd/shop_go/internal/interfaces/http/server"
)

// @Version 1.0.0
// @Title Shop API
// @Description Shop API designed for learning purposes
// @ContactName Lozerd
// @ContactEmail dmisha29411@gmail.com
// @ContactURL https://github.com/Lozerd
// @LicenseName MIT
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Security Jwt-access-token read write
// @SecurityScheme Jwt-access-token http bearer JSON Web Token authentication with required prefix "Bearer" for common users. Token might be fetched during API authentication flow.
func main() {
	// logging.InitLogging()

    c := dependencies.Init()
    err := c.Invoke(server.NewServer)
    if err != nil {
        panic(err)
    }
}
