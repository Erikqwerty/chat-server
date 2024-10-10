package main

import (
	"log"

	"github.com/erikqwerty/chat-server/internal/server"
)

const grpcPort = 50051

func main() {
	chatServer := &server.ChatServer{}
	srv := server.NewServer(grpcPort, chatServer)

	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
