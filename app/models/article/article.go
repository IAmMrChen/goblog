package article

import (
	"cyc/goblog/app/models"
	"cyc/goblog/app/models/category"
	"cyc/goblog/app/models/user"
	"cyc/goblog/pkg/route"
	"html/template"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	Title string
	Body  string
	HtmlBody template.HTML
	UserID uint64 `gorm: "not null;index"`
	User user.User
	CategoryID uint64 `gorm:"not null;default:4;index"`
	Category category.Category
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(article.ID), 10))
}

// CreatedAtDate 创建日期
func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}