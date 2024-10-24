package service

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

// ChatService - интерфейс определяющий логику chat-server
type ChatService interface {
	// CreateChat - Логика создания чата
	CreateChat(context.Context, *model.CreateChat) (int64, error)
	// JoinChat - Логика добавление нового пользователя в чат
	JoinChat(context.Context, *model.ChatMember) (*model.JoinChat, error)
	// DeleteChat - Логика удаления чата
	DeleteChat(context.Context, int64) error
	// SendMessage - Логика отправки сообщения
	SendMessage(context.Context, *model.Message) error
}
