package controllers

import (
	"net/http"
)

type ArticlesController struct {

}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取URL参数
	//id := GetRouteVariable("id", r)
	//
	//// 2. 读取对应的文章数据
	//article, err := getArticleByID(id)
	//
	//// 3. 如果出现错误
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		// 3.1 数据未找到
	//		w.WriteHeader(http.StatusNotFound)
	//		fmt.Fprint(w, "404 文章未找到")
	//	} else {
	//		// 3.2 数据库错误
	//		logger.LogError(err)
	//		w.WriteHeader(http.StatusInternalServerError)
	//		fmt.Fprint(w, "500 服务器内部错误")
	//	}
	//} else {
	//
	//	tmpl, err := template.New("show.gohtml").
	//		Funcs(template.FuncMap{
	//			"RouteName2URL": route.Name2URL,
	//			"Int64ToString": types.Int64ToString,
	//		}).
	//		ParseFiles("resources/views/articles/show.gohtml")
	//
	//	logger.LogError(err)
	//
	//	tmpl.Execute(w, article)
	//}
}