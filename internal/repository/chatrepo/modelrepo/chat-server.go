package modelrepo

import "time"

// Chat представляет чат с его атрибутами
type Chat struct {
	ID        int       `db:"id"`
	ChatName  string    `db:"chat_name"`
	CreatedAt time.Time `db:"created_at"`
}

// ChatMember представляет участника чата
type ChatMember struct {
	ChatID    int       `db:"chat_id"`
	UserEmail string    `db:"user_email"`
	JoinedAt  time.Time `db:"joined_at"`
}

// Message представляет сообщение в базе данных
type Message struct {
	ID        int       `db:"id"`
	ChatID    int       `db:"chat_id"`
	UserEmail string    `db:"user_email"`
	Text      string    `db:"text"`
	Timestamp time.Time `db:"timestamp"`
}

// Log - структура для логирования действий в БД
type Log struct {
	ID              int64     `db:"id"`
	ActionType      string    `db:"action_type"`
	ActionDetails   string    `db:"action_details"`
	ActionTimestamp time.Time `db:"action_timestamp"`
}
