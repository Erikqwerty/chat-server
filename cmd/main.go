package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50051

// ChatServer используется для реализации сервера gRPC, который предоставляет методы, описанные в ChatAPIV1
type ChatServer struct {
	desc.UnimplementedChatAPIV1Server
}

// Create обрабатывает создание нового чата.
func (s *ChatServer) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Cоздание нового чата: %v", req.Usernames)
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

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatAPIV1Server(s, &ChatServer{})

	log.Printf("server listening at :%v", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Faider to server: ", err)
	}
}
