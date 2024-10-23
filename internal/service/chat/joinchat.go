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
	return nil, nil
}
