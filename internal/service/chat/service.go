package chatservice

import (
	"github.com/erikqwerty/chat-server/internal/repository"
	dev "github.com/erikqwerty/chat-server/internal/service"
)

var _ dev.ChatService = (*service)(nil)

type service struct {
	chatRepository repository.ChatServerRepository
}

func NewService(chatRepository repository.ChatServerRepository) dev.ChatService {
	return &service{
		chatRepository: chatRepository,
	}
}
