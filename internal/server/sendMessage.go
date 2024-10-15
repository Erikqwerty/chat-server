package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// SendMessage отправляет сообщение в чат.
func (s *ChatServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {

	_, err := s.DB.CreateMessage(ctx, int(req.ChatId), req.From, req.Text)
	if err != nil {
		return nil, err
	}

	log.Printf("Отправка сообщения на сервер: User: %v; message: %v; time: %v", req.From, req.Text, req.Timestamp)
	return nil, nil
}

// TODO: Надо проверять а существует ли чат в который меседж отправляем, и email отправителя находится ли в нужном чате
