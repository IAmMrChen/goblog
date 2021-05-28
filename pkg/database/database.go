package database

import (
	"cyc/goblog/pkg/logger"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

var DB *sql.DB

// Initialize 初始化数据库
func Initialize() {
	initDB()
	createTables()
}
// 初始化数据库链接
func initDB()  {
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "123456",
		Addr:                 "127.0.0.1:3305", // 公司
		//Addr:                 "127.0.0.1:3306",   // 家里
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	DB.SetConnMaxIdleTime(25)
	// 设置每个链接的过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)

	// 尝试链接，失败报错
	err = DB.Ping()
	logger.LogError(err)
}
// 创建表
func createTables()  {
	createArticlesSql := `CREATE TABLE IF NOT EXISTS articles(
    	id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    	title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    	body longtext COLLATE utf8mb4_unicode_ci
	);`

	_, err := DB.Exec(createArticlesSql)

	logger.LogError(err)
}