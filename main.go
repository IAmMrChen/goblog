package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)


var router = mux.NewRouter()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	fmt.Println(vars)
	id := vars["id"]
	fmt.Fprint(w, "文章ID：" + id)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {


	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "请输入内容")
		return
	}

	title := r.PostForm.Get("title")

	fmt.Fprintf(w, "r.Form 中 title 的值为: %v \n", r.FormValue("title"))
	fmt.Fprintf(w, "POST PostForm: %v \n", r.PostForm)
	fmt.Fprintf(w, "POST Form: %v \n", r.Form)
	fmt.Fprintf(w, "title的值为: %v ", title)

	
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

func articlesCreateHandle(w http.ResponseWriter, r *http.Request)  {
	html := `
    <!DOCTYPE html>
<html lang="en">
<head>
    <title>创建文章 —— 我的技术博客</title>
</head>
<body>
    <form action="%s" method="post">
        <p><input type="text" name="title"></p>
        <p><textarea name="body" cols="30" rows="10"></textarea></p>
        <p><button type="submit">提交</button></p>
    </form>
</body>
</html>
    `

	storeURL, _ := router.Get("articles.store").URL()
	fmt.Fprintf(w, html, storeURL)
}


func main()  {


	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")

	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")

	router.HandleFunc("/articles/create", articlesCreateHandle).Methods("GET").Name("articles.create")

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	//router.Use(forceHTMLMiddleware)

	// 通过命名路由获取 URL 示例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}