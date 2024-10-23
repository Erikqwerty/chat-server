package api

import (
	"context"

	"github.com/erikqwerty/chat-server/internal/convertor"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// CreateChat - обрабатывает запрос на создание чата
func (i *Implementation) CreateChat(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.CreateChat(ctx, convertor.ToModelCreateChatFromCreateReq(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}
