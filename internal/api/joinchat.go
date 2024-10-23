package api

import (
	"context"
	"errors"

	"github.com/erikqwerty/chat-server/internal/convertor"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// JoinChat - Обрабатывает запрос на подключение к чату
func (i *ImplChatServer) JoinChat(ctx context.Context, req *desc.JoinChatRequest) (*desc.JoinChatResponse, error) {

	if req.ChatId == 0 {
		return nil, errors.New("не указан id чата")
	}

	if req.UserEmail == "" {
		return nil, errors.New("не указан пользовтель для добавления в чат")
	}

	if !isValidEmail(req.UserEmail) {
		return nil, errors.New("email не валиден")
	}

	joinChat, err := i.chatService.JoinChat(ctx, convertor.ToModelChatMemberFromJoinChatRequest(req))
	if err != nil {
		return nil, err
	}

	return convertor.ToChatAPIJoinRespFromModelJoinChat(joinChat), nil
}
