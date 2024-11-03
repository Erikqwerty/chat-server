package chatservice

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

func (s *service) DeleteChat(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTX error

		errTX = s.chatRepository.DeleteChat(ctx, int(id))
		if errTX != nil {
			return errTX
		}

		errTX = s.chatRepository.CreateLog(ctx, &model.Log{
			ActionType:      actionTypeDeleteChat,
			ActionDetails:   details(ctx),
			ActionTimestamp: timeNowUTC3(),
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
