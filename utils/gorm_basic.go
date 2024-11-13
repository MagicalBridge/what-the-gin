package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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

//func GormDBConnection() {
//	// MySQL 数据库连接配置
//	dsn := "root:cpf@1234@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
//	// 创建数据库连接池
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Printf("Failed to connect to database: %v", err)
//	}
//
//	// 基于模型完成表结构的迁移和定义
//	err = db.AutoMigrate(&Article{})
//	if err != nil {
//		fmt.Printf("Failed to migrate table: %v", err)
//	}
//}

var DB *gorm.DB

func init() {
	// MySQL 数据库连接配置
	const dsn = "root:cpf@1234@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	// 创建数据库连接池
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}
	DB = db
}

// Create 插入数据的操作演示
func Create() {
	//	构建Article类型的数据
	article := &Article{
		Subject:     "GORM 基础的增删改查的操作",
		Likes:       0,
		Published:   true,
		PublishTime: time.Now(),
		AuthorID:    42,
	}

	//	DB.Create 完成数据库的 insert 操作
	if err := DB.Create(article).Error; err != nil {
		log.Fatal(err)
	}
}

// Retrieve 查询的数据代码演示
func Retrieve(id int) {
	//	初始化Article模型 零值
	article := &Article{}

	//	DB.first
	if err := DB.First(article, id).Error; err != nil {
		log.Fatal(err)
	}
	// 打印相关信息
	fmt.Println(article)
}

func Update() {
	//	获取需要更新的对象
	article := &Article{}

	if err := DB.First(article, 1).Error; err != nil {
		log.Fatal(err)
	}

	//  更新对象字段
	article.AuthorID = 23
	article.Likes = 15
	article.Subject = "GORM 的增删改查的练习"

	// 存储，DB.Save()
	if err := DB.Save(article).Error; err != nil {
		log.Fatal(err)
	}

	// 打印数据
	fmt.Println(article)
}

func Delete() {
	//	获取需要更新的对象
	article := &Article{}

	if err := DB.First(article, 1).Error; err != nil {
		log.Fatal(err)
	}

	// DB.Delete() 删除
	if err := DB.Delete(article).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println(article)
}
