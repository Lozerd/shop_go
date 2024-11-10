package repositories

import (
	"github.com/Lozerd/shop_go/internal/infrastructure/config"
)

type IConfigsRepository interface {
	Get() (*config.Configuration, error)
}
