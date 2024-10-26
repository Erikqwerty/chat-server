package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// DeleteChat - обрабатывает запрос на удаление чата
func (i *ImplChatServer) DeleteChat(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	err := i.chatService.DeleteChat(ctx, req.ChatId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
