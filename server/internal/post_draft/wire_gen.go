// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package post_draft

import (
	"github.com/chenmingyong0423/fnote/server/internal/post_draft/internal/repository"
	"github.com/chenmingyong0423/fnote/server/internal/post_draft/internal/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/post_draft/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/post_draft/internal/web"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitPostDraftModule(mongoDB *mongo.Database) *Model {
	postDraftDao := dao.NewPostDraftDao(mongoDB)
	postDraftRepository := repository.NewPostDraftRepository(postDraftDao)
	postDraftService := service.NewPostDraftService(postDraftRepository)
	postDraftHandler := web.NewPostDraftHandler(postDraftService)
	model := &Model{
		Svc: postDraftService,
		Hdl: postDraftHandler,
	}
	return model
}

// wire.go:

var PostDraftProviders = wire.NewSet(web.NewPostDraftHandler, service.NewPostDraftService, repository.NewPostDraftRepository, dao.NewPostDraftDao, wire.Bind(new(service.IPostDraftService), new(*service.PostDraftService)), wire.Bind(new(repository.IPostDraftRepository), new(*repository.PostDraftRepository)), wire.Bind(new(dao.IPostDraftDao), new(*dao.PostDraftDao)))
