package database

import (
	"fmt"
	"go-chat/model"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	PG *gorm.DB
)

func InitDB() error {
	var err error

	if err = godotenv.Load(); err != nil {
		return err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略 ErrRecordNotFound 错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	config := &gorm.Config{
		PrepareStmt: true,
		Logger:      newLogger,
	}

	PG, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return err
	}

	// 启用 uuid-ossp 扩展
	if err := PG.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return err
	}
	if err = PG.Migrator().DropTable(&model.TextMessage{}); err != nil {
		return err
	}

	if err = PG.AutoMigrate(model.TextMessage{}); err != nil {
		return err
	}
	return nil
}
