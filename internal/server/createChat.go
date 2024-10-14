package server

import (
	"context"
	"log"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// CreateChat обрабатывает создание нового чата.
func (s *ChatServer) CreateChat(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	chatID, err := s.DB.CreateChat(ctx, req.ChatName)
	if err != nil {
		return &desc.CreateResponse{}, err
	}

	log.Printf("Cоздание нового чата: %v", req.Emails)
	return &desc.CreateResponse{Id: int64(chatID)}, nil
}
