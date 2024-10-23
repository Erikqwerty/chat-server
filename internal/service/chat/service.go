package chatservice

import (
	"github.com/erikqwerty/chat-server/internal/client/db"
	"github.com/erikqwerty/chat-server/internal/repository"
	dev "github.com/erikqwerty/chat-server/internal/service"
)

var _ dev.ChatService = (*service)(nil)

type service struct {
	chatRepository repository.ChatServerRepository
	txManager      db.TxManager
}

func NewService(chatRepository repository.ChatServerRepository, txManager db.TxManager) dev.ChatService {
	return &service{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
