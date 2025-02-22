package service

import (
	"context"
	"go-chat/database"
	"go-chat/model"

	"github.com/google/uuid"
)

type MessageService struct{}

var MessageServiceApp = new(MessageService)

func (s *MessageService) PrivateSendText(ctx context.Context, content string, fromUserID, toUserID uint) (uuid.UUID, error) {
	textMessage := model.TextMessage{
		BaseMessage: model.BaseMessage{
			FromUserID:  fromUserID,
			ToUserID:    toUserID,
			ChatType:    1,
			MessageType: 1,
		},
		Content: content,
	}

	if err := database.PG.Create(&textMessage).Error; err != nil {
		return uuid.Nil, err
	}

	return textMessage.ID, nil
}

func (s *MessageService) PublicSendText(ctx context.Context, content string, fromUserID, toUserID uint) (uuid.UUID, error) {
	textMessage := model.TextMessage{
		BaseMessage: model.BaseMessage{
			FromUserID:  fromUserID,
			ToUserID:    toUserID,
			ChatType:    2,
			MessageType: 1,
		},
		Content: content,
	}

	if err := database.PG.Create(&textMessage).Error; err != nil {
		return uuid.Nil, err
	}

	return textMessage.ID, nil
}
