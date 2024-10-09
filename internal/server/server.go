package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// Server представляет gRPC сервер для ChatAPI.
type Server struct {
	grpcPort int
	chat     *ChatServer
}

// NewServer создает новый экземпляр Server.
func NewServer(grpcPort int, chat *ChatServer) *Server {
	return &Server{
		grpcPort: grpcPort,
		chat:     chat,
	}
}

// Start запускает gRPC сервер.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.grpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	desc.RegisterChatAPIV1Server(grpcServer, s.chat)

	log.Printf("server listening at :%v", s.grpcPort)

	return grpcServer.Serve(lis)
}
