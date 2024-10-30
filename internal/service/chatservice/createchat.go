package chatservice

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

// CreateChat - Логика создания чата
func (s *service) CreateChat(ctx context.Context, chat *model.CreateChat) (int64, error) {
	if chat == nil {
		return 0, ErrCreateChatReq()
	}

	var id int

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTX error

		id, errTX = s.chatRepository.CreateChat(ctx, chat.ChatName)
		if errTX != nil {
			return errTX
		}

		for _, member := range chat.MembersEmail {
			m := &model.ChatMember{
				ChatID:    id,
				UserEmail: member,
			}

			errTX = s.chatRepository.CreateChatMember(ctx, m)
			if errTX != nil {
				return errTX
			}
		}

		if err := s.writeLog(ctx, actionTypeCreateChat); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return int64(id), nil
}
