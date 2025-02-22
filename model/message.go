package model

import "github.com/google/uuid"

// 基本消息模型
// ChatType 1 表示私聊, 2 表示群聊
// MessageType 1 表示文本消息, 2 表示图片消息, 3 表示语音消息
type BaseMessage struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	MessageType uint8     `gorm:"type:smallint;not null;default:1" json:"message_type"`
	ChatType    uint8     `gorm:"type:smallint;not null;default:1" json:"chat_type"`
	FromUserID  uint      `gorm:"index;not null" json:"from_user_id"`
	ToUserID    uint      `gorm:"index;not null" json:"to_user_id"`
	IsRevoked   bool      `gorm:"not null;default:false" json:"is_revoked"`
}

// 文本消息模型
type TextMessage struct {
	BaseMessage
	Content string `gorm:"type:text;not null" json:"content"`
}
