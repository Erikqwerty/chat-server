package server

import (
	"context"
	"log"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// CreateChat обрабатывает создание нового чата.
func (s *ChatServer) CreateChat(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Cоздание нового чата: %v", req.Emails)
	return &desc.CreateResponse{Id: 1}, nil
}
