package service

import (
	"context"
	"go-chat/database"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestMain(m *testing.M) {
	const dsn string = "host=localhost user=testuser password=justfortest dbname=testdatabase sslmode=disable"
	TestDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel:      logger.Warn,     // 设置日志级别
				SlowThreshold: time.Second * 2, // 慢查询的时间阈值
				Colorful:      false,           // 是否开启日志输出的彩色
			},
		),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	database.PG = TestDB
	m.Run()
}

func TestSendText(t *testing.T) {
	ctx := context.Background()

	id, err := MessageServiceApp.PrivateSendText(ctx, "Hello", 1, 2)
	if err != nil {
		t.Fatalf("PrivateSendText failed: %v", err)
	}

	// 检查返回的 UUID 是否有效
	if id == uuid.Nil {
		t.Error("Expected valid UUID, got Nil")
	}

	// 使用 %s 来正确格式化 UUID
	t.Logf("PrivateSendText success: %s", id.String())

	id, err = MessageServiceApp.PublicSendText(ctx, "Hello", 1, 2)
	if err != nil {
		t.Fatalf("PublicSendText failed: %v", err)
	}

	if id == uuid.Nil {
		t.Error("Expected valid UUID, got Nil")
	}
	t.Logf("PublicSendText success: %s", id.String())
}

func BenchmarkSendText(b *testing.B) {
	ctx := context.Background()

	b.Run("PrivateSendText", func(b *testing.B) {
		for b.Loop() {
			_, err := MessageServiceApp.PrivateSendText(ctx, "Hello", 1, 2)
			if err != nil {
				b.Fatalf("PrivateSendText failed: %v", err)
			}
		}
	})

	b.Run("PublicSendText", func(b *testing.B) {
		for b.Loop() {
			_, err := MessageServiceApp.PublicSendText(ctx, "Hello", 1, 2)
			if err != nil {
				b.Fatalf("PublicSendText failed: %v", err)
			}
		}
	})
}
