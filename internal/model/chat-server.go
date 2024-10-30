package model

import "time"

// Chat представляет чат с его атрибутами
type Chat struct {
	ID        int
	ChatName  string
	CreatedAt time.Time
}

// ChatMember представляет участника чата
type ChatMember struct {
	ChatID    int
	UserEmail string
	JoinedAt  time.Time
}

// Message представляет сообщение в базе данных
type Message struct {
	ID        int
	ChatID    int
	UserEmail string
	Text      string
	Timestamp time.Time
}

// JoinChat - представляет ответ из сервисного слоя на добавление нового участника чата
type JoinChat struct {
	*Chat
	Members  []string
	Messages []*Message
}

// CreateChat - преставляет запрос на создание чата в сервисном слое
type CreateChat struct {
	ChatName     string
	MembersEmail []string
}

// Log - структура для логирования действий в БД
type Log struct {
	ID              int64
	ActionType      string
	ActionDetails   string
	ActionTimestamp time.Time
}
