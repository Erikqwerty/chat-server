package chatservice

import (
	"context"
	"time"

	"github.com/erikqwerty/chat-server/internal/model"
)

// CreateChat - Логика создания чата
func (s *service) CreateChat(ctx context.Context, chat *model.CreateChat) (int64, error) {

	id, err := s.chatRepository.CreateChat(ctx, chat.ChatName)
	if err != nil {
		return 0, err
	}
	for _, member := range chat.MembersEmail {
		m := &model.ChatMember{
			ChatID:    id,
			UserEmail: member,
			JoinedAt:  timeNowUTC3(),
		}
		err := s.chatRepository.CreateChatMember(ctx, m)
		if err != nil {
			return 0, err
		}
	}

	return int64(id), nil
}

// timeNowUTC3 + возвращает время +3
func timeNowUTC3() time.Time {
	return time.Now().In(time.FixedZone("UTC+3", 3*60*60))
}
