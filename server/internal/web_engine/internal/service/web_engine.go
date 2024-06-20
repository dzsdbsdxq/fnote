package service

import (
	"context"
	"github.com/chenmingyong0423/fnote/server/internal/category"
	"github.com/chenmingyong0423/fnote/server/internal/post"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/chenmingyong0423/go-eventbus"
)

type IWebEngineService interface {
	GetWebSiteConfig(ctx context.Context) (*website_config.IndexConfig, error)
	GetCategory(ctx context.Context) ([]category.Category, error)
	GetPostList(ctx context.Context, req *post.PostRequest) ([]*post.Post, int64, error)
	GetLatestPosts(ctx context.Context, count int64) ([]*post.Post, error)
}

var _ IWebEngineService = (*WebEngineService)(nil)

type WebEngineService struct {
	eventBus         *eventbus.EventBus
	websiteConfigMdl *website_config.Module
	categoryMdl      *category.Module
	postMdl          *post.Module
}

func NewWebEngineService(eventBus *eventbus.EventBus,
	websiteConfigMdl *website_config.Module,
	categoryMdl *category.Module,
	postMdl *post.Module) *WebEngineService {
	return &WebEngineService{
		eventBus:         eventBus,
		websiteConfigMdl: websiteConfigMdl,
		categoryMdl:      categoryMdl,
		postMdl:          postMdl,
	}
}

func (s *WebEngineService) GetWebSiteConfig(ctx context.Context) (*website_config.IndexConfig, error) {
	config, err := s.websiteConfigMdl.Svc.GetIndexConfig(ctx)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (s *WebEngineService) GetCategory(ctx context.Context) ([]category.Category, error) {
	menus, err := s.categoryMdl.Svc.GetMenus(ctx)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (s *WebEngineService) GetPostList(ctx context.Context, req *post.PostRequest) ([]*post.Post, int64, error) {
	posts, count, err := s.postMdl.Svc.GetPosts(ctx, req)
	if err != nil {
		return nil, count, err
	}
	return posts, count, err
}

func (s *WebEngineService) GetLatestPosts(ctx context.Context, count int64) ([]*post.Post, error) {
	posts, err := s.postMdl.Svc.GetLatestPosts(ctx, count)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
