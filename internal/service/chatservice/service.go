package chatservice

import (
	"context"

	"google.golang.org/grpc/peer"

	"github.com/erikqwerty/chat-server/internal/repository"
	dev "github.com/erikqwerty/chat-server/internal/service"
	"github.com/erikqwerty/chat-server/pkg/db"
)

var _ dev.ChatService = (*service)(nil)

const (
	actionTypeCreateChat  = "CREATE_CHAT"
	actionTypeSendMessage = "SEND_MESSAGE"
	actionTypeJoinChat    = "JOIN_CHAT"
	actionTypeDeleteChat  = "DELETE_CHAT"
)

type service struct {
	chatRepository repository.ChatServerRepository
	txManager      db.TxManager
}

// NewService - создает новый обьект сервисного слоя
func NewService(chatRepository repository.ChatServerRepository, txManager db.TxManager) dev.ChatService {
	return &service{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}

// details - информация о пользователе
func details(ctx context.Context) string {
	info := "Адрес:"

	peer, _ := peer.FromContext(ctx)
	if peer != nil {
		info += peer.Addr.String()
	} else {
		info = "детальная информация отсутствует"
	}

	return info
}
