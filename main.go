package main

import (
	"cyc/goblog/app/http/middlewares"
	"cyc/goblog/bootstrap"
	"cyc/goblog/config"
	config2 "cyc/goblog/pkg/config"
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

var router *mux.Router
var db *sql.DB

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main()  {
	router = bootstrap.SetupRoute()
	bootstrap.SetupDB()

	//database.Initialize()
	//db = database.DB


	http.ListenAndServe(":" + config2.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}