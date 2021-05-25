package article

import (
	"cyc/goblog/app/models"
	"cyc/goblog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	ID    int64
	Title string
	Body  string
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(a.ID, 10))
}