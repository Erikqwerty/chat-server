package repository

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

// ChatServerRepository - определяет контракт для работы с базой данных
type ChatServerRepository interface {
	Chat
	ChatMember
	Message
}

// Chat определяет контракт для работы с таблицой chats
type Chat interface {
	// CreateChat создает новый чат
	CreateChat(ctx context.Context, chatName string) (int, error)
	// ReadChat - Возвращает чат по его ID
	ReadChat(ctx context.Context, id int) (*model.Chat, error)
	// ReadChats - Возвращает список чатов
	ReadChats(сtx context.Context) ([]*model.Chat, error)
	// DeleteChat - удаляет чат по ID
	DeleteChat(ctx context.Context, id int) error
}

// ChatMember определяет контракт для работы с таблицой chat_members
type ChatMember interface {
	// CreateChatMember - добавляет участника в чат
	CreateChatMember(ctx context.Context, chatID int, userEmail string) error
	// ReadChatMember - возвращает  участника чата если он есть
	ReadChatMember(сtx context.Context, userEmail string, chatID int) (*model.ChatMember, error)
	// ReadChatMembers - возвращает список участников чата
	ReadChatMembers(сtx context.Context, chatID int) ([]*model.ChatMember, error)
	// DeleteChatMember - удаляет участника из чата
	DeleteChatMember(сtx context.Context, chatID int, userEmail string) error
}

// Message определяет контракт для работы с таблицой messages
type Message interface {
	// CreateMessage - отправляет сообщение в чат
	CreateMessage(ctx context.Context, chatID int, userEmail, text string) (int, error)
	// ReadMessages - возвращает все сообщения из чата
	ReadMessages(ctx context.Context, chatID int) ([]*model.Message, error)
	// DeleteMessage удаляет сообщение по ID
	DeleteMessage(ctx context.Context, id int) error
}
