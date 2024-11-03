package chatservice

import (
	"context"
	"time"

	"google.golang.org/grpc/peer"

	"github.com/erikqwerty/chat-server/internal/repository"
	dev "github.com/erikqwerty/chat-server/internal/service"
	"github.com/erikqwerty/chat-server/pkg/db"
)

var _ dev.ChatService = (*service)(nil)

const (
	actionTypeCreateChat  = "CreateChat"
	actionTypeSendMessage = "SendMessage "
	actionTypeJoinChat    = "JounChat"
	actionTypeDeleteChat  = "DeleteChat "
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
	info += peer.Addr.String()

	return info
}

// timeNowUTC3 + возвращает время +3
func timeNowUTC3() time.Time {
	return time.Now().In(time.FixedZone("UTC+3", 3*60*60))
}
