package services

import (
	repos "github.com/Lozerd/shop_go/internal/domain/repositories"
	"github.com/Lozerd/shop_go/internal/infrastructure/config"
)

type ConfigsService struct {
	repo repos.IConfigsRepository
}

func NewConfigsService(repo repos.IConfigsRepository) *ConfigsService {
	return &ConfigsService{repo: repo}
}

func (s *ConfigsService) Get() (*config.Configuration, error) {
    return s.repo.Get()
}
