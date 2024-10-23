package api

import (
	"context"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteChat(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatService.DeleteChat(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
