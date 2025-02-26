package service

import (
	"context"
	"crypto/rand"
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
	// 启用 uuid-ossp 扩展
	if err := database.Migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	m.Run()
}

func TestSendText(t *testing.T) {
	ctx := context.Background()
	fromID := uuid.New()
	toID := uuid.New()
	id, err := MessageServiceApp.PrivateSendText(ctx, rand.Text(), fromID, toID)
	if err != nil {
		t.Fatalf("PrivateSendText failed: %v", err)
	}

	// 检查返回的 UUID 是否有效
	if id == uuid.Nil {
		t.Error("Expected valid UUID, got Nil")
	}

	// 使用 %s 来正确格式化 UUID
	t.Logf("PrivateSendText success: %s", id.String())

	id, err = MessageServiceApp.PublicSendText(ctx, rand.Text(), fromID, toID)
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
	fromID := uuid.New()
	toID := uuid.New()

	b.Run("PrivateSendText", func(b *testing.B) {
		for b.Loop() {
			_, err := MessageServiceApp.PrivateSendText(ctx, rand.Text(), fromID, toID)
			if err != nil {
				b.Fatalf("PrivateSendText failed: %v", err)
			}
		}
	})

	b.Run("PublicSendText", func(b *testing.B) {
		for b.Loop() {
			_, err := MessageServiceApp.PublicSendText(ctx, rand.Text(), fromID, toID)
			if err != nil {
				b.Fatalf("PublicSendText failed: %v", err)
			}
		}
	})
}

func TestGetMessages(t *testing.T) {
	fromID, toID := uuid.New(), uuid.New()
	text := rand.Text()
	_, _ = MessageServiceApp.PrivateSendText(context.Background(), text, fromID, toID)
	_, _ = MessageServiceApp.PublicSendText(context.Background(), text, fromID, toID)
	pravateMessages, err := MessageServiceApp.GetPrivateMessages(context.Background(), fromID, toID)
	if err != nil {
		t.Fatalf("GetPrivateMessages failed: %v", err)
	}
	if len(pravateMessages) == 0 {
		t.Error("Expected messages, got 0")
	}
	if pravateMessages[0].Content != text {
		t.Errorf("Expected %s, got %s", text, pravateMessages[0].Content)
	}
	publicMessages, err := MessageServiceApp.GetPublicMessages(context.Background(), toID)
	if err != nil {
		t.Fatalf("GetPublicMessages failed: %v", err)
	}
	if len(publicMessages) == 0 {
		t.Error("Expected messages, got 0")
	}
	if publicMessages[0].Content != text {
		t.Errorf("Expected %s, got %s", text, publicMessages[0].Content)
	}
}
