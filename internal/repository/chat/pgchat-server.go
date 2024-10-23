package chatserverrepository

import (
	"github.com/erikqwerty/chat-server/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ repository.ChatServerRepository = (*repo)(nil)

type repo struct {
	*repoChat
	*repoChatMember
	*repoMessage
}

// NewRepo - Создает новый обьект repo, для работы с базой данных
func NewRepo(p *pgxpool.Pool) *repo {
	return &repo{
		repoChat:       &repoChat{pool: p},
		repoChatMember: &repoChatMember{pool: p},
		repoMessage:    &repoMessage{pool: p},
	}
}
