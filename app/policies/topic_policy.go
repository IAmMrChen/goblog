package policies

import (
	"cyc/goblog/app/models/article"
	"cyc/goblog/pkg/auth"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}