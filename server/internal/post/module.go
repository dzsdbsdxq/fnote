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

package post

import (
	"github.com/chenmingyong0423/fnote/server/internal/post/internal/domain"
	"github.com/chenmingyong0423/fnote/server/internal/post/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/post/internal/web"
)

type (
	Handler       = web.PostHandler
	Service       = service.IPostService
	Post          = domain.Post
	PageReq       = web.PageRequest
	PostRequest   = domain.PostRequest
	PageRequest   = domain.PageRequest
	Page2         = domain.Page2
	Sorting       = domain.Sorting
	PrimaryPost   = domain.PrimaryPost
	ExtraPost     = domain.ExtraPost
	Category4Post = domain.Category4Post
	Tag4Post      = domain.Tag4Post
	DetailPost    = domain.DetailPostVO
	Module        struct {
		Svc Service
		Hdl *Handler
	}
)
