package api

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/convertor"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// JoinChat - Обрабатывает запрос на подключение к чату
func (i *ImplChatServer) JoinChat(ctx context.Context, req *desc.JoinChatRequest) (*desc.JoinChatResponse, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	joinChat, err := i.chatService.JoinChat(ctx, convertor.ToModelChatMemberFromJoinChatRequest(req))
	if err != nil {
		return nil, err
	}

	return convertor.ToChatAPIJoinRespFromModelJoinChat(joinChat), nil
}
