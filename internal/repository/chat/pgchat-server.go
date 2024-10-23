package chatserverrepository

import (
	"github.com/erikqwerty/chat-server/internal/client/db"
	"github.com/erikqwerty/chat-server/internal/repository"
)

var _ repository.ChatServerRepository = (*repo)(nil)

type repo struct {
	repoChat
	repoChatMember
	repoMessage
	repoLoger
}

// NewRepo - Создает новый обьект repo, для работы с базой данных
func NewRepo(dbc db.Client) *repo {
	return &repo{
		repoChat:       repoChat{db: dbc},
		repoChatMember: repoChatMember{db: dbc},
		repoMessage:    repoMessage{db: dbc},
		repoLoger:      repoLoger{db: dbc},
	}
}
