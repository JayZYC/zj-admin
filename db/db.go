package db

import (
	"fmt"
	"log"
	"os"
	"zj-admin/config"

	"gorm.io/driver/postgres" // 如果是mysql这个地方就是mysql的driver包
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// Init 创建数据库连接池
func Init() {
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	database := os.Getenv("db_database")
	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		username,
		database,
		password)

	var err error
	db, err = gorm.Open(postgres.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 关闭默认复数表名
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if config.Debug() {
		// 启用Logger，显示详细日志
		db = db.Debug()
	}

	if err := db.Exec(StoreProcedure).Error; err != nil {
		log.Fatal("存储过程执行失败:", err)
	}
}
