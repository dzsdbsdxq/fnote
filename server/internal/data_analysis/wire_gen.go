// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package data_analysis

import (
	service3 "github.com/chenmingyong0423/fnote/server/internal/comment/service"
	service2 "github.com/chenmingyong0423/fnote/server/internal/count_stats/service"
	"github.com/chenmingyong0423/fnote/server/internal/data_analysis/internal/web"
	"github.com/chenmingyong0423/fnote/server/internal/post_like"
	"github.com/chenmingyong0423/fnote/server/internal/visit_log/service"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitDataAnalysisModule(mongoDB *mongo.Database, vlServ service.IVisitLogService, csServ service2.ICountStatsService, posLikeModule *post_like.Module, commentServ service3.ICommentService) *Module {
	iPostLikeService := posLikeModule.Svc
	dataAnalysisHandler := web.NewDataAnalysisHandler(vlServ, csServ, iPostLikeService, commentServ)
	module := &Module{
		Hdl: dataAnalysisHandler,
	}
	return module
}

// wire.go:

var DataAnalysisProviders = wire.NewSet(web.NewDataAnalysisHandler)