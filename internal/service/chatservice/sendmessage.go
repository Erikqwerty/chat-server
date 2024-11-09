package chatservice

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

func (s *service) SendMessage(ctx context.Context, msg *model.Message) error {
	if msg == nil {
		return ErrSendMessage()
	}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTX error

		_, errTX = s.chatRepository.CreateMessage(ctx, msg)
		if errTX != nil {
			return errTX
		}

		errTX = s.chatRepository.CreateLog(ctx, &model.Log{
			ActionType:    actionTypeSendMessage,
			ActionDetails: details(ctx),
		})
		if errTX != nil {
			return errTX
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
