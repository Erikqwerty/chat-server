package server

import (
	"context"
	"log"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// JoinChat Подключение к чату.
func (s *ChatServer) JoinChat(_ context.Context, req *desc.JoinChatRequest) (*desc.JoinChatResponse, error) {
	log.Printf("Пользовтель %v хочет присоединится к чату ID %v ", req.UserEmail, req.ChatId)
	return nil, nil
}
