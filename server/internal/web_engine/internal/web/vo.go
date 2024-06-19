package web

import (
	"github.com/chenmingyong0423/fnote/server/internal/category"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
)

type MetaInfo struct {
	Author   string `json:"author"`
	KeyWords string `json:"key_words"`
	Robots   string `json:"robots"`
}
type IndexHomeVO struct {
	BlogName string                        `json:"blog_name"`
	BlogDesc string                        `json:"blog_desc"`
	Meta     *website_config.WebsiteConfig `json:"meta"`
	Menus    []category.Category           `json:"menus"`
}
