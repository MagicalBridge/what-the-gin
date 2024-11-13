package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	//	嵌入基础模型
	gorm.Model
	//	定义字段
	Subject     string
	Likes       uint
	Published   bool
	PublishTime time.Time
	AuthorID    uint
}

func GormDBConnection() string {
	// MySQL 数据库连接配置
	dsn := "root:cpf@1234@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	// 创建数据库连接池
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}

	// 基于模型完成表结构的迁移和定义
	err = db.AutoMigrate(&Article{})
	if err != nil {
		fmt.Printf("Failed to migrate table: %v", err)
	}
	return dsn
}
