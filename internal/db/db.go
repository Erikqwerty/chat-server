package db

import "time"

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

type Message struct {
	ID        int       // Уникальный идентификатор сообщения
	ChatID    int       // Идентификатор чата, в который отправлено сообщение (ссылка на таблицу Chat)
	UserEmail string    // Email пользователя, отправившего сообщение
	Text      string    // Содержимое сообщения
	Timestamp time.Time // Время отправки сообщения
}

type DB interface {
	// CreateChat создает новый чат
	CreateChat(chatName string) (int, error)

	// GetChat(id int) (*Chat, error)
	// ListChats() ([]*Chat, error)

	// DeleteChat удаляет чат по ID
	DeleteChat(id int) error

	// AddChatMember добавляет участника в чат
	AddChatMember(chatID int, userEmail string) error

	// ListChatMembers возвращает список участников чата
	// ListChatMembers(chatID int) ([]*ChatMember, error)

	// RemoveChatMember удаляет участника из чата
	// RemoveChatMember(chatID int, userEmail string) error

	// SendMessage отправляет сообщение в чат
	SendMessage(chatID int, userEmail, text string) (int, error)

	// ListMessages возвращает все сообщения из чата
	// ListMessages(chatID int) ([]*Message, error)

	// DeleteMessage удаляет сообщение по ID
	// DeleteMessage(id int) error
}