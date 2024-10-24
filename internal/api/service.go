package api

import (
	"github.com/erikqwerty/chat-server/internal/service"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// ImplChatServer - имплементирует gRPC методы
type ImplChatServer struct {
	desc.ChatAPIV1Server
	chatService service.ChatService
}

// NewChatServerGRPCImplementation - Создает новый обьект имплементирующий gRPC сервер
func NewChatServerGRPCImplementation(chatService service.ChatService) *ImplChatServer {
	return &ImplChatServer{chatService: chatService}
}
