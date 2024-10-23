package chatservice

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

func (s *service) SendMessage(ctx context.Context, msg *model.Message) error {
	_, err := s.chatRepository.CreateMessage(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}
