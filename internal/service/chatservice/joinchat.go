package chatservice

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

func (s *service) JoinChat(ctx context.Context, chatMember *model.ChatMember) (*model.JoinChat, error) {
	if chatMember == nil {
		return nil, ErrJoinChatReq()
	}

	var (
		chat     *model.Chat
		members  []*model.ChatMember
		messages []*model.Message
	)

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTX error

		errTX = s.chatRepository.CreateChatMember(ctx, chatMember)
		if errTX != nil {
			return errTX
		}

		chat, errTX = s.chatRepository.ReadChat(ctx, chatMember.ChatID)
		if errTX != nil {
			return errTX
		}

		members, errTX = s.chatRepository.ReadChatMembers(ctx, chatMember.ChatID)
		if errTX != nil {
			return errTX
		}

		messages, errTX = s.chatRepository.ReadMessages(ctx, chatMember.ChatID)
		if errTX != nil {
			return errTX
		}

		errTX = s.chatRepository.CreateLog(ctx, &model.Log{
			ActionType:      actionTypeJoinChat,
			ActionDetails:   details(ctx),
			ActionTimestamp: timeNowUTC3(),
		})
		if errTX != nil {
			return errTX
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &model.JoinChat{
		Chat:     chat,
		Members:  membersEmails(members),
		Messages: messages,
	}, nil
}

func membersEmails(members []*model.ChatMember) []string {
	emails := make([]string, len(members))
	for i, member := range members {
		emails[i] = member.UserEmail
	}
	return emails
}
