package api

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/erikqwerty/chat-server/internal/convertor"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

// CreateChat - обрабатывает запрос на создание чата
func (i *ImplChatServer) CreateChat(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.ChatName == "" {
		return nil, errors.New("не указанно название чата")
	}

	if len(req.Emails) == 0 {
		return nil, errors.New("не указанны участники чата")
	}

	err := validEmails(req.Emails)
	if err != nil {
		return nil, err
	}

	id, err := i.chatService.CreateChat(ctx, convertor.ToModelCreateChatFromCreateReq(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}

// validEmails - возвращает ошибку в виде списока не валидных email
func validEmails(emails []string) error {
	var errtext string

	for _, email := range emails {
		if !isValidEmail(email) {
			errtext += fmt.Sprintf("email: %v не валиден;", email)
		}
	}

	if errtext != "" {
		return errors.New(errtext)
	}

	return nil
}

// isValidEmail проверяет валидность email-адреса. Возвращает true если валидно.
func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}
