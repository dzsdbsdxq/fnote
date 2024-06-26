// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package count_stats

import (
	"github.com/chenmingyong0423/fnote/server/internal/count_stats/internal/repository"
	"github.com/chenmingyong0423/fnote/server/internal/count_stats/internal/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/count_stats/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/count_stats/internal/web"
	"github.com/chenmingyong0423/go-eventbus"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitCountStatsModule(mongoDB *mongo.Database, eventbus2 *eventbus.EventBus) *Module {
	countStatsDao := dao.NewCountStatsDao(mongoDB)
	countStatsRepository := repository.NewCountStatsRepository(countStatsDao)
	countStatsService := service.NewCountStatsService(countStatsRepository, eventbus2)
	countStatsHandler := web.NewCountStatsHandler(countStatsService)
	module := &Module{
		Svc: countStatsService,
		Hdl: countStatsHandler,
	}
	return module
}

// wire.go:

var CountStatsProviders = wire.NewSet(web.NewCountStatsHandler, service.NewCountStatsService, repository.NewCountStatsRepository, dao.NewCountStatsDao, wire.Bind(new(service.ICountStatsService), new(*service.CountStatsService)), wire.Bind(new(repository.ICountStatsRepository), new(*repository.CountStatsRepository)), wire.Bind(new(dao.ICountStatsDao), new(*dao.CountStatsDao)))
