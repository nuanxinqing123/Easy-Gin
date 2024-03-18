package initialize

import (
	"os"

	"gorm.io/gorm"

	"Easy-Gin/internal/model"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	// 数据表：自动迁移
	err := db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		&model.User{},
	)
	if err != nil {
		os.Exit(0)
	}
}
