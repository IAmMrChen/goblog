package controllers

import (
	"cyc/goblog/app/models/article"
	"cyc/goblog/app/policies"
	"cyc/goblog/app/requests"
	"cyc/goblog/pkg/auth"
	"cyc/goblog/pkg/flash"
	"cyc/goblog/pkg/logger"
	"cyc/goblog/pkg/route"
	"cyc/goblog/pkg/view"
	"fmt"
	"net/http"
)

type ArticlesController struct {
	BaseController
}

func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取URL参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	articles, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		// ---  4. 读取成功，显示文章 ---
		view.Render(w, view.D{
			"Article": articles,
			"CanModifyArticle": policies.CanModifyArticle(articles),
		}, "articles.show", "articles._article_meta")

	}
}

func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	// 1. 获取结果集
	articles, pagerData, err := article.GetAll(r, 10)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Articles": articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")

	}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request)  {
	view.RenderCreateOrEdit(w, view.D{}, "articles.create", "articles._form_field")
}

// Store 文章创建页面
func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	_article := article.Article{
		Title: title,
		Body:  body,
		UserID: auth.User().ID,
	}

	// 2. 表单验证
	errors := requests.ValidateArticleForm(_article)
	//fmt.Println(errors)
	// 检查是否有错误
	if len(errors) == 0 {
		_article.Create()

		if _article.ID > 0 {
			indexURL := route.Name2URL("articles.show", "id", _article.GetStringID())
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内容错误")
		}

	} else {
		view.Render(w, view.D{
			"Article": _article,
			"Errors": errors,
		}, "articles.create", "articles._form_field")
	}
}

func (ac *ArticlesController) Edit(w http.ResponseWriter, r *http.Request)  {
	// 1. 获取URL参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_article, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			// 4. 读取成功，显示表单
			view.RenderCreateOrEdit(w, view.D{
				"Article": _article,
				"Errors": view.D{},
			}, "articles.edit", "articles._form_field")
		}
	}
}

func (ac *ArticlesController) Update(w http.ResponseWriter, r *http.Request)  {
	// 1. 获取URL参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_article, err := article.Get(id)
	fmt.Println("_article is")
	fmt.Println(_article)
	// 3. 如果出现错误
	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		// 4. 未出现错误

		// 检查权限
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			// 4.1 表单验证
			_article.Title = r.PostFormValue("title")
			_article.Body = r.PostFormValue("body")
			fmt.Println("_article is")
			fmt.Println(_article)
			errors := requests.ValidateArticleForm(_article)

			if len(errors) == 0 {

				// 4.2 表单验证通过，更新数据
				rowsAffected, err := _article.Update()

				if err != nil {
					// 数据库错误
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "500 服务器内部错误")
					return
				}

				// √ 更新成功，跳转到文章详情页
				if rowsAffected > 0 {
					showURL := route.Name2URL("articles.show", "id", id)
					http.Redirect(w, r, showURL, http.StatusFound)
				} else {
					fmt.Fprint(w, "您没有做任何更改！")
				}
			} else {

				// 4.3 表单验证不通过，显示理由
				view.Render(w, view.D{
					"Article": _article,
					"Errors":  errors,
				}, "articles.edit", "articles._form_field")
			}
		}

	}
}

func (ac *ArticlesController) Delete(w http.ResponseWriter, r *http.Request)  {
	// 1.获取URL参数
	id := route.GetRouteVariable("id", r)

	// 2. 获取已经存在的文章
	_article, err := article.Get(id)

	// 3.如果出现错误
	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		// 检查权限
		if !policies.CanModifyArticle(_article) {
			flash.Warning("您没有权限执行此操作！")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			// 4. 未出现错误，执行删除操作
			rowAffected, err := _article.Delete()

			// 4.1 发生错误
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			} else {
				// 4.2 未发生错误
				if rowAffected > 0 {
					// 重定向到文章列表页面
					indexURL := route.Name2URL("articles.index")
					http.Redirect(w, r, indexURL, http.StatusFound)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w, "404 文章未找到")
				}
			}
		}
	}
}
