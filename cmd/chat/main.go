package main

import (
	"flag"
	"log"

	"github.com/erikqwerty/chat-server/internal/server"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()

	chatServer, err := server.NewChatApp(configPath)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(chatServer)

	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
