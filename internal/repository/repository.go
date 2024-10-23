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
	CreateChat(context.Context, string) (int, error)
	// ReadChat - Возвращает чат по его ID
	ReadChat(context.Context, int) (*model.Chat, error)
	// ReadChats - Возвращает список чатов
	ReadChats(context.Context) ([]*model.Chat, error)
	// DeleteChat - удаляет чат по ID
	DeleteChat(context.Context, int) error
}

// ChatMember определяет контракт для работы с таблицой chat_members
type ChatMember interface {
	// CreateChatMember - добавляет участника в чат
	CreateChatMember(context.Context, *model.ChatMember) error
	// ReadChatMember - возвращает  участника чата если он есть
	ReadChatMember(context.Context, *model.ChatMember) (*model.ChatMember, error)
	// ReadChatMembers - возвращает список участников чата
	ReadChatMembers(context.Context, int) ([]*model.ChatMember, error)
	// DeleteChatMember - удаляет участника из чата
	DeleteChatMember(context.Context, *model.ChatMember) error
}

// Message определяет контракт для работы с таблицой messages
type Message interface {
	// CreateMessage - отправляет сообщение в чат
	CreateMessage(context.Context, *model.Message) (int, error)
	// ReadMessages - возвращает все сообщения из чата
	ReadMessages(context.Context, int) ([]*model.Message, error)
	// DeleteMessage удаляет сообщение по ID
	DeleteMessage(context.Context, int) error
}
