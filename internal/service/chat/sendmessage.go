package chatservice

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

func (s *service) SendMessage(ctx context.Context, msg *model.Message) error {

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		_, err := s.chatRepository.CreateMessage(ctx, msg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
