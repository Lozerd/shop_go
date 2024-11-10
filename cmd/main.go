package main

import (
	"log"

	"github.com/Lozerd/shop_go/internal/infrastructure/dependencies"
	"github.com/Lozerd/shop_go/internal/interfaces/http/client"
	"github.com/Lozerd/shop_go/internal/interfaces/http/server"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

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
    g.Go(func() error {
        return c.Invoke(server.NewServer)
    })
    g.Go(func() error {
        return c.Invoke(client.NewClient)
    })

    if err := g.Wait(); err != nil {
        log.Fatal(err)
    }
}
