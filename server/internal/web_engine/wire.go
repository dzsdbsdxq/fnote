//go:build wireinject

package web_engine

import (
	"github.com/chenmingyong0423/fnote/server/internal/category"
	"github.com/chenmingyong0423/fnote/server/internal/post"
	"github.com/chenmingyong0423/fnote/server/internal/web_engine/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/web_engine/internal/web"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/chenmingyong0423/go-eventbus"
	"github.com/google/wire"
)

var ConfigProviders = wire.NewSet(web.NewWebEngineHandler, service.NewWebEngineService,
	wire.Bind(new(service.IWebEngineService), new(*service.WebEngineService)),
)

func InitWebEngineModule(eventBus *eventbus.EventBus,
	websiteConfigMdl *website_config.Module,
	categoryMdl *category.Module, postMdl *post.Module) *Module {
	panic(wire.Build(
		ConfigProviders,
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
