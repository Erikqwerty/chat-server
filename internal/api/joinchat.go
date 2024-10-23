package api

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/convertor"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// JoinChat - Обрабатывает запрос на подключение к чату
func (i *Implementation) JoinChat(ctx context.Context, req *desc.JoinChatRequest) (*desc.JoinChatResponse, error) {
	joinChat, err := i.chatService.JoinChat(ctx, convertor.ToModelChatMemberFromJoinChatRequest(req))
	if err != nil {
		return nil, err
	}
	return convertor.ToChatAPIJoinRespFromModelJoinChat(joinChat), nil
}
