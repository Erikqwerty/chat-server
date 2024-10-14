package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// DeleteChat обрабатывает удаление чата.
func (s *ChatServer) DeleteChat(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.DB.DeleteChat(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	log.Printf("Удаление чата из системы по его идентификатору: %v", req.Id)
	return nil, nil
}
