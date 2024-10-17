package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"

	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// CreateChat обрабатывает создание нового чата.
func (s *ChatServer) CreateChat(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	validEmails, errv := validateEmails(req.Emails)

	// TODO: Нужна middleware или какая-то проверка, зарегистрованных пользователей,
	// чтобы не создавать чат с email не существующих пользователей

	chatID, err := s.DB.CreateChat(ctx, req.ChatName)
	if err != nil {
		return &desc.CreateResponse{}, err
	}

	log.Printf("Cоздание нового чата: %v", validEmails)
	return &desc.CreateResponse{Id: int64(chatID)}, errv
}

// validateEmails - проверяет список переданных email на валидность и возвращает валидный список email,
// а также ошибку с невалидными email
func validateEmails(emails []string) ([]string, error) {
	validEmail := make([]string, 0)
	var err string
	for _, email := range emails {
		if isValidEmail(email) {
			validEmail = append(validEmail, email)
		} else {
			err += fmt.Sprintf("ошибка не валидный email: %v;", email)
		}
	}

	if err != "" {
		return validEmail, errors.New(err)
	}
	return validEmail, nil
}

// isValidEmail - проверяет валидность email. Возвращает true если валидно.
func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
