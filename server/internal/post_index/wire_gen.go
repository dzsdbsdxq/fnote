// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package post_index

import (
	"github.com/chenmingyong0423/fnote/server/internal/post_index/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/post_index/internal/web"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitPostIndexModule(cfgServ *website_config.Model) *Model {
	baiduService := service.NewBaiduService()
	iWebsiteConfigService := cfgServ.Svc
	postIndexService := service.NewPostIndexService(baiduService, iWebsiteConfigService)
	postIndexHandler := web.NewPostIndexHandler(postIndexService)
	model := &Model{
		Svc: postIndexService,
		Hdl: postIndexHandler,
	}
	return model
}

// wire.go:

var PostIndexProviders = wire.NewSet(web.NewPostIndexHandler, service.NewPostIndexService, service.NewBaiduService, wire.Bind(new(service.IPostIndexService), new(*service.PostIndexService)))
