package model

import (
	"github.com/google/uuid"
)

// User 用户模型
type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Username string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"username"`
	Status   uint8     `gorm:"type:smallint;not null;default:1" json:"status"`
}
