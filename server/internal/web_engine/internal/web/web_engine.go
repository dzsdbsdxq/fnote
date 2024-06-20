package web

import (
	"fmt"
	"github.com/chenmingyong0423/fnote/server/internal/category"
	apiwrap "github.com/chenmingyong0423/fnote/server/internal/pkg/web/wrap"
	"github.com/chenmingyong0423/fnote/server/internal/post"
	"github.com/chenmingyong0423/fnote/server/internal/web_engine/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func NewWebEngineHandler(serv service.IWebEngineService) *WebEngineHandler {
	return &WebEngineHandler{serv: serv}
}

type WebEngineHandler struct {
	serv service.IWebEngineService
}

func (w *WebEngineHandler) RegisterGinRoutes(engine *gin.Engine) {
	engine.Static("/themes", "./templates/default")
	w.InitHomeFuncMap(engine)
	engine.LoadHTMLGlob("./templates/default/*.tpl")
	engine.GET("/", apiwrap.WrapHtml(w.IndexHome, "index"))
	engine.GET("/post/:id", apiwrap.WrapHtml(w.IndexHome, "index"))
}

func (w *WebEngineHandler) IndexHome(ctx *gin.Context) (*apiwrap.ResponseBody[IndexHomeVO], error) {
	return apiwrap.SuccessResponseWithData(IndexHomeVO{}), nil
}

func (w *WebEngineHandler) InitHomeFuncMap(engine *gin.Engine) {
	engine.SetFuncMap(template.FuncMap{
		"func_web_site_config": func() *website_config.IndexConfig {
			config, _ := w.serv.GetWebSiteConfig(&gin.Context{})
			return config
		},
		"func_get_menus": func() []category.Category {
			menus, _ := w.serv.GetCategory(&gin.Context{})
			return menus
		},
		"func_latest_posts": func(count int64) []*post.Post {
			latestPost, _ := w.serv.GetLatestPosts(&gin.Context{}, count)
			return latestPost
		},
		"func_posts": func(page int64, pageSize int64, field string, order string, category []string, tags []string) []*post.Post {
			req := &post.PostRequest{
				Categories: category,
				Tags:       tags,
				PageRequest: post.PageRequest{
					Page2: post.Page2{
						PageNo:   page,
						PageSize: pageSize,
					},
					Sorting: post.Sorting{
						Field: &field,
						Order: &order,
					},
				},
			}
			list, _, err := w.serv.GetPostList(&gin.Context{}, req)
			fmt.Println(err)
			return list
		},
		"func_format_timestamp": func(timestamp int64, format string) string {
			t := time.Unix(timestamp, 0)
			// 格式化时间
			if format == "date" {
				return t.Format("2006-01-02")
			}
			return t.Format("2006-01-02 15:04:05")
		},
	})
}
