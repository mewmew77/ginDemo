package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// GetMysqlDB todo 单例模型
func GetMysqlDB() *gorm.DB {
	dsn := "user:password@tcp(127.0.0.1:3306)/test_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database init fail, error = %v", err)
	}
	return db
}
