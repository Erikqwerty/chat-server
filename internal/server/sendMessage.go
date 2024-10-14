package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// SendMessage отправляет сообщение в чат.
func (s *ChatServer) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Отправка сообщения на сервер: User: %v; message: %v; time: %v", req.From, req.Text, req.Timestamp)
	return nil, nil
}
