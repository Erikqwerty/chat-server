package server

import (
	"context"
	"log"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// JoinChat Подключение к чату.
func (s *ChatServer) JoinChat(ctx context.Context, req *desc.JoinChatRequest) (*desc.JoinChatResponse, error) {
	log.Printf("Пользовтель %v хочет присоединится к чату ID %v ", req.UserEmail, req.ChatId)

	err := s.DB.CreateChatMember(ctx, int(req.ChatId), req.UserEmail)
	if err != nil {
		return nil, err
	}
	emails, err := s.emailMembers(ctx, int(req.ChatId))
	if err != nil {
		return nil, err
	}
	messages, err := s.messagesMembers(ctx, int(req.ChatId))
	if err != nil {
		return nil, err
	}
	chat, err := s.DB.ReadChat(ctx, int(req.ChatId))
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
