// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package message_template

import (
	"github.com/chenmingyong0423/fnote/server/internal/message_template/internal/repository"
	"github.com/chenmingyong0423/fnote/server/internal/message_template/internal/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/message_template/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/message_template/internal/web"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitMessageTemplateModule(mongoDB *mongo.Database) *Module {
	messageTemplateDao := dao.NewMessageTemplateDao(mongoDB)
	messageTemplateRepository := repository.NewMessageTemplateRepository(messageTemplateDao)
	messageTemplateService := service.NewMessageTemplateService(messageTemplateRepository)
	messageTemplateHandler := web.NewMessageTemplateHandler(messageTemplateService)
	module := &Module{
		Svc: messageTemplateService,
		Hdl: messageTemplateHandler,
	}
	return module
}

// wire.go:

var MessageTemplateProviders = wire.NewSet(web.NewMessageTemplateHandler, service.NewMessageTemplateService, repository.NewMessageTemplateRepository, dao.NewMessageTemplateDao, wire.Bind(new(service.IMessageTemplateService), new(*service.MessageTemplateService)), wire.Bind(new(repository.IMessageTemplateRepository), new(*repository.MessageTemplateRepository)), wire.Bind(new(dao.IMessageTemplateDao), new(*dao.MessageTemplateDao)))
