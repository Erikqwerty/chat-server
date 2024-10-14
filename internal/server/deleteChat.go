package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// DeleteChat обрабатывает удаление чата.
func (s *ChatServer) DeleteChat(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Удаление чата из системы по его идентификатору: %v", req.Id)
	return nil, nil
}
