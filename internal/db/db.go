package db

import (
	"context"
	"time"
)

// Chat представляет чат с его атрибутами
type Chat struct {
	ID        int       // Уникальный идентификатор чата в базе данных, используется как первичный ключ
	ChatName  string    // Название чата, которое отображается участникам
	CreatedAt time.Time // Время создания чата
}

// ChatMember представляет участника чата
type ChatMember struct {
	ChatID    int       // Идентификатор чата, в котором состоит пользователь (ссылка на таблицу Chat)
	UserEmail string    // Email пользователя, присоединившегося к чату
	JoinedAt  time.Time // Время, когда пользователь присоединился к чату
}

// Message представляет сообщение в базе данных
type Message struct {
	ID        int       // Уникальный идентификатор сообщения
	ChatID    int       // Идентификатор чата, в который отправлено сообщение (ссылка на таблицу Chat)
	UserEmail string    // Email пользователя, отправившего сообщение
	Text      string    // Содержимое сообщения
	Timestamp time.Time // Время отправки сообщения
}

// DB - определяет контракт для работы с базой данных и определяет методы CRUD
type DB interface {

	//					Chats

	// CreateChat создает новый чат
	CreateChat(ctx context.Context, chatName string) (int, error)

	// ReadChat - Возвращает чат по его ID
	// ReadChat(ctx context.Context, id int) (*Chat, error)

	// ReadChats - Возвращает список чатов
	ReadChats(сtx context.Context) ([]*Chat, error)

	// UpdateChat

	// DeleteChat - удаляет чат по ID
	DeleteChat(ctx context.Context, id int) error

	// 					Members

	// CreateChatMember - добавляет участника в чат
	CreateChatMember(ctx context.Context, chatID int, userEmail string) error

	// ReadChatMembers - возвращает список участников чата
	ReadChatMembers(сtx context.Context, chatID int) ([]*ChatMember, error)

	// DeleteChatMember - удаляет участника из чата
	DeleteChatMember(сtx context.Context, chatID int, userEmail string) error

	// 					Message

	// CreateMessage - отправляет сообщение в чат
	CreateMessage(ctx context.Context, chatID int, userEmail, text string) (int, error)

	// ListMessages - возвращает все сообщения из чата
	ReadMessages(ctx context.Context, chatID int) ([]*Message, error)

	// DeleteMessage удаляет сообщение по ID
	DeleteMessage(ctx context.Context, id int) error
}
