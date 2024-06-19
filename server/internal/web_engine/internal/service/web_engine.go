package service

import (
	"context"
	"github.com/chenmingyong0423/fnote/server/internal/category"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/chenmingyong0423/go-eventbus"
)

type IWebEngineService interface {
	GetWebSiteConfig(ctx context.Context) (*website_config.WebsiteConfig, error)
	GetCategory(ctx context.Context) ([]category.Category, error)
}

var _ IWebEngineService = (*WebEngineService)(nil)

type WebEngineService struct {
	eventBus             *eventbus.EventBus
	websiteConfigService *website_config.Module
	categoryService      *category.Module
}

func NewWebEngineService(eventBus *eventbus.EventBus,
	websiteConfigService *website_config.Module,
	categoryService *category.Module) *WebEngineService {
	return &WebEngineService{
		eventBus:             eventBus,
		websiteConfigService: websiteConfigService,
		categoryService:      categoryService,
	}
}
func (s *WebEngineService) GetWebSiteConfig(ctx context.Context) (*website_config.WebsiteConfig, error) {
	config, err := s.websiteConfigService.Svc.GetIndexConfig(ctx)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (s *WebEngineService) GetCategory(ctx context.Context) ([]category.Category, error) {
	menus, err := s.categoryService.Svc.GetMenus(ctx)
	if err != nil {
		return nil, err
	}
	return menus, nil
}
