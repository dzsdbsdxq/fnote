package web

import "github.com/chenmingyong0423/fnote/server/internal/post"

type MetaInfo struct {
	Author   string `json:"author"`
	KeyWords string `json:"key_words"`
	Robots   string `json:"robots"`
}
type IndexHomeVO struct{}
type PostDetail struct {
	Id   string          `json:"id"`
	Post post.DetailPost `json:"post"`
}
