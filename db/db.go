package db

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
	"zj-admin/config"
	"zj-admin/model"
)
import "github.com/jinzhu/gorm"

var db *gorm.DB

// Init 创建数据库连接池
func Init() *gorm.DB {
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
	db, err = gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}

	//set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected
	db.SingularTable(true)

	if config.Debug() {
		db = db.Debug()
	}

	// 调用注册函数
	// 在使用gorm创建新的数据记录时，自动加上id和时间
	db.Callback().Create().Before("gorm:create").Register("before_created", beforeCreated)
	db.Callback().Update().Before("gorm:update").Register("before_updated", beforeUpdated)

	/*建表或更新表*/
	db.AutoMigrate(
		&model.Organization{},
		&model.Role{},
		&model.User{},
	)

	if err := db.Exec(StoreProcedure).Error; err != nil {
		log.Fatal("存储过程执行失败:", err)
	}

	return db
}

func beforeCreated(scope *gorm.Scope) {
	now := time.Now().UnixNano() / 1e6
	scope.SetColumn("id", uuid.New())
	scope.SetColumn("create_time", now)
	scope.SetColumn("update_time", now)
}

func beforeUpdated(scope *gorm.Scope) {
	scope.SetColumn("update_time", time.Now().UnixNano()/1e6)
}
