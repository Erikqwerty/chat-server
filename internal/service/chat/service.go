package chatservice

import (
	"context"
	"time"

	"github.com/erikqwerty/chat-server/internal/client/db"
	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/repository"
	dev "github.com/erikqwerty/chat-server/internal/service"
	"google.golang.org/grpc/peer"
)

var _ dev.ChatService = (*service)(nil)

const (
	actionTypeCreateChat  = "CreateChat"
	actionTypeSendMessage = "SendMessage "
	actionTypeJounChat    = "JounChat"
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

// createLog - записывает лог в базу даных
func (s *service) createLog(ctx context.Context, actionType string) error {
	errTx := s.chatRepository.CreateLog(ctx, &model.Log{
		ActionType:      actionType,
		ActionDetails:   details(ctx),
		ActionTimestamp: timeNowUTC3(),
	})

	if errTx != nil {
		return errTx
	}
	return nil
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
