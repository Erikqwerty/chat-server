package server

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// SendMessage отправляет сообщение в чат.
func (s *ChatServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := s.existsEmailInChat(ctx, req.From, req.ChatId)
	if err != nil {
		return nil, err
	}

	_, err = s.DB.CreateMessage(ctx, int(req.ChatId), req.From, req.Text)
	if err != nil {
		return nil, err
	}

	log.Printf("Отправка сообщения на сервер: User: %v; message: %v; time: %v", req.From, req.Text, req.Timestamp)
	return nil, nil

}

// existsEmailInChat - Проверяем может ли пользователь отправить сообщение в указанный чат
// возвращаем True если может.
func (s *ChatServer) existsEmailInChat(ctx context.Context, fromEmail string, chatID int64) error {
	chat, err := s.DB.ReadChat(ctx, int(chatID))
	if err != nil {
		return fmt.Errorf("чат с указанным id не существует: %w", err)
	}
	if chat.ID > 0 {
		member, err := s.DB.ReadChatMember(ctx, fromEmail, chat.ID)
		if err != nil {
			return err
		}
		if member != nil {
			return nil // пользователь состоит в чате
		}
	}
	return fmt.Errorf("пользователь %v не состоит в чате %v", fromEmail, chatID)
}
