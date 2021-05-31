package bootstrap

import (
	"cyc/goblog/app/models/article"
	"cyc/goblog/app/models/user"
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
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB)  {
	// 自动迁移
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}