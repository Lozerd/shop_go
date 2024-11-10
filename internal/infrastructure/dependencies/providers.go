package dependencies

import (
	"log"

	appServices "github.com/Lozerd/shop_go/internal/app/services"
	"github.com/Lozerd/shop_go/internal/domain/services"
	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	"github.com/Lozerd/shop_go/internal/infrastructure/db"
	"github.com/Lozerd/shop_go/internal/infrastructure/db/repositories"
	"go.uber.org/dig"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

var container *dig.Container

// Little hack, to access Container from anywhere
func loadContainer() *dig.Container {
    if container == nil {
        container = dig.New()
    }
    return container
}

func GetContainer() *dig.Container {
	if container == nil {
		loadContainer()
	}
    return container
}

func Init() (c *dig.Container) {
	c = GetContainer()
	ProvideConfig(c)
	ProvideDB(c)
	ProvideRepositories(c)
	ProvideServices(c)
	return
}

func ProvideConfig(c *dig.Container) {
	err := c.Provide(
		func() *config.Configuration {
			return config.GetConfig()
		},
	)
	checkErr(err)
}

func ProvideDB(c *dig.Container) {
	err := c.Provide(func(c *config.Configuration) *db.Connection {
		return db.NewConnection(c)
	})
	checkErr(err)
}

func ProvideRepositories(c *dig.Container) {
	err := c.Provide(func() *repositories.ConfigsRepository {
		return repositories.NewConfigsRepository()
	})
	checkErr(err)
}

func ProvideServices(c *dig.Container) {
	err := c.Provide(func(r *repositories.ConfigsRepository) *services.ConfigsService {
		return services.NewConfigsService(r)
	})
	checkErr(err)
	err = c.Provide(func(s *services.ConfigsService) *appServices.ConfigsAppService {
		return appServices.NewConfigsAppService(s)
	})
	checkErr(err)
}
