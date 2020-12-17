package bootstrap

import (
	"cms/app/models/article"
	"cms/app/models/clerk"
	"cms/pkg/config"
	"cms/pkg/model"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	//连接数、空闲连接数、过期时间
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	sqlDB.SetConnMaxIdleTime(time.Duration(config.GetInt("database.mysql.max_life_seconds")))

	migrate(db)
}

func migrate(db *gorm.DB)  {
	fmt.Println("Start AutoMigrate")
	db.AutoMigrate(
		&article.Article{},
		&clerk.Clerk{},
	)
}
