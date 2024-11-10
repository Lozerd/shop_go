package services

import (
	"github.com/Lozerd/shop_go/internal/app/dtos"
	"github.com/Lozerd/shop_go/internal/domain/services"
	"go.uber.org/dig"
)

type ConfigsAppService struct {
	service *services.ConfigsService
}

type ConfigsAppServiceDeps struct {
    dig.In

    Service *services.ConfigsService
}

func NewConfigsAppService(service *services.ConfigsService) *ConfigsAppService {
    return &ConfigsAppService{service: service}
}

func (s *ConfigsAppService) Get() (*dtos.ConfigurationDTO, error) {
    c, err := s.service.Get()
    if err != nil {
        return &dtos.ConfigurationDTO{}, err
    }

    return &dtos.ConfigurationDTO{
        ApiVersion: c.ApiVersion,
        ApiPrefix: c.ApiPrefix,
    }, nil
}
