package controllers

import (
	"cyc/goblog/app/models/article"
	"cyc/goblog/pkg/logger"
	"cyc/goblog/pkg/route"
	"cyc/goblog/pkg/types"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

type ArticlesController struct {

}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// todo 这边这个方法后面看看怎么激活
	// 1. 获取URL参数
	//id := getRouteVariable("id", r)
	vars := mux.Vars(r)
	id := vars["id"]

	// 2. 读取对应的文章数据
	articles, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {

		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2URL": route.Name2URL,
				"Int64ToString": types.Int64ToString,
			}).
			ParseFiles("resources/views/articles/show.gohtml")

		logger.LogError(err)

		tmpl.Execute(w, articles)
	}
}

func (*ArticlesController)Index(w http.ResponseWriter, r *http.Request) {
	// 获取结果集
	articles, err := article.GetAll()

	if err != nil {
		// 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		// 2.加载模版
		tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
		logger.LogError(err)

		// 3.渲染模版，将所有文章的数据传输出去
		tmpl.Execute(w, articles)
	}
}