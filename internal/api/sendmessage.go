package api

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/erikqwerty/chat-server/internal/convertor"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// SendMessage - обрабатывает запрос на отправку сообщения
func (i *ImplChatServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	if req.From == "" {
		return nil, errors.New("не указан отправитель сообщения")
	}

	if req.Text == "" {
		return nil, errors.New("текст сообщения пустой")
	}

	if req.ChatId == 0 {
		return nil, errors.New("не указан чат в который отправляется сообщение")
	}

	if !isValidEmail(req.From) {
		return nil, errors.New("email не валиден")
	}

	err := i.chatService.SendMessage(ctx, convertor.ToModelMessageFromReqSendMessage(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
