package main

import (
	"cyc/goblog/app/http/middlewares"
	"cyc/goblog/bootstrap"
	"cyc/goblog/pkg/database"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

var router *mux.Router
var db *sql.DB

func main()  {
	router = bootstrap.SetupRoute()
	bootstrap.SetupDB()

	database.Initialize()
	db = database.DB


	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}