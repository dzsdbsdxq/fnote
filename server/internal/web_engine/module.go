package web_engine

import (
	"github.com/chenmingyong0423/fnote/server/internal/web_engine/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/web_engine/internal/web"
)

type (
	Handler = web.WebEngineHandler
	Service = service.IWebEngineService
	Module  struct {
		Svc Service
		Hdl *Handler
	}
)
