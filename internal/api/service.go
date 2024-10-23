package api

import (
	"github.com/erikqwerty/chat-server/internal/service"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// Implementation - имплементирует gRPC методы
type Implementation struct {
	desc.ChatAPIV1Server
	chatService service.ChatService
}

// NewImplementation - Создает новый обьект имплементирующий gRPC сервер
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{chatService: chatService}
}
