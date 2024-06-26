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

package domain

type PostDraft struct {
	Id               string               `json:"_id"`
	Author           string               `json:"author"`
	Title            string               `json:"title"`
	Summary          string               `json:"summary"`
	CoverImg         string               `json:"cover_img"`
	Categories       []Category4PostDraft `json:"category"`
	Tags             []Tag4PostDraft      `json:"tags"`
	StickyWeight     int                  `json:"sticky_weight"`
	Content          string               `json:"content"`
	MetaDescription  string               `json:"meta_description"`
	MetaKeywords     string               `json:"meta_keywords"`
	WordCount        int                  `json:"word_count"`
	IsDisplayed      bool                 `json:"is_displayed"`
	IsCommentAllowed bool                 `json:"is_comment_allowed"`
	CreatedAt        int64                `json:"created_at"`
}

type Category4PostDraft struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Tag4PostDraft struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Page struct {
	// 当前页
	PageNo int64
	// 每页数量
	PageSize int64
	// 排序字段
	Field string
	// 排序规则
	Order string
	// 搜索内容
	Keyword string
}

func (p *Page) OrderConvertToInt() int {
	switch p.Order {
	case "ASC":
		return 1
	case "DESC":
		return -1
	default:
		return -1
	}
}

type PageQuery struct {
	Size    int64
	Skip    int64
	Keyword string
	Field   string
	Order   int
}
