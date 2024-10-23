package chatservice

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

func (s *service) JoinChat(ctx context.Context, chatMember *model.ChatMember) (*model.JoinChat, error) {
	var (
		chat     *model.Chat
		members  []*model.ChatMember
		messages []*model.Message
	)

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		err := s.chatRepository.CreateChatMember(ctx, chatMember)
		if err != nil {
			return err
		}

		chat, err = s.chatRepository.ReadChat(ctx, chatMember.ChatID)
		if err != nil {
			return err
		}

		members, err = s.chatRepository.ReadChatMembers(ctx, chatMember.ChatID)
		if err != nil {
			return err
		}

		messages, err = s.chatRepository.ReadMessages(ctx, chatMember.ChatID)
		if err != nil {
			return err
		}

		if err := s.createLog(ctx, actionTypeJounChat); err != nil {
			return err
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
