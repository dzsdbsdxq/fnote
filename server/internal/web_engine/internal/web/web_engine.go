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
	"strconv"
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
	engine.GET("/post/:id", apiwrap.WrapHtml(w.PostDetail, "post"))
	engine.GET("/categories/:id", apiwrap.WrapHtml(w.GetCategories, "categories"))
}

func (w *WebEngineHandler) IndexHome(ctx *gin.Context) (*CommonParams[IndexHomeVO], error) {
	return &CommonParams[IndexHomeVO]{
		PageIdx: "index",
	}, nil
}
func (w *WebEngineHandler) PostDetail(ctx *gin.Context) (*CommonParams[PostDetail], error) {
	sug := ctx.Param("id")
	postDetail, err := w.serv.GetPostDetailById(ctx, sug, ctx.ClientIP())
	if err != nil {
		return nil, err
	}
	return &CommonParams[PostDetail]{
		PageIdx: "postDetail",
		Data:    PostDetail{Id: sug, Post: *postDetail},
	}, nil
}
func (w *WebEngineHandler) GetCategories(ctx *gin.Context) (*CommonParams[post.PageReq], error) {
	pageIdx := ctx.Param("id")
	pageNo, _ := strconv.ParseInt(ctx.Param("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.Param("pageSize"), 10, 64)
	if pageNo == 0 {
		pageNo = 1
	}
	if pageSize == 0 {
		pageSize = 30
	}
	return &CommonParams[post.PageReq]{
		PageIdx: pageIdx,
		PageRequest: post.PageReq{
			PageNo:   pageNo,
			PageSize: pageSize,
		},
	}, nil

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
		"func_posts": func(page int64, pageSize int64, field string, order string, category string, tags string) []*post.Post {
			req := &post.PostRequest{
				Categories: []string{category},
				Tags:       []string{tags},
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
