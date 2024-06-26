// Copyright 2024 chenmingyong0423

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
	"strings"

	"github.com/chenmingyong0423/fnote/server/internal/pkg/jwtutil"
	"github.com/gin-gonic/gin"
)

// JwtParseMiddleware jwt 解析中间件
func JwtParseMiddleware(isWebsiteInitialized func() bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		if !isWebsiteInitialized() && (uri == "/admin-api/files/upload" || uri == "/admin-api/configs/initialization") {
			ctx.Next()
			return
		}
		// 非 admin 接口不需要 jwt
		if !strings.HasPrefix(uri, "/admin-api") {
			ctx.Next()
			return
		}
		// 这些接口不需要 jwt
		if uri == "/admin-api/login" || uri == "/admin-api/configs/check-initialization" || uri == "/admin-api/configs/website/meta" {
			ctx.Next()
			return
		}

		jwtStr := ctx.GetHeader("Authorization")
		if jwtStr == "" {
			ctx.AbortWithStatusJSON(401, nil)
			return
		}
		// 解析 jwt
		claims, err := jwtutil.ParseJwt(jwtStr)
		if err != nil {
			ctx.AbortWithStatusJSON(401, nil)
			return
		}
		ctx.Set("jwtClaims", claims)
		ctx.Next()
	}
}
