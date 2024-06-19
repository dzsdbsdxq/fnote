// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ioc

import (
	"github.com/chenmingyong0423/fnote/server/internal/web_engine"
	"io"
	"log/slog"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/chenmingyong0423/fnote/server/internal/backup"

	"github.com/chenmingyong0423/fnote/server/internal/visit_log"

	"github.com/chenmingyong0423/fnote/server/internal/message_template"

	"github.com/chenmingyong0423/fnote/server/internal/category"
	"github.com/chenmingyong0423/fnote/server/internal/tag"

	"github.com/chenmingyong0423/fnote/server/internal/file"

	"github.com/chenmingyong0423/fnote/server/internal/count_stats"

	"github.com/chenmingyong0423/fnote/server/internal/friend"

	"github.com/chenmingyong0423/fnote/server/internal/post_visit"

	"github.com/chenmingyong0423/fnote/server/internal/comment"

	"github.com/chenmingyong0423/fnote/server/internal/data_analysis"

	"github.com/chenmingyong0423/fnote/server/internal/post"
	"github.com/chenmingyong0423/fnote/server/internal/post_like"

	"github.com/chenmingyong0423/fnote/server/internal/aggregate_post"

	"github.com/chenmingyong0423/fnote/server/internal/post_draft"
	"github.com/chenmingyong0423/fnote/server/internal/post_index"

	"github.com/chenmingyong0423/fnote/server/internal/website_config"

	"github.com/spf13/viper"

	"github.com/chenmingyong0423/ginx/middlewares/id"
	"github.com/chenmingyong0423/ginx/middlewares/log"
	"github.com/gin-contrib/cors"

	myValidator "github.com/chenmingyong0423/fnote/server/internal/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func NewGinEngine(fileHdr *file.Handler, ctgHdr *category.Handler, cmtHdr *comment.Handler, cfgHdr *website_config.Handler, frdHdr *friend.Handler, postHdr *post.Handler, vlHdr *visit_log.Handler, msgTplHandler *message_template.Handler, tagsHandler *tag.Handler, daHandler *data_analysis.Handler, csHandler *count_stats.Handler, backupHandler *backup.Handler, middleware []gin.HandlerFunc, validators Validators, postIndexHdr *post_index.Handler, postDraftHdr *post_draft.Handler, aggregatePostHdr *aggregate_post.Handler, postLikesHdr *post_like.Handler, postVisitHdr *post_visit.Handler, webEngineHdr *web_engine.Handler) (*gin.Engine, error) {
	engine := gin.New()
	engine.Use(gin.Recovery())

	// 参数校验器注册
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for k, v := range validators {
			err := validate.RegisterValidation(k, v)
			if err != nil {
				return nil, err
			}
		}
	}

	// 中间件注册
	engine.Use(middleware...)

	// 注册路由
	{
		webEngineHdr.RegisterGinRoutes(engine)

		ctgHdr.RegisterGinRoutes(engine)
		cmtHdr.RegisterGinRoutes(engine)
		cfgHdr.RegisterGinRoutes(engine)
		frdHdr.RegisterGinRoutes(engine)
		postHdr.RegisterGinRoutes(engine)
		vlHdr.RegisterGinRoutes(engine)
		msgTplHandler.RegisterGinRoutes(engine)
		tagsHandler.RegisterGinRoutes(engine)
		fileHdr.RegisterGinRoutes(engine)
		daHandler.RegisterGinRoutes(engine)
		csHandler.RegisterGinRoutes(engine)
		backupHandler.RegisterGinRoutes(engine)
		postIndexHdr.RegisterGinRoutes(engine)
		postDraftHdr.RegisterGinRoutes(engine)
		aggregatePostHdr.RegisterGinRoutes(engine)
		postLikesHdr.RegisterGinRoutes(engine)
		postVisitHdr.RegisterGinRoutes(engine)
	}
	return engine, nil
}

func InitMiddlewares(writer io.Writer, isWebsiteInitialized func() bool) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.LoggerWithWriter(writer),
		id.RequestId(),
		log.RequestLogger(*log.NewLoggerConfig(func(level string) slog.Level {
			switch level {
			case "DEBUG":
				return slog.LevelDebug
			case "INFO":
				return slog.LevelInfo
			case "WARN":
				return slog.LevelWarn
			case "ERROR":
				return slog.LevelError
			default:
				return slog.LevelInfo
			}
		}(viper.GetString("logger.level")), log.WithSkipPaths([]string{"/admin-api/files/upload", "/admin-api/recovery", "/admin-api/backup"}), log.WithSkipFunc(func(ctx *gin.Context) bool {
			url := ctx.Request.URL.Path
			return strings.HasPrefix(url, "/static/")
		}))),
		cors.New(cors.Config{
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				if slices.Contains(viper.GetStringSlice("gin.allowed_origins"), "*") {
					return true
				}
				return slices.ContainsFunc(viper.GetStringSlice("gin.allowed_origins"), func(s string) bool {
					return strings.Contains(origin, s)
				})
			},
			AllowMethods:  viper.GetStringSlice("gin.allowed_methods"),
			AllowHeaders:  viper.GetStringSlice("gin.allowed_headers"),
			ExposeHeaders: viper.GetStringSlice("gin.exposed_headers"),
			MaxAge:        12 * time.Hour,
		}),
		func(ctx *gin.Context) {
			uri := ctx.Request.RequestURI
			if (isWebsiteInitialized() || strings.HasPrefix(uri, "/static")) || uri == "/admin-api/files/upload" || uri == "/admin-api/configs/initialization" {
				ctx.Next()
			} else {
				ctx.JSON(http.StatusServiceUnavailable, nil)
				ctx.Abort()
			}
		},
		JwtParseMiddleware(isWebsiteInitialized),
	}
}

type Validators map[string]func(fl validator.FieldLevel) bool

func InitGinValidators() Validators {
	return map[string]func(fl validator.FieldLevel) bool{
		"validateEmailFormat": myValidator.ValidateEmailFormat,
	}
}
