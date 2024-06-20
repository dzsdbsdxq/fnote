package service

import (
	"context"
	"errors"
	"github.com/chenmingyong0423/fnote/server/internal/category"
	apiwrap "github.com/chenmingyong0423/fnote/server/internal/pkg/web/wrap"
	"github.com/chenmingyong0423/fnote/server/internal/post"
	"github.com/chenmingyong0423/fnote/server/internal/post_like"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/chenmingyong0423/go-eventbus"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type IWebEngineService interface {
	GetWebSiteConfig(ctx context.Context) (*website_config.IndexConfig, error)
	GetCategory(ctx context.Context) ([]category.Category, error)
	GetPostList(ctx context.Context, req *post.PostRequest) ([]*post.Post, int64, error)
	GetLatestPosts(ctx context.Context, count int64) ([]*post.Post, error)
	GetPostDetailById(ctx context.Context, id string, ip string) (*post.DetailPost, error)
}

var _ IWebEngineService = (*WebEngineService)(nil)

type WebEngineService struct {
	eventBus         *eventbus.EventBus
	websiteConfigMdl *website_config.Module
	categoryMdl      *category.Module
	postMdl          *post.Module
	postLikeMdl      *post_like.Module
}

func NewWebEngineService(eventBus *eventbus.EventBus,
	websiteConfigMdl *website_config.Module,
	categoryMdl *category.Module,
	postMdl *post.Module,
	postLikeMdl *post_like.Module) *WebEngineService {
	return &WebEngineService{
		eventBus:         eventBus,
		websiteConfigMdl: websiteConfigMdl,
		categoryMdl:      categoryMdl,
		postMdl:          postMdl,
		postLikeMdl:      postLikeMdl,
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
func (s *WebEngineService) GetPostDetailById(ctx context.Context, id string, ip string) (*post.DetailPost, error) {
	postDetail, err := s.postMdl.Svc.GetPunishedPostById(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apiwrap.NewErrorResponseBody(http.StatusBadRequest, "The postId does not exist.")
		}
		return nil, err
	}
	// 查询点赞状态
	liked, err := s.postLikeMdl.Svc.GetLikeStatus(ctx, postDetail.PrimaryPost.Id, ip)
	if err != nil {
		return nil, err
	}
	return &post.DetailPost{
		PrimaryPost: postDetail.PrimaryPost,
		ExtraPost:   postDetail.ExtraPost,
		IsLiked:     liked,
	}, nil
}
