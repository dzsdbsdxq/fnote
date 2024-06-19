// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chenmingyong0423/fnote/server/internal/aggregate_post"
	"github.com/chenmingyong0423/fnote/server/internal/backup"
	"github.com/chenmingyong0423/fnote/server/internal/category"
	"github.com/chenmingyong0423/fnote/server/internal/comment"
	"github.com/chenmingyong0423/fnote/server/internal/count_stats"
	"github.com/chenmingyong0423/fnote/server/internal/data_analysis"
	"github.com/chenmingyong0423/fnote/server/internal/email"
	"github.com/chenmingyong0423/fnote/server/internal/file"
	"github.com/chenmingyong0423/fnote/server/internal/friend"
	"github.com/chenmingyong0423/fnote/server/internal/global"
	"github.com/chenmingyong0423/fnote/server/internal/ioc"
	"github.com/chenmingyong0423/fnote/server/internal/message"
	"github.com/chenmingyong0423/fnote/server/internal/message_template"
	"github.com/chenmingyong0423/fnote/server/internal/post"
	"github.com/chenmingyong0423/fnote/server/internal/post_draft"
	"github.com/chenmingyong0423/fnote/server/internal/post_index"
	"github.com/chenmingyong0423/fnote/server/internal/post_like"
	"github.com/chenmingyong0423/fnote/server/internal/post_visit"
	"github.com/chenmingyong0423/fnote/server/internal/tag"
	"github.com/chenmingyong0423/fnote/server/internal/visit_log"
	"github.com/chenmingyong0423/fnote/server/internal/web_engine"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func initializeApp() (*gin.Engine, error) {
	database := ioc.NewMongoDB()
	eventBus := ioc.NewEventBus()
	module := file.InitFileModule(database, eventBus)
	fileHandler := module.Hdl
	categoryModule := category.InitCategoryModule(database, eventBus)
	categoryHandler := categoryModule.Hdl
	emailModule := email.InitEmailModule(database)
	message_templateModule := message_template.InitMessageTemplateModule(database)
	website_configModule := website_config.InitWebsiteConfigModule(database)
	messageModule := message.InitMessageModule(emailModule, message_templateModule, website_configModule)
	post_likeModule := post_like.InitPostLikeModule(database)
	postModule := post.InitPostModule(database, website_configModule, post_likeModule, eventBus)
	commentModule := comment.InitCommentModule(database, messageModule, website_configModule, postModule, eventBus)
	commentHandler := commentModule.Hdl
	websiteConfigHandler := website_configModule.Hdl
	friendModule := friend.InitFriendModule(database, messageModule, website_configModule)
	friendHandler := friendModule.Hdl
	postHandler := postModule.Hdl
	visit_logModule := visit_log.InitVisitLogModule(database, eventBus)
	visitLogHandler := visit_logModule.Hdl
	messageTemplateHandler := message_templateModule.Hdl
	tagModule := tag.InitTagModule(database, eventBus)
	tagHandler := tagModule.Hdl
	count_statsModule := count_stats.InitCountStatsModule(database, eventBus)
	data_analysisModule := data_analysis.InitDataAnalysisModule(database, count_statsModule, post_likeModule, commentModule, visit_logModule)
	dataAnalysisHandler := data_analysisModule.Hdl
	countStatsHandler := count_statsModule.Hdl
	backupModule := backup.InitBackupModule(database)
	backupHandler := backupModule.Hdl
	writer := ioc.InitLogger()
	v, err := global.IsWebsiteInitializedFn(database)
	if err != nil {
		return nil, err
	}
	v2 := ioc.InitMiddlewares(writer, v)
	validators := ioc.InitGinValidators()
	post_indexModule := post_index.InitPostIndexModule(website_configModule, categoryModule, tagModule, postModule, module)
	postIndexHandler := post_indexModule.Hdl
	post_draftModule := post_draft.InitPostDraftModule(database)
	postDraftHandler := post_draftModule.Hdl
	aggregate_postModule := aggregate_post.InitAggregatePostModule(postModule, post_draftModule)
	aggregatePostHandler := aggregate_postModule.Hdl
	postLikeHandler := post_likeModule.Hdl
	post_visitModule := post_visit.InitPostVisitModule(database)
	postVisitHandler := post_visitModule.Hdl
	web_engineModule := web_engine.InitWebEngineModule(eventBus,website_configModule,categoryModule)
	webEngineHandler := web_engineModule.Hdl
	engine, err := ioc.NewGinEngine(fileHandler, categoryHandler, commentHandler, websiteConfigHandler, friendHandler, postHandler, visitLogHandler, messageTemplateHandler, tagHandler, dataAnalysisHandler, countStatsHandler, backupHandler, v2, validators, postIndexHandler, postDraftHandler, aggregatePostHandler, postLikeHandler, postVisitHandler, webEngineHandler)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
