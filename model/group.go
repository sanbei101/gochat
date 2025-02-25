package model

import "github.com/google/uuid"

type Group struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"name"`
}
