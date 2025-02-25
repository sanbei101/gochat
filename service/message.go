package service

import (
	"context"
	"go-chat/database"
	"go-chat/model"

	"github.com/google/uuid"
)

type MessageService struct{}

var MessageServiceApp = new(MessageService)

func (s *MessageService) PrivateSendText(ctx context.Context, content string, FromID, toUserID uuid.UUID) (uuid.UUID, error) {
	textMessage := model.TextMessage{
		BaseMessage: model.BaseMessage{
			FromID:      FromID,
			ToID:        toUserID,
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

func (s *MessageService) PublicSendText(ctx context.Context, content string, FromID, toUserID uuid.UUID) (uuid.UUID, error) {
	textMessage := model.TextMessage{
		BaseMessage: model.BaseMessage{
			FromID:      FromID,
			ToID:        toUserID,
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

func (s *MessageService) GetPrivateMessages(ctx context.Context, fromID, toID uuid.UUID) ([]model.TextMessage, error) {
	var messages []model.TextMessage
	err := database.PG.Where("from_id = ? AND to_id = ? AND chat_type = 1", fromID, toID).Find(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *MessageService) GetPublicMessages(ctx context.Context, groupID uuid.UUID) ([]model.TextMessage, error) {
	var messages []model.TextMessage
	err := database.PG.Where("to_id = ? AND chat_type = 2", groupID).Find(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}
