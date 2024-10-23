package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/erikqwerty/chat-server/internal/convertor"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.chatService.SendMessage(ctx, convertor.ToModelMessageFromReqSendMessage(req))
	if err != nil {
		return nil, err
	}
	return nil, nil
}
