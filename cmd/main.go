package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/erikqwerty/chat-server/pkg/ChatAPI_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50051

type chat_server struct {
	desc.UnimplementedChatAPIV1Server
}

func (s *chat_server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Usernames: %v", req.Usernames)
	return &desc.CreateResponse{Id: 1}, nil
}

func (s *chat_server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user_id: %v", req.Id)
	return nil, nil
}

func (s *chat_server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("User: %v; message: %v; time: %v", req.From, req.Text, req.Timestamp)
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatAPIV1Server(s, &chat_server{})

	log.Printf("server listening at :%v", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Faider to server: ", err)
	}
}
