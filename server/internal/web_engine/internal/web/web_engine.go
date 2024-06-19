package web

import (
	apiwrap "github.com/chenmingyong0423/fnote/server/internal/pkg/web/wrap"
	"github.com/chenmingyong0423/fnote/server/internal/web_engine/internal/service"
	"github.com/gin-gonic/gin"
)

func NewWebEngineHandler(serv service.IWebEngineService) *WebEngineHandler {
	return &WebEngineHandler{serv: serv}
}

type WebEngineHandler struct {
	serv service.IWebEngineService
}

func (w *WebEngineHandler) RegisterGinRoutes(engine *gin.Engine) {
	engine.Static("/themes", "./templates/default")
	//websiteRouter.InitHomeFuncMap(Router)
	engine.LoadHTMLGlob("./templates/default/*.tpl")

	engine.GET("/", apiwrap.WrapHtml(w.IndexHome, "index"))
	engine.GET("/post/:id", apiwrap.WrapHtml(w.IndexHome, "index"))
}

func (w *WebEngineHandler) IndexHome(ctx *gin.Context) (*apiwrap.ResponseBody[IndexHomeVO], error) {
	config, err := w.serv.GetWebSiteConfig(ctx)
	if err != nil {
		return nil, err
	}
	menus, err := w.serv.GetCategory(ctx)
	return apiwrap.SuccessResponseWithData(IndexHomeVO{
		BlogName: "博客首页",
		Meta:     config,
		Menus:    menus,
	}), nil
}
