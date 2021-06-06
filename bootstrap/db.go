package bootstrap

import (
	"cyc/goblog/app/models/article"
	"cyc/goblog/app/models/category"
	"cyc/goblog/app/models/user"
	"cyc/goblog/pkg/config"
	"cyc/goblog/pkg/model"
	"gorm.io/gorm"
	"time"
)

func SetupDB()  {
	// 建立数据库链接池
	db := model.ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Minute)

	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB)  {
	// 自动迁移
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
	)
}