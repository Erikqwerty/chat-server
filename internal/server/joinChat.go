package server

import (
	"context"
	"fmt"
	"log"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// JoinChat - Обрабатывае подключение к чату.
func (s *ChatServer) JoinChat(ctx context.Context, req *desc.JoinChatRequest) (*desc.JoinChatResponse, error) {
	log.Printf("Пользовтель %v хочет присоединится к чату ID %v ", req.UserEmail, req.ChatId)

	// проверяем валидность переданного email
	_, err := validateEmails([]string{req.UserEmail})
	if err != nil {
		return nil, err
	}

	// получаю чат и проверяю его наличие
	chat, err := s.DB.ReadChat(ctx, int(req.ChatId))
	if err != nil {
		return nil, fmt.Errorf("чат c id %v не существует либо не был найден: %w", req.ChatId, err)
	}

	// Добавляем пользователя в чат
	err = s.DB.CreateChatMember(ctx, int(req.ChatId), req.UserEmail)
	if err != nil {
		return nil, err
	}

	// Получаем список всех пользователей
	emails, err := s.emailMembers(ctx, int(req.ChatId))
	if err != nil {
		return nil, err
	}

	// Получаем список сообщений
	messages, err := s.messagesMembers(ctx, int(req.ChatId))
	if err != nil {
		return nil, err
	}

	return &desc.JoinChatResponse{
		ChatId:       req.ChatId,
		ChatName:     chat.ChatName,
		Participants: emails,
		Messages:     messages,
	}, nil
}

// emailMembers - возвращает список пользовтелей чата
func (s *ChatServer) emailMembers(ctx context.Context, chatID int) ([]string, error) {
	members, err := s.DB.ReadChatMembers(ctx, chatID)
	if err != nil {
		return nil, err
	}
	emails := make([]string, 0, 10)
	for i := 0; i < len(members); i++ {
		emails = append(emails, members[i].UserEmail)
	}
	return emails, nil
}

// messagesMembers - возвращает список сообщений из чата
func (s *ChatServer) messagesMembers(ctx context.Context, chatID int) ([]*desc.Message, error) {
	messagesDB, err := s.DB.ReadMessages(ctx, chatID)
	if err != nil {
		return nil, err
	}

	responseMessges := make([]*desc.Message, 0, 10)

	for i := 0; i < len(messagesDB); i++ {
		responseMessges = append(responseMessges,
			&desc.Message{
				Text:      messagesDB[i].Text,
				From:      messagesDB[i].UserEmail,
				Timestamp: timestamppb.New(messagesDB[i].Timestamp),
			})
	}

	return responseMessges, err
}
