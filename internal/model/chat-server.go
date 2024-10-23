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

// JoinChat
type JoinChat struct {
	*Chat
	Members  []string
	Messages []*Message
}

type CreateChat struct {
	ChatName     string
	MembersEmail []string
}
