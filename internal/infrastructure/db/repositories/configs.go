package repositories

import (
	"github.com/Lozerd/shop_go/internal/infrastructure/config"
)

type ConfigsRepository struct{}

func NewConfigsRepository() *ConfigsRepository {
	return &ConfigsRepository{}
}

func (r *ConfigsRepository) Get() (*config.Configuration, error) {
	return config.GetConfig(), nil
}
