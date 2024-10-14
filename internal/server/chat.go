package server

import (
	"fmt"

	"github.com/erikqwerty/chat-server/internal/config"
	"github.com/erikqwerty/chat-server/internal/db"
	"github.com/erikqwerty/chat-server/internal/db/pg"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// ChatServer реализует методы ChatAPIV1.
type ChatServer struct {
	desc.UnimplementedChatAPIV1Server
	Config *config.Config
	DB     db.DB
}

// NewAuthApp - Создает структуру приложения чата, загружая конфигурации
func NewChatApp(path string) (*ChatServer, error) {
	chat := &ChatServer{}
	conf, err := config.New(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурации %v", err)
	}
	chat.Config = conf

	database, err := pg.New(conf.DB.DSN())
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения dsn для подключения к базе данных %v", err)
	}
	chat.DB = database

	return chat, nil
}
