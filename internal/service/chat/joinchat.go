package chatservice

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/model"
)

func (s *service) JoinChat(ctx context.Context, chatMember *model.ChatMember) (*model.JoinChat, error) {
	err := s.chatRepository.CreateChatMember(ctx, chatMember)
	if err != nil {
		return nil, err
	}

	chat, err := s.chatRepository.ReadChat(ctx, chatMember.ChatID)
	if err != nil {
		return nil, err
	}

	members, err := s.chatRepository.ReadChatMembers(ctx, chatMember.ChatID)
	if err != nil {
		return nil, err
	}

	emails := make([]string, len(members))
	for i, member := range members {
		emails[i] = member.UserEmail
	}

	messages, err := s.chatRepository.ReadMessages(ctx, chatMember.ChatID)
	if err != nil {
		return nil, err
	}

	return &model.JoinChat{
		Chat:     chat,
		Members:  emails,
		Messages: messages,
	}, nil
}
