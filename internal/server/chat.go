package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// ChatServer реализует методы ChatAPIV1.
type ChatServer struct {
	desc.UnimplementedChatAPIV1Server
}

// Create обрабатывает создание нового чата.
func (s *ChatServer) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Cоздание нового чата: %v", req.UserIds)
	return &desc.CreateResponse{Id: 1}, nil
}

// Delete обрабатывает удаление чата.
func (s *ChatServer) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Удаление чата из системы по его идентификатору: %v", req.Id)
	return nil, nil
}

// SendMessage отправляет сообщение в чат.
func (s *ChatServer) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Отправка сообщения на сервер: User: %v; message: %v; time: %v", req.From, req.Text, req.Timestamp)
	return nil, nil
}
